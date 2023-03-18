package repository_unit_test_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-postgres-menu/helper"
	"go-postgres-menu/middleware"
	"go-postgres-menu/models"
	"go-postgres-menu/repository"
	"testing"
)

// UNIT TEST FOR REPOSITORY PACKAGE
func TestInsert(t *testing.T) {
	db := middleware.CreateConnection()
	defer db.Close()

	_, err := db.Exec("DELETE FROM master")
	if err != nil {
		helper.Panic(err)
	}

	master := models.Master{
		Name: "Nasi Goreng",
	}
	id := repository.Insert(master)

	var name string
	err = db.QueryRow("SELECT name FROM master WHERE id = $1", id).Scan(&name)
	if err != nil {
		helper.Panic(err)
	}

	assert.Equal(t, "Nasi Goreng", name)
}

func TestInsertCategory(t *testing.T) {
	category := models.Category{
		MasterId:     1,
		NameCategory: "Makanan",
	}

	id := repository.InsertCategory(category)
	assert.NotEqual(t, id, 0)
}

func TestInsertBahan(t *testing.T) {
	bahan := models.Bahan{
		MasterId:   1,
		CategoryId: 1,
		NameBahan:  "Bawang Merah",
	}

	repository.InsertBahan(bahan)
}

func TestGetMenu(t *testing.T) {
	db := middleware.CreateConnection()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS master (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		helper.Panic(err)
	}

	_, err = db.Exec("INSERT INTO master (name) VALUES ($1)", "Nasi Goreng")
	if err != nil {
		helper.Panic(err)
	}
	result, err := repository.GetMenu(1)
	if err != nil {
		helper.Panic(err)
	}
	if result.Name != "Nasi Goreng" {
		t.Errorf("Expected Nasi Goreng, got %s", result.Name)
	}
}

func TestUpdateMenu(t *testing.T) {
	db := middleware.CreateConnection()
	defer db.Close()

	menu := models.Master{
		Name: "Menu 1",
	}
	id := repository.Insert(menu)
	menu = models.Master{
		Name: "Menu 2",
	}
	rowsAffected := repository.UpdateMenu(id, menu)
	assert.Equal(t, int64(1), rowsAffected)
}

func TestUpdateBahan(t *testing.T) {
	bahan := models.Bahan{
		MasterId:  1,
		NameBahan: "Bahan 1",
	}

	repository.InsertBahan(bahan)

	bahan.NameBahan = "Bahan 2"
	rowsAffected := repository.UpdateBahan(1, bahan) // id = 1 should be the id of the bahan we just inserted

	if rowsAffected != 1 {
		t.Errorf("Expected 1 row to be affected but was %d", rowsAffected)
	}
}

func TestDeleteMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM master WHERE master_id=$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.DeleteCategory(1)
	if err != nil {
		t.Errorf("error was not expected while updating record: %s", err)
	}
}

func TestDeleteCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM category WHERE master_id=$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.DeleteCategory(1)
	if err != nil {
		t.Errorf("error was not expected while updating record: %s", err)
	}
}

func TestDeleteBahan(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM bahan WHERE master_id=$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.DeleteBahan(1)
	if err != nil {
		t.Errorf("error was not expected while updating record: %s", err)
	}
}

func TestCheckMenu(t *testing.T) {
	db, mock, err := sqlmock.New()
	helper.Panic(err)
	defer middleware.CloseConnection(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "menu1")
	mock.ExpectQuery(`SELECT id, name FROM master WHERE id=$1`).WithArgs(1).WillReturnRows(rows)

	master, err := repository.CheckMenu(1)
	helper.Panic(err)

	assert.Equal(t, int64(1), master.Id)
	assert.Equal(t, "menu1", master.Name)
}

func TestRandomInt(t *testing.T) {
	result := helper.RandomInt()
	fmt.Println(result)
}
