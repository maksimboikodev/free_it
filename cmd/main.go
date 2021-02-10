package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/maksimboikodev/test/pkg/gorilla"
	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	LogFile string `envconfig:"LOG_FILE"`
}

func main() {
	var eConf EnvConfig
	envconfig.Process("", &eConf)

	file, _ := os.OpenFile(eConf.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	var log = logrus.New()
	log.Out = file

	/*tempHistoryHandler := &usecases.HistoryHander{
		Logger:  log,
		Samples: []usecases.Fahrenheit{},
	}

	usecases.CheckAndSave(tempHistoryHandler, usecases.Celsius(55), 2)

	datatypes.List(tempHistoryHandler)
	datatypes.TurnFIFO(tempHistoryHandler)
	str := "Golang"
	datatypes.Reverse(str, tempHistoryHandler)
	datatypes.ReverseString(str, tempHistoryHandler)
	repository := urlshortener.NewURLStore()
	tempHistoryHandler.Info("Key New:  ", repository.Set("5_Url", "ab.com"))
	tempHistoryHandler.Info("Url: ", repository.Get("12_URL"))

	if _, error := urlshortener.Sqrt(25); error != nil {
		tempHistoryHandler.Info("err: ", error)
	}

	mess := "message"
	tempHistoryHandler.Info("started ", mess)
	ch := make(chan string)
	go chanell.DoSomething(ch, tempHistoryHandler)
	ch <- mess
	tempHistoryHandler.Info("push chanell  ", mess)

	db, err := storage.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	h := storage.NewPersonRepository(db)

	p := storage.User{First_name: "CVack", Last_name: "jack", Age: 30}
	err = h.AddRecord(&p)
	if err != nil {
		fmt.Println(err)
	}
	sel, err := h.FindAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(sel)
	*/

	router := mux.NewRouter()
	gorilla.Students = append(gorilla.Students, gorilla.Freeit{ID: "1", First_name: "Maksim", Last_name: "Boiko", Position: "Student"})
	gorilla.Students = append(gorilla.Students, gorilla.Freeit{ID: "2", First_name: "Vladimir", Last_name: "Vladimir", Position: "Student"})
	gorilla.Students = append(gorilla.Students, gorilla.Freeit{ID: "3", First_name: "Ekaterina", Last_name: "Shemerey", Position: "Mentor"})
	router.HandleFunc("/students", gorilla.GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", gorilla.GetStudentV1).Methods("GET").Headers("Version", "v1")
	router.HandleFunc("/students/{id}", gorilla.GetStudentV2).Methods("GET").Headers("Version", "v2")
	router.HandleFunc("/students", gorilla.CreateStudent).Methods("POST")
	/*{ POST Request
		"First_name": "Alex",
		"Last_name": "Alex",
		"Position": "Alex"
	}*/
	router.HandleFunc("/students/{id}", gorilla.DeleteStudent).Methods("DELETE")
	/*{ Delete request
	    "id":"727887",
	    "First_name": "Alex",
	    "Last_name": "Alex",
	    "Position": "Alex"
	}*/
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
