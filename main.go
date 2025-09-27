package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	// This is a placeholder for the main function.
	fmt.Println("Hello, World!")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})
	// Create a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body is required",
			})
		}

		todo.ID = fmt.Sprintf("%d", len(todos)+1)

		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"error": "Todo Not Found",
		}) // No Content or appropriate response
	})

	// Delete a todo

	// app.Delete("/api/todos/id", func(c, *fiber.Ctx) error {

	// })

	log.Fatal(app.Listen(":4000"))
}
