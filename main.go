package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"time"
	"todo-togo/controller"
	"todo-togo/db"
	"todo-togo/entity"
	"todo-togo/repository"
	"todo-togo/service"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Final Task! %s", time.Now())
}

func oldmain() {
	//http.HandleFunc("/", greet)
	//http.ListenAndServe(":8088", nil)

	dbConnect := db.NewSqliteConnection()
	todoRepo := repository.NewTodoRepo(dbConnect)

	coba,err:=todoRepo.SelectTodo(entity.Todo{
		ID: 4,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(coba)

	//baru := entity.Todo{
	//	Title:       "Test Insert",
	//	Description: "Insertion test from repos",
	//	DueDate:     "2021-10-05",
	//	PIC:         "bm",
	//	Status:      "New",
	//}
	//coba, err := todoRepo.CreateTodo(baru)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(coba)
	
	/*update := entity.Todo{
		ID:          5,
		Title:       "Test Insert [UPDATED]",
		Description: "Insertion test from repos",
		DueDate:     "2021-10-05",
		PIC:         "bm",
		Status:      "OnGoing",
	}

	coba, err = todoRepo.UpdateTodo(update)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(coba)*/

	/*update := entity.Todo{
		ID:          5,
	}

	err = todoRepo.DeleteTodo(update)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Success!")
*/

	stat := repository.NewStatusRepo(dbConnect)
	a,_:=stat.SelectAllStatus()
	log.Println(a)

	user := repository.NewUserRepo(dbConnect)
	u,_:=user.SelectUser(entity.User{
		UserID: 1,
	})
	log.Println(u)
}

func main() {
	app := fiber.New()

	dbConn := db.NewSqliteConnection()

	todoRepo := repository.NewTodoRepo(dbConn)
	todoService := service.NewTodoService(&todoRepo)
	todoController := controller.NewController(&todoService)

	statusRepo := repository.NewStatusRepo(dbConn)
	statusService := service.NewStatusService(&statusRepo)
	statusService.PrepareAllStatus()

	todoController.Route(app)

	log.Fatal(app.Listen(":8088"))
}