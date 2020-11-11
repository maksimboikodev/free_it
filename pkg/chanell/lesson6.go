package chanell

import (
	"time"

	"github.com/maksimboikodev/test/pkg/usecases"
)

func DoSomething(ch chan string, h *usecases.HistoryHander) {
	time.Sleep(3 * time.Second)
	h.Info("pull chanell ", <-ch)
}
