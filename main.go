package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	usage = []byte(`<a href="www.shaalx-echouj.daoapp.io?site=www.baidu.com" ><h1>www.shaalx-echouj.daoapp.io?site=www.baidu.com</h1></a>` + "\n" + `
		<a href="www.shaalx-echouj.daoapp.io?site=blog.csdn.net/archi_xiao" ><h1>Archi_xiao 's blog (CSDN)</h1></a>` + "\n")
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
	rw.Write(usage)
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
	if len(site) < 1 {
		site = "127.0.0.1:80/echouj?well=I'm_comming&but=where_are_you?"
	}
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
