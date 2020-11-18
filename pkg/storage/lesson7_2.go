package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Users struct {
	first_name string
	last_name  string
	age        int
}

func newrec() Users {
	a := Users{"maksim", "boiko", 31}
	return a
}

func Connect() {
	connStr := "user=maksim password=pass dbname=freeit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Open connections", db.Stats().OpenConnections)
	err = db.Ping()
	if err != nil {
		fmt.Println("Errrrooooorrrr connect")
	}
	fmt.Println("Open connections", db.Stats().OpenConnections)
	defer db.Close()

	rows, err := db.Query("SELECT first name,last name,age FROM freeit")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var firstname string
		var lastname string
		var age int
		err = rows.Scan(&firstname, &lastname, &age)
		fmt.Println("Stroka", firstname, lastname, age)
	}
	rows.Close()
}
