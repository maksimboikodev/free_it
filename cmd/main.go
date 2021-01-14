package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/maksimboikodev/test/pkg/chanell"
	"github.com/maksimboikodev/test/pkg/csvwork"
	"github.com/maksimboikodev/test/pkg/datatypes"
	"github.com/maksimboikodev/test/pkg/storage"
	"github.com/maksimboikodev/test/pkg/urlshortener"
	"github.com/maksimboikodev/test/pkg/usecases"
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

	var d float64
	fmt.Print("enter temperature: ")
	fmt.Scan(&d)
	fmt.Println(d)

	tempHistoryHandler := &usecases.HistoryHander{
		Logger:  log,
		Samples: []usecases.Fahrenheit{},
	}

	usecases.CheckAndSave(tempHistoryHandler, usecases.Celsius(d), 2)

	datatypes.List(tempHistoryHandler)
	datatypes.TurnFIFO(tempHistoryHandler)
	str := "Golang"
	datatypes.Reverse(str, tempHistoryHandler)
	datatypes.ReverseString(str, tempHistoryHandler)
	repository := urlshortener.NewURLStore()
	tempHistoryHandler.Info("Key New:  ", repository.Set("5_Url", "ab.com"))
	tempHistoryHandler.Info("Url: ", repository.Get("12_URL"))

	if _, error := urlshortener.Sqrt(d); error != nil {
		tempHistoryHandler.Info("err: ", error)
	}

	mess := "message"
	tempHistoryHandler.Info("started ", mess)
	ch := make(chan string)
	go chanell.DoSomething(ch, tempHistoryHandler)
	ch <- mess
	tempHistoryHandler.Info("push chanell  ", mess)

	csv, err := csvwork.Readcsv()
	if err != nil {
		tempHistoryHandler.Info("err: ", err)
	}
	tempHistoryHandler.Info("result func Readcsv  ", csv)

	db, err := storage.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	h := storage.NewPersonRepository(db)

	sel, error := h.FindAll()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(sel)
	add := storage.User{"qwe", "rty", 123}
	err = h.AddRecord(add)
	if err != nil {
		fmt.Println(err)
	}
}
