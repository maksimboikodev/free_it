package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	first_name string
	last_name  string
	age        int
}
type PersonRepository struct {
	database *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "free"
)

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

func ConnectDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	return db, err
}

func (repository *PersonRepository) FindAll() ([]User, error) {
	rows, err := repository.database.Query("SELECT * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.first_name, &p.last_name, &p.age)
		if err != nil {
			return nil, err
			continue
		}
		users = append(users, p)
	}
	return users, err
}

func (repository *PersonRepository) AddRecord() error {
	p := User{}
	_, err := repository.database.Exec("INSERT INTO users VALUES($1, $2, $3)", p.first_name, p.last_name, p.age)
	return err
}
