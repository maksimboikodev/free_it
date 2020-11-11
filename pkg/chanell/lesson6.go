package chanell

import (
	"time"

	"github.com/maksimboikodev/test/pkg/usecases"
)

func Write(ch chan string, inmess string, h *usecases.HistoryHander) {
	h.Info("push chanell -", inmess)
	time.Sleep(2 * time.Second)
	ch <- inmess
}
