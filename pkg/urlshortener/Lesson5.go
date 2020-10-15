package urlshortener

import "fmt"

type URLStore struct {
	urls map[string]string
}

func NewURLStore() *URLStore {
	var urls URLStore
	url := map[string]string{
		"Tut":     "tut.by",
		"Onliner": "Onliner.by",
		"Dev":     "dev.by",
		"Mail":    "mail.ru"}
	fmt.Println(url["Tut"])
	urls = *url
	fmt.Println(urls)

}
