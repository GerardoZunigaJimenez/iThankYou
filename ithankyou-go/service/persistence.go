package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var err error

func init() {
	createConnection()
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("problem to ping the database server", err)
		return
	}
}

func createConnection() {
	db, err = sql.Open("mysql", "service:ironmansucks@tcp(localhost:3306)/iThankYou?charset=utf8")
	if err != nil {
		log.Println("problem to open a database connection", err)
		return
	}
}

func FetchUserByEmail(email string) User {
	// data to be used in query
	var u User

	createConnection()
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM users where email = '` + email + `';`)
	if err != nil {
		log.Println("problem to execute the query", err)
		return u
	}
	defer rows.Close()

	// query
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.LastName, &u.FirstName, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			log.Println("problem to parse the rows", err)
			return u
		}
	}
	return u
}
