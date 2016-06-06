package main

import (
	"fmt"
	// "html/template"
	"log"
	"net/http"
	// "os"
	// "strings"
)

func main() {
	http.HandleFunc("/", index)
	log.Println("Start listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	log.Println("Index home")

	res.Write([]byte(fmt.Sprintf("welcome ! 妮3")))
}
