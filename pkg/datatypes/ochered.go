package datatypes

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func OcheredFIFO(h *LoggerDz) {
	var queue []string
	queue = append(queue, "Hello ")
	queue = append(queue, "world!")
	fmt.Print(queue)

	for len(queue) > 0 {
		queue[0] = ""
		queue = queue[1:]
		fmt.Println(queue)
		
		s:=(h.Samples,queue)
		h.Samples=s
		h.Logger.Info("add to log",h.Samples)
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func OdnSpis() {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	n3 := &ListNode{Val: 3}
	n2.Next = n3
	fmt.Println(n1.Val, n2.Val, n3.Val)
	i := n1
	for i != nil {
		fmt.Println(i.Val)
		i = i.Next
	}

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}

type LoggerDz struct {
	Samples string
	Logger  *logrus.Logger
}
