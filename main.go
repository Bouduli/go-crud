package main

import (
	"encoding/json"
	"fmt"
	"go-crud/internal/router"
	"go-crud/internal/types"
	"net/http"

	"github.com/google/uuid"
)

var TODOs = []types.Note{
	{
		Title:  "Title1",
		Text:   "message",
		Author: uuid.New().String(),
	},
	{
		Title:  "Title2",
		Text:   "message",
		Author: uuid.New().String(),
	},
	{
		Title:  "Title3",
		Text:   "message",
		Author: uuid.New().String(),
	}}

func main() {

	mux := http.NewServeMux()

	todoRouter := router.NewRouter("/todo")
	todoRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "todo index")
	})
	todoRouter.GET("/what", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "something else")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		jsonWriter := json.NewEncoder(w)

		if err := jsonWriter.Encode(TODOs); err != nil {
			http.Error(w, "unable to stringify stuff", http.StatusInternalServerError)
		}

	})

	mux.Handle("/todo/", todoRouter.Serve())

	fmt.Println("server started at localhost:80")
	http.ListenAndServe("localhost:80", mux)
}
