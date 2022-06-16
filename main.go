package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/huelet/analytics/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	routes.PageVisit(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
