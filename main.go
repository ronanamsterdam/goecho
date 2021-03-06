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
	
	/**
	* @api {GET} /get echoes back what you send it in the query params
	* @apiName GetEcho
	* @apiParam (Query) {String} key1 value of key1
	* @apiParam (Query) {String} [key2] value of key2
	* @apiSuccess (Success_200) data the query params encoded as a JSON object
	* @apiError (400) error.message the error message
	 */
	r.HandleFunc("/get", GetEcho).Methods("GET")
	/**
	* @api {POST} /post echoes back what you send it in the request body
	* @apiName PostEcho
	* @apiParam (Body) {String} key1 value of key1
	* @apiParam (Body) {String} [key2] value of key2
	* @apiSuccess (Success_200) data the request body encoded as a JSON object
	* @apiError (400) error.message the error message
	 */
	r.HandleFunc("/post", PostEcho).Methods("POST")
	
	loggingRouter := handlers.LoggingHandler(os.Stdout, r)

	http.Handle("/", loggingRouter)
	fmt.Println("Starting server at 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
