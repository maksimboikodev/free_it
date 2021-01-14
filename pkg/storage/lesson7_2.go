package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	First_name string
	Last_name  string
	Age        int
}
type PersonRepository struct {
	database *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "freeit"
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
		err := rows.Scan(&p.First_name, &p.Last_name, &p.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, p)
	}
	return users, err
}

func (repository *PersonRepository) AddRecord(p *User) error {
	_, err := repository.database.Exec("INSERT INTO users VALUES($1, $2, $3)", p.First_name, p.Last_name, p.Age)
	return err
}
