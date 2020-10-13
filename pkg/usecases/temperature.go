package usecases

import "github.com/sirupsen/logrus"

type Celsius float32
type Fahrenheit float32

type HistoryHander struct {
	Samples []Fahrenheit
	*logrus.Logger
}

func CheckAndSave(h *HistoryHander, t Celsius, n int) {
	for i := 0; i < n; i++ {
		f := toFahrenheit(t) + Fahrenheit(i)
		h.Samples = append(h.Samples, f)
		h.Info("предыдущие изменения и текущая температура - ", h.Samples, f)
	}
}
func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = Fahrenheit((t * 9 / 5) + 32)
	return temp
}
