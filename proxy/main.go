package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	var router *chi.Mux = SetupRouter()

	http.ListenAndServe(":8080", router)
}
