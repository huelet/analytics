package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
