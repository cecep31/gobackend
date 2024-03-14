package server

import (
	"os"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"

	"gobackend/pkg"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"go.uber.org/zap"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	if os.Getenv("ALLOW_ORIGIN") != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "https://pilput.dev, https://dash.pilput.dev, http://pilput.test, http://localhost:3000, http://localhost:5432",
		}))
	} else {
		app.Use(cors.New(cors.Config{}))
	}
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Use(etag.New())
	if os.Getenv("LIMITER") != "" {
		max, _ := strconv.Atoi(os.Getenv("LIMITER"))
		app.Use(limiter.New(limiter.Config{
			Max: max,
		}))
	}
	logger, _ := zap.NewProduction()
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

}

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
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
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		BodyLimit:   100 * 1024 * 1024,
	})

	setupMiddlewares(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "server ready",
		})
	})

	return app
}

func Listen(app *fiber.App) error {
	app.Use(handleNotFound)
	return app.Listen(getPort())
}

func handleNotFound(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}

func getPort() string {
	envPort := os.Getenv("PORT")
	defaultPort := ":8080"
	return ":" + strings.TrimPrefix(envPort, ":") + defaultPort[1:]
}
