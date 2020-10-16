package urlshortener

import "fmt"

type URLStore struct {
	Urls map[string]string
}

func NewURLStore() *URLStore {
	var urls *URLStore
	urls = &URLStore{}
	fmt.Println(urls)
	return urls
}
