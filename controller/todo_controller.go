package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	_ "todo-togo/docs"
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

// GetAll godoc
// @Description Get all todo list without any params required
// @Summary Get all todos list
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} model.JsonSuccessTodosReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /todo [get]
func (controller *TodoController) GetAll(c *fiber.Ctx) error  {
	todos, err := controller.todo.GetAllTodo()
	if err != nil {
		log.Println("error to get all todos,",err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error to get all todos",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"todos":   todos,
	})
}

// PostTodo godoc
// @Description Add new single todo task
// @Summary Add new single todo task
// @Tags Todo
// @Accept json
// @Produce json
// @Param Task body model.CreateTodoRequest true "Add todo"
// @Success 200 {object} model.JsonSuccessTodosReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /todo [post]
func (controller *TodoController) PostTodo(c *fiber.Ctx) error  {
	var req model.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("error to parse body new todo,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error to parse new todos",
		})
	}

	newTodo, err := controller.todo.AddTodo(req)
	if err != nil {
		log.Println("error to create new todo,",err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error to create new todos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"todo": newTodo,
	})
}

// DeleteTodo godoc
// @Description Delete a todo task by task ID
// @Summary Delete a todo task
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} model.JsonDeleteTodosReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /todo/{id} [delete]
func (controller *TodoController) Delete(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("error to parse requested id,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
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
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error when delete todo",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"todos":   nil,
	})
}

// UpdateTodo godoc
// @Description Update todo task
// @Summary Update todo task
// @Tags Todo
// @Accept json
// @Produce json
// @Param Task body model.CreateTodoRequest true "Modify todo"
// @Success 200 {object} model.JsonSuccessTodosReturn
// @Failure 500 {object} model.JsonInternalErrorReturn
// @Failure 404 {object} model.JsonBadRequestErrorReturn
// @Router /todo [put]
func (controller *TodoController) Update(c *fiber.Ctx) error  {
	var req model.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("error to parse body-request updated todo,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error to parse updated todos",
		})
	}

	//check if ID exist
	_, err := controller.todo.GetTodo(req)
	if err != nil {
		log.Println("error id not found,",err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "Error on update, record not found",
		})
	}

	//todo move this to service (validate if deleted)
	if req.Status > 2 {
		log.Println("unable to update deleted / finished task")
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error": "unable to update deleted / finished task",
		})
	}

	todo, err := controller.todo.ModifyTodo(req)
	if err != nil {
		log.Println("error on modify todo record,",err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error": "Error when modify todo",
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"todos":   todo,
	})
}

