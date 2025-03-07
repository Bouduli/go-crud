package handlers

import (
	"encoding/json"
	"fmt"
	"go-crud/internal/types"
	utils "go-crud/internal/utils"
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

// Struct containing the different handler methods
type TodoHandler struct{}

// Index
//
//	GET / route
func (h *TodoHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	jsonWriter.Encode(TODOs)
}

// Show
//
// GET /:id
func (h *TodoHandler) Show(w http.ResponseWriter, r *http.Request) {
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

		utils.Response{W: w}.Status(http.StatusNotFound).ProblemJson(types.ProblemJson{
			Type:   "example.com/not-found",
			Status: http.StatusNotFound,
			Title:  "not found",
			Detail: fmt.Sprintf("A TODO with id \"%s\" wasnt found", param),
		})
	}

}

// Create
//
// POST /
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {

	//limit the size of request-body (trying to parse larger files result in error)
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	var TODO types.Note
	if err := json.NewDecoder(r.Body).Decode(&TODO); err != nil {

		utils.Response{W: w}.ErrorMap(err)
		return
	}

	TODO.Id = uuid.NewString()

	valid := TODO.Validate()
	if !valid.Ok {
		utils.Response{W: w}.Status(http.StatusBadRequest).ProblemJson(types.ProblemJson{
			Type:   "example.com/bad-request",
			Status: http.StatusBadRequest,
			Title:  "Invalid TODO",
			Detail: fmt.Sprintf("TODO is missing required properties: %v", valid.ErrorFields),
		})
		return
	}

	TODOs = append(TODOs, TODO)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(TODO)
}
