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

/*func (s *URLStore) Get(key string) string {

	for key, value := range s.Urls {
		fmt.Println("Ключ:", key, "Значение", value)

		if _, ok := s.Urls["TUT"]; ok {
			fmt.Println("Найдено", key, value, ok)
		}

		/*if key, ok := s.Urls["Key"]; ok {
			fmt.Println(key, ok)
		}
	}
	return key
}*/

/*func (s *URLStore) Set(key, url string) bool {
	// ваш код
}*/

/*func (s *URLStore) Count() int {
	i := 1
	for key := range s.Urls {
		fmt.Println(key)
		i++
	}
	return i
}*/
/*func (s *URLStore) Count() int {
	sum,ok := s.Urls{
		fmt.Println(sum)
	}

	return len(sum)
}*/
/*func (s *URLStore) Get(key string) string {

	for key, value := range s.Urls {
		fmt.Println("Ключ:", key, "Значение", value)

	}

	key, ok := s.Urls["TUT"]
	if ok {
		fmt.Println("Найдено", key, ok)
		return key
	}
	return key
}*/

/*func (s *URLStore) Get(key string) string {

	for key, value := range s.Urls {
		fmt.Println("Ключ:", key, "Значение", value)

	}

	if key, ok := s.Urls["TUT"]; ok {
		fmt.Println(key, ok)
	}
	return key
}*/
func (s *URLStore) Get(key string) string {
	var val string
	_, ok := s.Urls["TUT"]
	if ok {
		val = s.Urls["TUT"]

	}
	return val
}
