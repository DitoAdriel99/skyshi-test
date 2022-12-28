package todo

import (
	"encoding/json"
	"go-learn/entities"
	"go-learn/library/response"
	"go-learn/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		payload     entities.TodoPayload
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
	)
	rawID := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return

	}

	data, err := repositories.NewTodoRepositories().FindById(id)
	if err != nil {
		response := *errResponse.WithError("ID not Found")
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	if payload.Title == "" {
		payload.Title = data.Title
	}

	if payload.ActivityID == 0 {
		payload.ActivityID = data.ActivityID
	}

	if payload.Priority == "" {
		payload.Priority = data.Priority
	}

	time := time.Now().Local()
	update := entities.Todo{
		ID:         data.ID,
		Title:      payload.Title,
		ActivityID: payload.ActivityID,
		IsActive:   payload.IsActive,
		Priority:   payload.Priority,
		UpdatedAt:  time.String(),
		CreatedAt:  time.String(),
	}

	res, err := repositories.NewTodoRepositories().UpdateTodo(&update)
	if err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	response := *succResponse.WithData(&res)
	object, err := json.Marshal(response)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(object)
}
