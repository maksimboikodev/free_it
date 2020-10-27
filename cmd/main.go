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
	baseurl := urlshortener.NewURLStore()
	baseurl.Put("abrakadabra.com")
	baseurl.Put("google.com")
	baseurl.Put("dabra.com")
	baseurl.Put("abrakadabra.com")
	baseurl.Put("kadabra.com")
	baseurl.Put("kadabra.com")
	baseurl.Put("abra.com")
	fmt.Println("number of keys: ", baseurl.Count())
	fmt.Println("Key New:  ", baseurl.Set("3_Url", "abrakada.com"))
	fmt.Println("Key New:  ", baseurl.Set("3_Url", "ab.com"))
	fmt.Println("ALL MAP:  ", baseurl)
	fmt.Println("Url: ", baseurl.Get("12_URL"))
	fmt.Println("Url: ", baseurl.Get("5_URL"))
}
