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
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPage)

	r.GET("/client", controller.SelectClients)
	r.POST("/client/insert", controller.InsertClient)
	r.POST("/client/update/subscription", controller.UpdateClientSubscription)
	r.POST("/client/update/height&weight", controller.UpdateHeightAndWeight)
	r.POST("/client/delete", controller.DeleteClient)

	r.GET("/trainer", controller.SelectTrainers)
	r.POST("/trainer/insert", controller.InsertTrainer)
	r.POST("/trainer/delete", controller.DeleteTrainer)

	r.GET("/group", controller.SelectGroups)
	r.POST("/group/select", controller.SelectGroup)
	r.POST("/group/client/insert", controller.InsertClientIntoGroup)
	r.POST("/group/client/delete", controller.DeleteClientFromGroup)
	r.POST("/group/delete", controller.DeleteGroup)
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
