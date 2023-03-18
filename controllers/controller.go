package controllers

import (
	"github.com/gorilla/mux"
	"go-postgres-menu/helper"
	"go-postgres-menu/models"
	"go-postgres-menu/services"
	"log"
	"net/http"
	"strconv"
)

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	helper.ReadFromRequestBody(r, &request)

	services.Create(request)

	res := models.Response{
		Status:  200,
		Message: "Menu created successfully",
		Data:    request,
	}
	helper.WriteToResponseBody(w, res)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.Panic(err)
	}

	master, err := services.GetMenu(int64(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := models.Response{
			Status:  http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    err.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
	} else {
		res := models.Response{
			Status:  200,
			Message: "successfully",
			Data:    master,
		}
		helper.WriteToResponseBody(w, res)
	}
}

func GetAllMenu(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	masters := services.GetAll(query)
	res := models.Response{
		Status:  200,
		Message: "successfully",
		Data:    masters,
	}
	helper.WriteToResponseBody(w, res)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	menuRequest := models.Request{}
	helper.ReadFromRequestBody(r, &menuRequest)
	_, err = services.UpdateMenu(int64(id), menuRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := models.Response{
			Status:  http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    err.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
	} else {
		res := models.Response{
			ID:      int64(id),
			Status:  200,
			Message: "Menu updated successfully",
			Data:    nil,
		}

		helper.WriteToResponseBody(w, res)
	}
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}
	_, err = services.DeleteMenu(int64(id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := models.Response{
			Status:  http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    err.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
	} else {
		res := models.Response{
			ID:      int64(id),
			Status:  200,
			Message: "Menu deleted successfully",
			Data:    nil,
		}

		helper.WriteToResponseBody(w, res)
	}
}
