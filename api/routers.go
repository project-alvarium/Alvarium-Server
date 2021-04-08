/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

// Route holds info of routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array
type Routes []Route

// NewRouter creates new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{
		"AddInsight",
		strings.ToUpper("Post"),
		"/api/insights",
		AddInsight,
	},

	Route{
		"GetInsightById",
		strings.ToUpper("Get"),
		"/api/insights/{id}",
		GetInsightByID,
	},
	Route{
		"addContent",
		strings.ToUpper("Post"),
		"/api/data",
		addContent,
	},
	Route{
		"getContent",
		strings.ToUpper("Get"),
		"/api/data/{id}",
		getContent,
	},
}
