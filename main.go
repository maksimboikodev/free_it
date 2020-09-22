package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Celsius float32
type Fahrenheit float32
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
	c := Celsius(d)
	log.Info("температура по фаренгейту  ", toFahrenheit(c))

}

func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = Fahrenheit((t * 9 / 5) + 32)
	return temp
}
