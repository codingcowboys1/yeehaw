package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

const (
	port = ":8080"
)

func init() {
	sayings := []string{
		"Dont squat with your spurs on!",
		"Never buy your salsa in New York city.",
		"Tenderfoots gotta do the coffee runs.",
		"Home is where you hang your ten gallon.",
	}
	rand.Seed(42)

	yeehaw := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v\n", sayings[rand.Intn(len(sayings))])
	}

	fmt.Printf("listening at http://localhost%v\n", port)
	http.HandleFunc("/", yeehaw)
	http.ListenAndServe(port, nil)
}

func main() {}
