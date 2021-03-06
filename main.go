/*
 * Fast Counter
 *
 * A fast incrementing counter API
 *
 * API version: 1.0.0
 * Contact: daren@darenliang.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
	"os"

	sw "github.com/darenliang/fastcounter/go"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	log.Printf("Fast Counter API started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(getPort(), router))
}
