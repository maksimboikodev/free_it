package chanell

import (
	"time"
)

func Read(ch chan string) {
	time.Sleep(2 * time.Second)
	<-ch
}
