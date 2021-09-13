package cmd

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Start(d Definition) {
	router := chi.NewRouter()
	Routes(router, d)

	http.ListenAndServe(":8080", router)
}

func Routes(router *chi.Mux, d Definition) {

	router.Get("/transactions/{user_id}", d.TransactionHandler.Get)
	router.Post("/transactions/{user_id}/{receptor_id}/{amount}", d.TransactionHandler.Create)
	router.Get("/user/{user_id}", d.UserHandler.Get)

}