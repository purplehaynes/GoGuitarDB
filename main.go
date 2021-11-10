package main

import (
	"GuitarDB/handlers"
	repo "GuitarDB/repository"
	"GuitarDB/service"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fn := "guitardb.json"

	ext := filepath.Ext(fn)
	if ext != ".json" {
		log.Fatalln("File extension is invalid")
	}

	r := repo.NewRepository(fn)

	svc := service.NewService(r)

	hdlr := handlers.NewGuitarHandler(svc)

	router := handlers.ConfigureRouter(hdlr)

	svr := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())
}