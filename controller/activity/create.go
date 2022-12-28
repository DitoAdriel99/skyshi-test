package activity

import (
	"encoding/json"
	"fmt"
	"go-learn/entities"
	"go-learn/library/response"
	"go-learn/repositories"
	"net/http"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {
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

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return

	}

	time := time.Now().Local()

	data := entities.Activity{
		Title:     payload.Title,
		Email:     payload.Email,
		UpdatedAt: time.String(),
		CreatedAt: time.String(),
	}
	fmt.Println("dataa", data)

	err := repositories.NewActivityRepositories().Create(&data)
	if err != nil {
		fmt.Println("dimanaa")
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	response := *succResponse.WithData(data)
	object, err := json.Marshal(response)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(object)

}
