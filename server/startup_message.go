package main

import (
	"fmt"
	"log"
	"net/http"
)

func displayStartupMessage() {
	fmt.Println("-------------------- FluxDB --------------------")
	fmt.Print("Thank you for using FluxDB. Made with love by\nthe Golang community.\n\n")
	fmt.Printf("Running on port %d\n", defaultPort)
	fmt.Println("------------------------------------------------")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil))
}
