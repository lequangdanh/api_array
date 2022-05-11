package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: 1, Name: "walk the dog", Completed: false},
	{Id: 2, Name: "walk the cat", Completed: false},
	{Id: 3, Name: "walk the fish", Completed: false},
}

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})
	app.Get("/todos", GetTodos)
	app.Post("/todos", CreateTodo)
	app.Get("/todo/{id}", GetTodo)
	app.Listen(":8080")

}
func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}
func CreateTodo(ctx *fiber.Ctx) error {
	body := new(Todo)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	todo := Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}
	todos = append(todos, todo)
	return ctx.Status(fiber.StatusOK).JSON(todos)
}
func GetTodo(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return err
	}
	for _, todo := range todos {
		if todo.Id == id {
			return ctx.Status(fiber.StatusOK).JSON(todo)

		}
	}
	return ctx.Status(fiber.StatusOK).JSON(todos)

}
