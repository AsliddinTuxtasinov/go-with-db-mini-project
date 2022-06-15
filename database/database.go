package database

import (
	"database/sql"
	"fmt"
	"go-with-db/user"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Open MySql Database
func OpenDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@/")
	return
}

// Create database IF NOT EXISTS
func CreateAndUseDB(db *sql.DB, name string) {
	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v;\n", name))
	db.Exec(fmt.Sprintf("USE %v;\n", name))
}

// Create database table IF NOT EXISTS
func CreateTableDB(db *sql.DB, tableName string) (res sql.Result, err error) {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v(
		id int(100) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		username varchar(45) NOT NULL,
		password varchar(45) NOT NULL,
		isActive tinyint(1) DEFAULT NULL,
		UNIQUE (id)
	);`, tableName)
	res, err = db.Exec(query)
	return
}

// Insert data into database
func InserDataToDB(db *sql.DB, username, password string, isActive int) (sql.Result, error) {

	res, err := db.Exec(
		"INSERT INTO users(username, password, isActive) VALUES (?, ?, ?);", username, password, isActive)
	return res, err
}

// Insert multiple data into database
func InserMultipleDataToDB(db *sql.DB, data []user.User) (err error) {
	for _, v := range data {
		username, password, isActive := v.GetFieldsFromUser()
		_, err := InserDataToDB(db, username, password, isActive)

		if err != nil {
			break
		}
	}
	return
}

// Select data from database
func SelectDataFromToDB(db *sql.DB, tableName string) (row *sql.Rows, err error) {
	row, err = db.Query(fmt.Sprintf("SELECT * FROM %v;\n", tableName))
	return

}

// Select data from database by username
func SelectDataByUsernameFromToDB(db *sql.DB, tableName, username string) (row *sql.Rows, err error) {
	row, err = db.Query(fmt.Sprintf("SELECT * FROM %v WHERE username = \"%s\";\n", tableName, username))
	return

}

// Select data from database with limit
func SelectDataWithLimitFromToDB(db *sql.DB, tableName string, limit int) (row *sql.Rows, err error) {
	row, err = db.Query(fmt.Sprintf("SELECT * FROM %s LIMIT %d;\n", tableName, limit))
	return

}

// Delete data from database by username
func DeleteDataByUsernameFromToDB(db *sql.DB, tableName, username string) (row *sql.Rows, err error) {
	row, err = db.Query(fmt.Sprintf("DELETE FROM %s  WHERE username = \"%s\";\n", tableName, username))
	return

}

// Update data from database by username
func UpdateDataByUsernameFromToDB(db *sql.DB, tableName, username, column, value string) (row *sql.Rows, err error) {
	row, err = db.Query(fmt.Sprintf("UPDATE %s SET %s=\"%s\" WHERE username = \"%s\";\n", tableName, column, value, username))
	return

}

// check error
func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
