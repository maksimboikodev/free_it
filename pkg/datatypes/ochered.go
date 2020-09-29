package datatypes

import "fmt"

func OcheredFIFO() {
	var queue []string

	queue = append(queue, "Hello ") // Добавление в очередь
	queue = append(queue, "world!")
	fmt.Print(queue)

	for len(queue) > 0 {
		//fmt.Print(queue[0]) // Первый элемент
		queue = queue[1:] // Удаление из очереди
		fmt.Print(queue)
		//queue[0] = ""
		//queue = queue[1:]
	}

	fmt.Print(queue)

}
