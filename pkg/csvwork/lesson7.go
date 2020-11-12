package csvwork

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type book struct {
	Title string
}

func Createcsv() {
	rows := [][]string{
		{"Title", "Price", "Quantity"},
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
	}
	csvwriter.Flush()
	csvfile.Close()
}

func Readcsv() {
	file, err := os.Open("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var b = book{scanner.Text()}
		//fmt.Println(scanner.Text())
		fmt.Println(b)
	}
}
