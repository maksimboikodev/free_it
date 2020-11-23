package csvwork

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type book struct {
	title    string
	price    string
	quantity string
}

func Createcsv() {
	rows := [][]string{
		{"The ABC of Go", "255", "1500"},
		{"Functional Programming with Go", "56", "280"},
		{"Go for It", "459", "356"},
		{"The Go Way", "55", "500"},
	}

	csvfile, err := os.Create("products.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvfile)
	for _, row := range rows {
		_ = csvwriter.Write(row)
		fmt.Println(row)
	}
	csvwriter.Flush()
	csvfile.Close()
	Readcsv()
}

func Readcsv() {
	file, err := os.Open("products.csv")
	if err != nil {
		fmt.Println("file not found")
		Createcsv()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sr []book
	for scanner.Scan() {
		b := strings.Split(scanner.Text(), ",")
		sr = append(sr, book{title: b[0], price: b[1], quantity: b[2]})
	}
	fmt.Println(sr)
}
