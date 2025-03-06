package routers

import (
	"encoding/json"
	"fmt"
	"go-crud/internal/types"
	"go-crud/internal/utils/router"
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
var TodoRouter router.Router

// init function allows us to initialize variables
func init() {
	TodoRouter = router.NewRouter("/todo")

	TodoRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonWriter := json.NewEncoder(w)

		jsonWriter.Encode(TODOs)
	})
	TodoRouter.GET("/{id}", func(w http.ResponseWriter, r *http.Request) {
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
}
