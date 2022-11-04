package main

import (
	"BackAPI/internal/API"
	"BackAPI/internal/database"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	err := database.OpenBD("user=postgres password=ВСТАВЬ_СВОЙ_СУКА dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = database.CloseDB()
		if err != nil {
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/post_client", API.PostClient)
	http.HandleFunc("/post_trainer", API.PostTrainer)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}
