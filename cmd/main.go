package main

import (
	"os"

	"github.com/maksimboikodev/test/pkg/csvwork"

	"github.com/kelseyhightower/envconfig"
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

	/*var d float64
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
	storage := urlshortener.NewURLStore()
	tempHistoryHandler.Info("Key New:  ", storage.Set("3_Url", "abrakada.com"))
	tempHistoryHandler.Info("Key New:  ", storage.Set("5_Url", "ab.com"))
	tempHistoryHandler.Info("ALL MAP:  ", storage)
	tempHistoryHandler.Info("Url: ", storage.Get("12_URL"))
	tempHistoryHandler.Info("Url: ", storage.Get("5_Url"))

	if _, error := urlshortener.Sqrt(d); error != nil {
		tempHistoryHandler.Info("err: ", error)
	}

	mess := "message"
	tempHistoryHandler.Info("started ", mess)
	ch := make(chan string)
	go chanell.DoSomething(ch, tempHistoryHandler)
	ch <- mess
	tempHistoryHandler.Info("push chanell  ", mess)*/
	csvwork.Readcsv()
}
