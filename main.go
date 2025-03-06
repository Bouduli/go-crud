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
		Id:     uuid.New().String(),
		Title:  "Title1",
		Text:   "message",
		Author: uuid.New().String(),
	},
	{
		Id:     uuid.New().String(),
		Title:  "Title2",
		Text:   "message",
		Author: uuid.New().String(),
	},
	{
		Id:     uuid.New().String(),
		Title:  "Title3",
		Text:   "message",
		Author: uuid.New().String(),
	}}

func main() {

	mux := http.NewServeMux()

	todoRouter := router.NewRouter("/todo")
	todoRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonWriter := json.NewEncoder(w)

		jsonWriter.Encode(TODOs)
	})
	todoRouter.GET("/{id}", func(w http.ResponseWriter, r *http.Request) {
		param := r.PathValue("id")

		var TODO *types.Note
		for i := range TODOs {
			if TODOs[i].Id == param {
				TODO = &TODOs[i]
			}
		}

		if TODO != nil {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TODO)
		} else {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(types.ProblemJson{
				Type:   "example.com/not-found",
				Status: http.StatusNotFound,
				Title:  "not found",
				Detail: fmt.Sprintf("A TODO with id \"%s\" wasnt found", param),
			})
		}

	})
	todoRouter.GET("/what", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "something else")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "index, html coming soon!")

	})

	mux.Handle("/todo/", todoRouter.Serve())

	fmt.Println("server started at localhost:80")
	http.ListenAndServe("localhost:80", mux)
}
