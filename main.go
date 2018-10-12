package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	/**
	* @api {GET} / get help about this block
	* @apiName Intro
	 */
	r.HandleFunc("/", Intro)
	

	loggingRouter := handlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", loggingRouter)
	fmt.Println("Starting server at 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
