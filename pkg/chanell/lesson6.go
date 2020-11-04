package chanell

import (
	"fmt"
	"time"
)

func Read(c chan string) {
	time.Sleep(2 * time.Second)
	b := <-c
	fmt.Println("The value at the exit from the channel ", b)
}
