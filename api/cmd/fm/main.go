package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
