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
type PersonRep struct {
	database *sql.DB
}

func NewPersonRep(database *sql.DB) *PersonRep {
	return &PersonRep{database: database}
}

func ConnectDatabase() (*sql.DB, error) {
	connStr := "user=postgres password=pass  dbname=free sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connect")
	}
	fmt.Println("Open connections", db.Stats().OpenConnections)
	return db, err
}

func (repository *PersonRep) FindAll() {
	rows, err := repository.database.Query("SELECT * from users")
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
}

func (repository *PersonRep) AddRecord() {
	_, err := repository.database.Exec("insert into users (first_name, last_name, age) values ('M', 'M', $1)", 100)
	if err != nil {
		panic(err)
	}
}
