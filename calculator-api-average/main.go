package main

import (
	"crud-simple/helper"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("Start main")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	server := http.Server{Addr: "localhost:8080", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
