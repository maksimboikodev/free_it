package urlshortener

import (
	"fmt"
	"math"
	"strconv"
)

type URLStore struct {
	Urls map[string]string
}
type ErrNegativeSqrt float64

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

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot square root a negative number ", float64(e))
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, ErrNegativeSqrt(value)
	}
	return math.Sqrt(value), nil
}
