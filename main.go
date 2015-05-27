package main

import (
	. "github.com/shaalx/echo/oauth2"
	"log"
	"net/http"
)

var (
	OA *OAGithub
)

func init() {
	OA = NewOAGithub("8ba2991113e68b4805c1", "b551e8a640d53904d82f95ae0d84915ba4dc0571", "user")
}

func signin(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, OA.AuthURL(), 302)
}

func main() {
	log.Println("server start...")
	// http.Handle("/css/", http.FileServer(http.Dir("templates")))
	// http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/signin", signin)
	http.ListenAndServe(":80", nil)
}
