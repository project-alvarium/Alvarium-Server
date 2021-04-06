package main

import (
	"database-manager/api"
	"database-manager/collections"
	"database-manager/configuration"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	httpRouter := api.NewRouter()
	configuration.InitConfig()
	srv := &http.Server{
		Handler: httpRouter,
		Addr:    ":" + fmt.Sprint(configuration.Config.HTTPPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	collections.Database()
	log.Fatal(srv.ListenAndServe())
	log.Println("listening")
}
