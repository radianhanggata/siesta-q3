package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/radianhanggata/siesta-coding-test/q3/handler"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/format/:countrycode/:phonenumber", handler.FormatPhoneNumber)
}
