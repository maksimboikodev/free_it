package urlshortener

import (
	"strconv"
)

type URLStore struct {
	Urls map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		Urls: make(map[string]string),
	}
}

func (s *URLStore) Get(key string) string {
	if url, ok := s.Urls[key]; ok {
		return url
	}
	return ""
}

func (s *URLStore) Set(key, url string) bool {
	_, ok := s.Urls[key]
	if !ok {
		s.Urls[key] = url
	}
	return !ok
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
