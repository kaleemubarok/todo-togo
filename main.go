package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"todo-togo/db"
	"todo-togo/entity"
	"todo-togo/repository"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Final Task! %s", time.Now())
}

func main() {
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
