package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-postgres-menu/helper"
	"go-postgres-menu/middleware"
	"go-postgres-menu/models"
	"log"
)

func Insert(master models.Master) int64 {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `INSERT INTO master (name) VALUES ($1) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, master.Name).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return id
}

func InsertCategory(category models.Category) int64 {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `INSERT INTO category (master_id, name_category) VALUES ($1, $2) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, category.MasterId, category.NameCategory).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return id
}

func InsertBahan(bahan models.Bahan) {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	_, err := db.Exec("INSERT INTO bahan (master_id, category_id, name_bahan) VALUES ($1, $2, $3)", bahan.MasterId, bahan.CategoryId, bahan.NameBahan)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMenu(id int64) (models.ResponseResult, error) {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	stmt := `SELECT m.name, c.name_category, json_agg(b.name_bahan) AS bahan
			 FROM master m
			 LEFT JOIN category c ON m.id = c.master_id
			 LEFT JOIN bahan b ON c.id = b.category_id
			 WHERE m.id = $1
			 GROUP BY m.name, c.name_category`

	rows, err := db.Query(stmt, id)
	if err != nil {
		helper.Panic(err)
	}
	defer rows.Close()

	var recipes []models.ResponseResult
	for rows.Next() {
		var name, category string
		var ingredients []byte
		err := rows.Scan(&name, &category, &ingredients)
		if err != nil {
			log.Fatal(err)
		}
		var ingredientList []string
		err = json.Unmarshal(ingredients, &ingredientList)
		if err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, models.ResponseResult{
			Name:         name,
			NameCategory: category,
			NameBahan:    ingredientList,
		})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	var finalResult models.ResponseResult
	for _, recipe := range recipes {
		finalResult.Name = recipe.Name
		finalResult.NameCategory = recipe.NameCategory
		finalResult.NameBahan = recipe.NameBahan
	}
	return finalResult, err
}

func GetAllMenu(name string) ([]models.ResponseResult, error) {
	conn, err := middleware.CreateConnectionPgx()
	if err != nil {
		helper.Panic(err)
	}
	defer middleware.CloseConnectionPgx(conn)

	query := `
        SELECT m.name, c.name_category, json_agg(b.name_bahan) AS bahan
        FROM master m
        LEFT JOIN category c ON m.id = c.master_id
        LEFT JOIN bahan b ON c.id = b.category_id
    `

	if name != "" {
		query += `WHERE m.name ILIKE '%` + name + `%' or c.name_category ILIKE '%` + name + `%' or b.name_bahan ILIKE '%` + name + `%'`
	}
	query += `GROUP BY m.name, c.name_category`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		helper.Panic(err)
	}
	defer rows.Close()

	var results []models.ResponseResult
	for rows.Next() {
		var name, category string
		var bahan []string
		err := rows.Scan(&name, &category, &bahan)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			helper.Panic(err)
		}
		results = append(results, models.ResponseResult{
			Name:         name,
			NameCategory: category,
			NameBahan:    bahan,
		})
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Failed to iterate over rows:", err)
		helper.Panic(err)
	}

	return results, err
}

func UpdateMenu(id int64, menu models.Master) int64 {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `UPDATE master SET name=$2 WHERE id=$1 RETURNING id`
	res, err := db.Exec(sqlStatement, id, menu.Name)
	if err != nil {
		helper.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprintf("Rows affected: %v", rowsAffected)

	return id
}

func UpdateCategory(id int64, category models.Category) int64 {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `UPDATE category SET name_category=$2 WHERE master_id=$1 RETURNING category.id`
	res, err := db.Exec(sqlStatement, id, category.NameCategory)
	if err != nil {
		helper.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		helper.Panic(err)
	}

	var categoryId int64
	err = db.QueryRow("SELECT id FROM category WHERE master_id = $1", id).Scan(&categoryId)
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprintf("Rows affected: %v", rowsAffected)
	return categoryId
}

func UpdateBahan(id int64, bahan models.Bahan) int64 {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `UPDATE bahan SET name_bahan=$2 WHERE master_id=$1`
	res, err := db.Exec(sqlStatement, id, bahan.NameBahan)
	if err != nil {
		helper.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprintf("Rows affected: %v", rowsAffected)

	return id
}

func DeleteBahanByMasterID(masterID int64) error {
	conn, err := middleware.CreateConnectionPgx()
	if err != nil {
		helper.Panic(err)
	}
	defer middleware.CloseConnectionPgx(conn)

	_, err = conn.Exec(context.Background(), `DELETE FROM bahan WHERE master_id = $1`, masterID)
	if err != nil {
		fmt.Println("Failed to delete bahan records:", err)
		return err
	}

	return nil
}

func CreateBahan(request models.Bahan) error {
	conn, err := middleware.CreateConnectionPgx()
	if err != nil {
		helper.Panic(err)
	}
	defer middleware.CloseConnectionPgx(conn)

	_, err = conn.Exec(context.Background(), `INSERT INTO bahan (name_bahan, category_id, master_id)VALUES ($1, $2, $3)`, request.NameBahan, request.CategoryId, request.MasterId)
	if err != nil {
		fmt.Println("Failed to create bahan record:", err)
		return err
	}

	return nil
}

func DeleteMenu(id int64) error {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `DELETE FROM master WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		helper.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprintf("Rows affected: %v", rowsAffected)

	return err
}

func DeleteCategory(id int64) error {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `DELETE FROM category WHERE master_id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		helper.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprintf("Rows affected: %v", rowsAffected)

	return err
}

func DeleteBahan(id int64) error {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `DELETE FROM bahan WHERE master_id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		helper.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		helper.Panic(err)
	}
	fmt.Sprintf("Rows affected: %v", rowsAffected)

	return err
}

func CheckMenu(id int64) (models.Master, error) {
	db := middleware.CreateConnection()
	defer middleware.CloseConnection(db)

	sqlStatement := `SELECT id, name FROM master WHERE id=$1`
	rows, err := db.QueryContext(context.Background(), sqlStatement, id)
	helper.Panic(err)
	defer rows.Close()

	master := models.Master{}
	if rows.Next() {
		err := rows.Scan(&master.Id, &master.Name)
		helper.Panic(err)
		return master, nil
	} else {
		return master, errors.New("menu is not found")
	}
}
