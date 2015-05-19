package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("ready...")
	http.HandleFunc("/echouj", echo)
	http.HandleFunc("/", httpout)
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

func httpout(rw http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	site := q.Get("site")
	log.Printf(" visit http://%s\n", site)
	resp, err := http.Get("http://" + site)
	if check_err(err) {
		log.Printf(" visit https://%s\n", site)
		resp, err = http.Get("https://" + site)
		if check_err(err) {
			rw.Write([]byte("site ： error"))
			return
		}
	}
	b, err := ioutil.ReadAll(resp.Body)
	if check_err(err) {
		rw.Write([]byte(err.Error()))
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
