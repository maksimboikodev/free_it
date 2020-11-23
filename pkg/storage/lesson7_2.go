package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Users struct {
	first_name string
	last_name  string
	age        int
}

func Connect() {
	connStr := "user=maksim password=pass  dbname=freeit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connect")
	}
	fmt.Println("Open connections", db.Stats().OpenConnections)
	defer db.Close()

	rows, err := db.Query("SELECT * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []Users{}

	for rows.Next() {
		p := Users{}
		err := rows.Scan(&p.first_name, &p.last_name, &p.age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	for _, p := range users {
		fmt.Println(p.first_name, p.last_name, p.age)
	}
	_, err = db.Exec("insert into users (first_name, last_name, age) values ('Maksim', 'Maksim')",
		100)
	if err != nil {
		panic(err)
	}
}
