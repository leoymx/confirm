package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", echo)
	err := http.ListenAndServe(":8784", nil)
	if check_err(err) {
		return
	}
}

func echo(rw http.ResponseWriter, req *http.Request) {
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
