package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"main/helper"
	"main/models"
	"main/repository"
	"main/services"
	"net/http"
	"strconv"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	request := models.UserRequest{}
	helper.ReadFromRequestBody(r, &request)

	services.SignUp(request)

	res := models.Response{
		Status:  200,
		Message: "User created successfully",
	}
	helper.WriteToResponseBody(w, res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	request := models.UserRequest{}
	helper.ReadFromRequestBody(r, &request)

	valid, err := repository.CheckUser(request.Username, request.Password)
	if err != nil {
		response := map[string]string{
			"status":  "error",
			"message": "An error occurred or username and password is incorrect",
		}
		jsonResponse, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write(jsonResponse)
		if err != nil {
			return
		}
		return
	}

	res := models.Response{
		Status:  200,
		Message: "User logged in successfully",
		Token:   valid,
	}
	helper.WriteToResponseBody(w, res)

}

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
	id, _ := strconv.Atoi(params["id"])
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
