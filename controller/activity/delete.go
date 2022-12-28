package activity

import (
	"encoding/json"
	"go-learn/library/response"
	"go-learn/repositories"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
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

	data, err := repositories.NewActivityRepositories().FindById(id)
	if err != nil {
		response := *errResponse.WithError("ID not Found")
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	err = repositories.NewActivityRepositories().DeleteActivity(data.ID)
	if err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	object, err := json.Marshal(succResponse)
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
