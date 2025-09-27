package main

import (
	"fmt"
	"net/http"

	"github.com/kodega2016/booking-app/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application running on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
