package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zthiagovalle/trivia-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/facts", handlers.ListFacts)

	app.Post("/api/v1/fact", handlers.CreateFact)
}
