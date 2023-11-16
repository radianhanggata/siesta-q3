package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	
	"github.com/radianhanggata/siesta-coding-test/q3/router"
)

func main() {
	app := fiber.New()
	router.SetupRoute(app)
	log.Fatal(app.Listen(":3000"))
}
