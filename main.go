package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alfanherya/jakarsafx/jakarsafx-registrasi/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	r := router.Router()
	fmt.Println("Starting server on the port 8090 ............")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
