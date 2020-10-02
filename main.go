package main

import (
	"database-manager/api"
	"database-manager/configuration"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	httpRouter := api.NewRouter()
	httpRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		path, _ := route.GetPathTemplate()
		httpRouter.
			Methods(methods...).
			Path(path).
			Name(route.GetName()).
			Handler(route.GetHandler())
		return nil
	})
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(configuration.Config.HTTPPort), httpRouter))
}
