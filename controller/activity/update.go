package activity

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
		payload     entities.ActivityPayload
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

	data, err := repositories.NewActivityRepositories().FindById(id)
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

	if payload.Email == "" {
		payload.Email = data.Email
	}

	time := time.Now().Local()
	objectActivity := entities.Activity{
		ID:        data.ID,
		Title:     payload.Title,
		Email:     payload.Email,
		UpdatedAt: time.String(),
		CreatedAt: data.CreatedAt,
	}

	res, err := repositories.NewActivityRepositories().UpdateActivity(&objectActivity)
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
