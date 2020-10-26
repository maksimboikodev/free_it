package urlshortener

import (
	"fmt"
	"strconv"
)

type URLStore struct {
	Urls map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		Urls: make(map[string]string)}
}

func (s *URLStore) Get(key string) string {
	if u, ok := s.Urls[key]; ok {
		fmt.Println(" url of ", u)
		return u
	}
	return ""
}

func (s *URLStore) Set(key, url string) bool {
	bl := false
	for _, value := range s.Urls {
		if value == url {
			bl = true
		}
	}
	if !bl {
		s.Urls[key] = url
		return true
	}
	return false

}
func (s *URLStore) Put(url string) string {
	k := genKey(len(s.Urls))
	g := s.Set(k, url)
	if g == true {
		s.Urls[k] = url
	}
	fmt.Print(g)
	fmt.Print(s.Urls)
	return ""
}

func (s *URLStore) Count() int {
	return len(s.Urls)
}
func genKey(n int) string {
	if n < 10 {
		shortUrl := strconv.Itoa(n) + "_URL"
		return shortUrl
	}
	shortUrl := strconv.Itoa(n) + "_urllll"
	return shortUrl

}
