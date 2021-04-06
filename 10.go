package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	respPost, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	bodyPost, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		log.Fatalln(err)
	}

	respUser, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatalln(err)
	}
	bodyUser, err := ioutil.ReadAll(respUser.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(bodyPost)
	sb2 := string(bodyUser)
	log.Printf("POST :", sb+"USER :", sb2)
}
