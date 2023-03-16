package server

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/goccy/go-json"

	"gobackend/database"
	"gobackend/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/websocket/v2"
)

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var clients = make(map[*websocket.Conn]*client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		// AllowOrigins: "https://pilput.dev, https://pilput.test, http://pilput.test",
		AllowOrigins: "*",
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Use(etag.New())
	if os.Getenv("ENABLE_LIMITER") != "" {
		app.Use(limiter.New())
	}
	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New())
	}
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			c.Locals("allowed", true)
			fmt.Println("jalan websocket")
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

}

func Create() *fiber.App {
	database.SetupDatabase()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*pkg.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(pkg.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(pkg.Error{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
		AppName:     "pilput-turbo",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	setupMiddlewares(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "server ready",
		})
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed")) // true
		log.Println(c.Params("id"))      // 123
		log.Println(c.Query("v"))        // 1.0
		log.Println(c.Cookies("token"))  // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, []byte("hello")); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	return app
}

func Listen(app *fiber.App) error {

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverPort := os.Getenv("SERVER_PORT")
	return app.Listen(fmt.Sprintf(":%s", serverPort))
}
