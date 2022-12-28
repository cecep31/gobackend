package payments

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Post("payment/linkaja", CreateLinkajaPayment)
}
