package chanell

import (
	"fmt"
	"time"
)

func Read(ch chan string, inchen string) {
	time.Sleep(2 * time.Second)
	fmt.Println("внутри", inchen)
	ch <- inchen
}
