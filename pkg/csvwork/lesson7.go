package csvwork

import (
	"encoding/csv"
	"log"
	"os"
)

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
