package urlshortener

import "fmt"

type URLStore struct {
	Urls map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{}
}

func (s *URLStore) Get(key string) string {
	if u, ok := s.Urls[key]; ok {
		fmt.Println(" url of ", u)
		return u
	}
	return ""
}

/*func NewURLStore() *URLStore {
	var urls *URLStore
	urls = &URLStore{}
	fmt.Println(urls)
	return urls
}


func (s *URLStore) Get(key string) string {
	var val string
	_, ok := s.Urls["TUT"]
	if ok {
		val = s.Urls["TUT"]

	}
	return val
}*/
func (s *URLStore) Set(key, url string) bool {
	_, ok := s.Urls[key]
	if ok {
		fmt.Println("Значение существует")
		return false
	} else {
		s.Urls[key] = url
		fmt.Println("Значение добавлено")
		return true

	}
}
