package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erdincmutlu/go-react-todo/server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Staring the server on port 9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}
