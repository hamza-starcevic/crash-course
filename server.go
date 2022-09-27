package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = ":8080"

func main() {

	//* Create a new router
	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))

	//* Handler for the home route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	//* Handler for the /posts route
	router.HandleFunc("/posts", GetPosts).Methods("GET")

	router.HandleFunc("/posts", AddPost).Methods("POST")
	//* Start the server
	log.Print("Server is running on port ", port)
	http.ListenAndServe(port, router)
}
