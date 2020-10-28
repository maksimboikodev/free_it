package main

import (
	"fmt"
	"os"

	"github.com/maksimboikodev/test/pkg/urlshortener"

	"github.com/kelseyhightower/envconfig"
	"github.com/maksimboikodev/test/pkg/datatypes"
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

	var d float32
	fmt.Print("Введите температуру: ")
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
	Storage := urlshortener.NewURLStore()
	tempHistoryHandler.Info("Key New:  ", Storage.Set("3_Url", "abrakada.com"))
	tempHistoryHandler.Info("Key New:  ", Storage.Set("5_Url", "ab.com"))
	tempHistoryHandler.Info("ALL MAP:  ", Storage)
	tempHistoryHandler.Info("Url: ", Storage.Get("12_URL"))
	tempHistoryHandler.Info("Url: ", Storage.Get("5_Url"))
}
