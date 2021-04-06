package main

import (
	"fmt"
	"net/http"
	"os"

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

	tempHistoryHandler := &usecases.HistoryHander{
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
	tempHistoryHandler.Info("push chanell  ", mess)*/

	db, err := storage.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	h := storage.NewPersonRepository(db)*/

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

	gorilla.DB = gorilla.CreateConnection()
	r := gorilla.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
