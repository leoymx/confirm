package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("ready...")
	http.HandleFunc("/echouj", echo)
	err := http.ListenAndServe(":80", nil)
	if check_err(err) {
		return
	}
	log.Println("go...")
}

func echo(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("[ECHO]"))
	q := req.URL.Query()
	b, err := json.Marshal(q)
	if check_err(err) {
		rw.Write([]byte("echo ： error"))
		return
	}
	rw.Write(b)
}

func check_err(err error) bool {
	if nil != err {
		return true
	}
	return false
}
