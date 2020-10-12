package main

import (
	"fmt"
	"os"

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

	/*var d float32
	fmt.Print("Введите температуру: ")
	fmt.Scan(&d)
	fmt.Println(d)*/

	tempHistoryHandler := &usecases.HistoryHander{
		Logger:  log,
		Samples: []usecases.Fahrenheit{},
	}

	usecases.CheckAndSave(tempHistoryHandler, usecases.Celsius( /*d*/ 10), 2)

	/*OcheredHandler := &datatypes.LoggerDz{
		Samples: datatypes.OcheredFIFO
		Logger: log
	}*/

	//datatypes.OcheredFIFO()
	datatypes.OdnSpis()
	datatypes.OcheredFIFO(tempHistoryHandler /*h * usecases.HistoryHander*/)
	str := "Golang"
	fmt.Println("", datatypes.Reverse(str))
}
