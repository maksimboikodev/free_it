package urlshortener

import (
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
	if _, ok := s.Urls[key]; !ok {
		return "URL not found"
	}
	return s.Urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	if _, ok := s.Urls[key]; ok {
		return false
	}
	s.Urls[key] = url
	return true
}

func (s *URLStore) Put(url string) string {
	key := genKey(s.Count())
	s.Set(key, url)
	return key
}

func (s *URLStore) Count() int {
	return len(s.Urls)
}

func genKey(n int) string {
	return strconv.Itoa(n) + "_URL"

}
