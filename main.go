package main

import (

	"project-two/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/movies", handlers.GetMovies)

	app.Patch("/movies/:id", handlers.UpdateMovie)

	app.Get("/movies/:id", handlers.GetMovieByID)

	app.Post("/movies", handlers.PostMovie)

	app.Delete("/movies/:id", handlers.DeleteMovie)

	app.Put("/movies/:id", handlers.PutMovie)

	app.Listen(":3000")
}
