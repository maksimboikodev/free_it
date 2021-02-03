package gorilla

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/maksimboikodev/test/pkg/csvwork"
	"github.com/maksimboikodev/test/pkg/storage"
)

func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("TEST")
	fmt.Fprint(w, response)
}

func ParseHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("index.html")
	Check(err)
	err = html.Execute(writer, nil)
	Check(err)
}

func ExecuteTemplate(text string, data interface{}) {
	tmpl, err := template.New("Parse").Parse(text)
	Check(err)
	err = tmpl.Execute(os.Stdout, data)
	Check(err)
}

func CsvHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("products.csv")
	Check(err)
	err = html.Execute(writer, nil)
	Check(err)
}

func ReadCsvHandler(writer http.ResponseWriter, request *http.Request) {

	csv, err := csvwork.Readcsv()
	CheckPanic(err)
	response := fmt.Sprintf("CSV", csv)
	fmt.Fprint(writer, response)
}

func DBHandler(writer http.ResponseWriter, request *http.Request) {

	db, err := storage.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	h := storage.NewPersonRepository(db)

	p := storage.User{First_name: "Mack", Last_name: "jack", Age: 35}
	err = h.AddRecord(&p)
	Check(err)
	sel, err := h.FindAll()
	CheckPanic(err)
	response := fmt.Sprintf("Print DB ", sel)
	fmt.Fprint(writer, response)
}
