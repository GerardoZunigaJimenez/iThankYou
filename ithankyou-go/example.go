package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"iThankYou/ithankyou-go/service"
	"io"
	"net/http"
)

var db *sql.DB
var err error

//This class is an example to test the mysql connection.  Please delete it when the project have been finished
func main() {
	//user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "service:ironmansucks@tcp(localhost:3306)/iThankYou?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	//http.HandleFunc("/create", create)
	//http.HandleFunc("/insert", insert)
	//http.HandleFunc("/read", read)
	//http.HandleFunc("/update", update)
	//http.HandleFunc("/delete", del)
	//http.HandleFunc("/drop", drop)

	err := http.ListenAndServe(":8080", nil)

	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM users;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var u service.User
	var su []service.User
	// data to be used in query
	var s, name string
	// query
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.LastName, &u.FirstName, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		check(err)
		su = append(su, u)
		s += name + "\n"
	}

	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(su)
	w.Write([]byte(payload))
}

func create(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
