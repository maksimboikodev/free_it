package chanel

import (
	"fmt"
	"time"
)

func read(c chan string) {
	time.Sleep(2 * time.Second)
	b := <-c
	fmt.Println("The value at the exit from the channel ", b)
	//usecases.HistoryHander.Info(b)
}
