package main

import (
	"log"
	"todo-togo/controller"
	"todo-togo/db"
	_ "todo-togo/docs"
	"todo-togo/repository"
	"todo-togo/service"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title TODO-TOGO a Todo API
// @version 1.0
// @description This is an auto-generated API Docs.
// @host localhost:8088
// @BasePath /
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.Handler) // swagger route

	dbConn := db.NewSqliteConnection()

	todoRepo := repository.NewTodoRepo(dbConn)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewController(&todoService)

	statusRepo := repository.NewStatusRepo(dbConn)
	statusService := service.NewStatusService(statusRepo)
	statusService.PrepareAllStatus()

	userRepo := repository.NewUserRepo(dbConn)
	userService := service.NewUserService(&userRepo)
	userController := controller.NewUserController(&userService)

	todoController.Route(app)
	userController.UserRoute(app)

	log.Fatal(app.Listen(":8088"))
}
