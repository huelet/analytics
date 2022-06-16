package routes

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/huelet/analytics/src"
)

type Visit struct {
	Path    string `fauna:"path"`
	Time    string `fauna:"time"`
	Browser string `fauna:"browser"`
	OS      string `fauna:"os"`
	Country string `fauna:"country"`
}

func PageVisit(app *fiber.App) {
	ctx := context.Background()
	app.Post("/api/visit", func(c *fiber.Ctx) error {
		db, err := src.ConnectDB()
		if err != nil {
			return err
		}
		res, err := db.Get(ctx, "demo").Result()
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	})
}
