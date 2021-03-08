/*
 * Fast Counter
 *
 * A fast incrementing counter API
 *
 * API version: 1.0.0
 * Contact: daren@darenliang.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() http.Handler {
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

	handler := cors.Default().Handler(router)
	return handler
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fast Counter API")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"HitKey",
		strings.ToUpper("Get"),
		"/hit/{key}",
		HitKey,
	},
}
