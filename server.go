package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var home = os.Getenv("HOME")

	fs := http.FileServer(http.Dir(home + "/public/"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":4747", nil)
}
