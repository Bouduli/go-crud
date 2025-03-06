package main

import (
	"fmt"
	"go-crud/internal/routers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "index, html coming soon!")

	})

	mux.Handle("/todo/", routers.TodoRouter.Serve())

	fmt.Println("server started at localhost:80")
	http.ListenAndServe("localhost:80", mux)
}
