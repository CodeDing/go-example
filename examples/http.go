//package main
//
//type Handler interface {
//	Read([]byte) (int64, error)
//	Write([]byte) (int64, error)
//}
//
//var Manager map[string]Handler
//
//var (
//	ErrAlreadyExist = errors.New("already exist it")
//)
//
//func Register(path string, h Handler) error {
//	if v, ok := Manager[path]; ok {
//		return ErrAdreadyExit
//	}
//	Manager[path] = h
//	return nil
//}

package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	q.Set("redirect", "http://www.baidu.com")
	u.RawQuery = q.Encode()
	fmt.Printf("Path => %s\n", u.Path)
	fmt.Printf("RawPath => %s\n", u.RawPath)
	fmt.Println("URL => ", u.String())
}
