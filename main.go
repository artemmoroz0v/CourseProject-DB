package main

import (
	"DB/app/controller"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
)

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("/public"))
	r.GET("/", controller.StartPage)
	r.GET("/trainers", controller.GetTrainers)
	r.GET("/clients", controller.GetClients)
	r.GET("/groups", controller.GetGroups)
	r.POST("/trainer/post", controller.PostTrainer)
	r.POST("/client/post", controller.PostClient)
	//r.POST("/group/post", controller.PostGroup)
}

func main() {
	err := server.OpenBD("user=postgres password=454565 dbname=postgres sslmode=disable")
	if err != nil {
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
