package main

import (
	"DB/app/controller"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("/public"))
	r.GET("/", controller.StartPage)
	r.GET("/trainers", controller.SelectTrainers)
	r.GET("/clients", controller.SelectClients)
	r.GET("/groups", controller.SelectGroups)
	r.POST("/trainer/post", controller.InsertTrainer)
	r.POST("/client/post", controller.InsertClient)
	r.POST("/group/choose", controller.SelectGroup)
	r.POST("/client/update", controller.UpdateSubscription)
	r.POST("/group/append", controller.InsertClientIntoGroup)
	//r.POST("/group/post", controller.PostGroup)
}

func main() {
	err := server.OpenBD("user=postgres password=454565 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		err = server.CloseDB()
		if err != nil {
			return
		}
	}()
	r := httprouter.New()
	routes(r)
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		return
	}
}
