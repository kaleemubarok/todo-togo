package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"todo-togo/model"
	"todo-togo/service"
)

type TodoController struct {
	todo service.ITodoService
}

func NewController(todoService *service.ITodoService) TodoController  {
	return	TodoController{todo: *todoService}
}

func (controller *TodoController) Route(app *fiber.App)  {
	app.Get("/todo", controller.GetAll)
	app.Post("/todo", controller.PostTodo)
	app.Put("/todo", controller.Update)
	app.Delete("/todo/:id", controller.Delete)
}

func (controller *TodoController) GetAll(c *fiber.Ctx) error  {
	todos, err := controller.todo.GetAllTodo()
	if err != nil {
		log.Println("error to get all todos,",err.Error())
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error to get all todos",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"todos":   todos,
	})
}

func (controller *TodoController) PostTodo(c *fiber.Ctx) error  {
	var req model.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("error to parse body new todo,",err.Error())
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error to parse new todos",
		})
	}

	newTodo, err := controller.todo.AddTodo(req)
	if err != nil {
		log.Println("error to create new todo,",err.Error())
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error to create new todos",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"todo": newTodo,
	})
}

func (controller *TodoController) Delete(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("error to parse requested id,",err.Error())
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error to get deleted todo",
		})
	}

	deletedID := model.CreateTodoRequest{
		ID: id,
	}

	err = controller.todo.DeleteTodo(deletedID)
	if err != nil {
		log.Println("error to delete todo record,",err.Error())
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error when delete todo",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"todos":   nil,
	})
}

func (controller *TodoController) Update(c *fiber.Ctx) error  {
	var req model.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("error to parse body-request updated todo,",err.Error())
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error to parse updated todos",
		})
	}

	//check if ID exist
	_, err := controller.todo.GetTodo(req)
	if err != nil {
		log.Println("error id not found,",err.Error())
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error on update, record not found",
		})
	}

	todo, err := controller.todo.ModifyTodo(req)
	if err != nil {
		log.Println("error on modify todo record,",err.Error())
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error when modify todo",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"todos":   todo,
	})
}

