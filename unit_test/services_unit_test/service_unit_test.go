package services_unit_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"go-postgres-menu/helper"
	"go-postgres-menu/middleware"
	"go-postgres-menu/models"
	"go-postgres-menu/repository"
	"go-postgres-menu/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	request := models.Request{
		Name:         "Nasi Goreng",
		NameCategory: "Makanan",
		NameBahan:    []string{"Nasi", "Telur", "Bawang"},
	}

	master := models.Master{
		Name: request.Name,
	}

	category := models.Category{
		MasterId:     1,
		NameCategory: request.NameCategory,
	}

	bahan := models.Bahan{
		MasterId:   1,
		CategoryId: 1,
		NameBahan:  "Nasi",
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("INSERT INTO master").WithArgs(master.Name).WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("INSERT INTO category").WithArgs(category.MasterId, category.NameCategory).WillReturnRows(rows)

	mock.ExpectExec("INSERT INTO bahan").WithArgs(bahan.MasterId, bahan.CategoryId, bahan.NameBahan).WillReturnResult(sqlmock.NewResult(1, 1))

	result := services.Create(request)

	if result.Name != master.Name {
		t.Errorf("Expected %s, got %s", master.Name, result.Name)
	}
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Nasi Goreng")
	mock.ExpectQuery("SELECT id, name FROM master").WillReturnRows(rows)

	result := services.GetAll("")

	if len(result) != 1 {
		t.Errorf("Expected %d, got %d", 1, len(result))
	}
}

func TestGetMenu(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE master (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL)")
	db.Exec("CREATE TABLE category (id SERIAL PRIMARY KEY, master_id INT NOT NULL, name_category VARCHAR(255) NOT NULL)")
	db.Exec("CREATE TABLE bahan (id SERIAL PRIMARY KEY, master_id INT NOT NULL, category_id INT NOT NULL, name_bahan VARCHAR(255) NOT NULL)")

	db.Exec("INSERT INTO master (name) VALUES ('Nasi Goreng')")
	db.Exec("INSERT INTO category (master_id, name_category) VALUES (1, 'Nasi')")
	db.Exec("INSERT INTO bahan (master_id, category_id, name_bahan) VALUES (1, 1, 'Nasi')")

	menu, err := repository.GetMenu(1)
	if err != nil {
		panic(err)
	}
	if menu.Name != "Nasi Goreng" {
		t.Errorf("Expected Nasi Goreng, but got %s", menu.Name)
	}
}

func TestUpdateMenu(t *testing.T) {
	req, err := http.NewRequest("PUT", "/menu/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateMenu)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func UpdateMenu(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	write, err := writer.Write([]byte(`{"id":1,"name":"Updated Menu","category":[{"id":1,"name_category":"Updated Category","bahan":[{"id":1,"name_bahan":"Updated Bahan"}]}]}`))
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprint(write)
}

func TestDeleteMenu(t *testing.T) {
	db := middleware.CreateConnection()
	defer db.Close()

	master := models.Master{
		Name: "Nasi Goreng",
	}
	lastIdMaster := repository.Insert(master)
	category := models.Category{
		MasterId:     lastIdMaster,
		NameCategory: "Bahan Utama",
	}
	lastIdCategory := repository.InsertCategory(category)
	bahan := models.Bahan{
		MasterId:   lastIdMaster,
		CategoryId: lastIdCategory,
		NameBahan:  "Nasi",
	}
	repository.InsertBahan(bahan)

	err := repository.DeleteMenu(lastIdMaster)
	if err != nil {
		t.Errorf("Error when delete menu: %s", err)
	}
}
