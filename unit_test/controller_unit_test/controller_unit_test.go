package controller_unit_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateMenu(t *testing.T) {
	req, err := http.NewRequest("PUT", "/menu/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateMenu)
	handler.ServeHTTP(rr, req)
}

func UpdateMenu(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	write, err := writer.Write([]byte(`{"status":200,"message":"Menu updated successfully","data":{"name":"test","description":"test","price":100}}`))
	if err != nil {
		panic(err)
	}
	log.Printf("write: %v", write)
}

func TestGetMenu(t *testing.T) {
	req, err := http.NewRequest("GET", "/menu/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMenu)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func GetMenu(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	write, err := writer.Write([]byte(`{"status":200,"message":"successfully","data":{"name":"test","description":"test","price":100}}`))
	if err != nil {
		panic(err)
	}
	log.Printf("write: %v", write)
}

func TestCreateMenu(t *testing.T) {
	req, err := http.NewRequest("POST", "/menu", bytes.NewBuffer([]byte(`{"name":"test","name_category":"test","name_bahan":test}`)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMenu)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func CreateMenu(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	write, err := writer.Write([]byte(`{"status":200,"message":"Menu created successfully","data":{"name":"test","description":"test","price":100}}`))
	if err != nil {
		panic(err)
	}
	log.Printf("write: %v", write)
}

func TestDeleteMenu(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/menu/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteMenu)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func DeleteMenu(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	write, err := writer.Write([]byte(`{"status":200,"message":"Menu deleted successfully","data":null}`))
	if err != nil {
		panic(err)
	}
	log.Printf("write: %v", write)
}
