package cmd

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Start(d Definition) {
	router := chi.NewRouter()
	Routes(router, d)

	printRoutes()

	http.ListenAndServe(":8080", router)
}

func Routes(router *chi.Mux, d Definition) {

	router.Get("/transactions/{user_id}", d.TransactionHandler.Get)
	router.Post("/transactions/{user_id}/{receptor_id}/{amount}", d.TransactionHandler.Create)
	router.Get("/user/{user_id}", d.UserHandler.Get)

}

func printRoutes() {
	fmt.Printf("Routes: \n")
	fmt.Printf(" %s\t%v\t\n", "/user/{user_id}", "[GET]  [check balance]")
	fmt.Printf(" %s\t%v\t\n", "/transactions/{user_id}/{receptor_id}/{amount}", "[POST] [create a transaction]")
	fmt.Printf(" %s\t%v\t\n", "/transactions/{user_id}", "[GET]  [check user's transactions]")
}
