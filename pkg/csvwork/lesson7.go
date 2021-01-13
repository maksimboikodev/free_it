package csvwork

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type book struct {
	title    string
	price    int
	quantity int
}

func Readcsv() ([]book, error) {
	file, err := os.Open("products.csv")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sr []book
	for scanner.Scan() {
		b := strings.Split(scanner.Text(), ",")
		title := b[0]
		price, err := strconv.Atoi(b[1])
		if err != nil {
			return nil, err
		}
		quantity, err := strconv.Atoi(b[2])
		if err != nil {
			return nil, err
		}
		sr = append(sr, book{title, price, quantity})
	}
	return sr, err
}
