package gorilla

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type Freeit struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Position string `json:"position"`
}
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

var DB *sql.DB

var config struct {
	PostgresDB struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		DBname   string `yaml:"DBname"`
	} `yaml:"PostgresDB"`
	AccessToken struct {
		Token string `yaml:"Token"`
	} `yaml:"Access-token"`
}

func ConfigDB(nameconfig string) string {
	configFile, err := ioutil.ReadFile(nameconfig)
	if err != nil {
		log.Fatalln("Failed to load config file")
	}

	err = yaml.Unmarshal([]byte(configFile), &config)
	if err != nil {
		log.Fatalf("cannot unmarshal config file: %v", err)
	}

	dbconf := config.PostgresDB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconf.Host, dbconf.Port, dbconf.User, dbconf.Password, dbconf.DBname)
	return psqlInfo
}

func CreateConnection(psqlInfo string) *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	hed := r.Header.Get("Access-token")
	if hed == config.AccessToken.Token {
		var user Freeit
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatalf("Unable to decode the request body.  %v", err)
		}
		insertID := insertUser(user)
		res := response{
			ID:      insertID,
			Message: "User created successfully",
		}
		json.NewEncoder(w).Encode(res)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	user, err := getUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := getAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	var user Freeit
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	updatedRows, err := updateUser(int64(id), user)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows, err := deleteUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func insertUser(user Freeit) int64 {
	sqlStatement := `INSERT INTO users (name, position, age) VALUES ($1, $2, $3) RETURNING userid`
	var id int64
	err := DB.QueryRow(sqlStatement, user.Name, user.Position, user.Age).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}

func getUser(id int64) (Freeit, error) {
	var user Freeit
	sqlStatement := `SELECT * FROM users WHERE userid=$1`
	row := DB.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Position)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func getAllUsers() ([]Freeit, error) {
	var users []Freeit
	sqlStatement := `SELECT * FROM users`
	rows, err := DB.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user Freeit
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Position)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)

	}
	return users, err
}

func updateUser(id int64, user Freeit) (int64, error) {
	sqlStatement := `UPDATE users SET name=$2, position=$3, age=$4 WHERE userid=$1`
	res, err := DB.Exec(sqlStatement, id, user.Name, user.Position, user.Age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected, err
}

func deleteUser(id int64) (int64, error) {
	sqlStatement := `DELETE FROM users WHERE userid=$1`
	res, err := DB.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, err
}

func Router() *mux.Router {
	router := mux.NewRouter()
	//curl -X GET http://localhost:8080/user/7
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	//curl -X GET http://localhost:8080/user
	router.HandleFunc("/user", GetAllUser).Methods("GET")
	//curl -X POST -H "Content-Type: application/json" -H "Access-token: Pass" -d "{\"name\":\"AAA\",\"age\":10,\"position\":\"A\"}" http://localhost:8080/user
	router.HandleFunc("/user", CreateUser).Methods("POST").Headers("Access-token", "Pass")
	//curl -X PUT -H "Content-Type: application/json" -d "{\"name\":\"AAA\",\"age\":10,\"position\":\"A\"}" http://localhost:8080/user/17
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	//curl -X DELETE http://localhost:8080/user/12
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	//curl -X GET http://localhost:8080/user/?id=12
	router.HandleFunc("/user/", QueryLine).Methods("GET")

	return router
}

//query response http://localhost:8080/api/queryline?id=1
func QueryLine(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	user, _ := getUser(int64(id))
	json.NewEncoder(w).Encode(user)
}
