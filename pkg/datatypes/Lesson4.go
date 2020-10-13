package datatypes

import (
	"github.com/maksimboikodev/test/pkg/usecases"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TurnFIFO(h *usecases.HistoryHander) {
	var queue []string
	queue = append(queue, "Hello ")
	queue = append(queue, "world!")

	for len(queue) > 0 {
		h.Info(queue)
		queue = queue[1:]
	}
}

func List(h *usecases.HistoryHander) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	n3 := &ListNode{Val: 3}
	n2.Next = n3

	i := n1
	for i != nil {
		h.Info(i.Val)
		i = i.Next
	}
}

func Reverse(s string, h *usecases.HistoryHander) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	h.Info(string(runes))
	return string(runes)
}
func ReverseString(s string, h *usecases.HistoryHander) string {
	var newWord []rune
	for _, v := range s {
		newWord = append([]rune{v}, newWord...)
	}
	h.Info(string(newWord))
	return string(newWord)
}
