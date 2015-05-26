package main

import (
	"github.com/Unknwon/macaron"
	. "github.com/shaalx/echo/oauth2"
	"net/http"
	// "github.com/macaron-contrib/pongo2"
)

var (
	OA *OAGithub
)

func init() {
	OA = NewOAGithub("8ba2991113e68b4805c1", "b551e8a640d53904d82f95ae0d84915ba4dc0571", "user")
}

func main() {
	m := macaron.Classic()
	// m.Use(pongo2.Pongoer())
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["avatar_url"] = "https://avatars.githubusercontent.com/u/5652582?v=3"
		ctx.HTML(200, "index") // 200 is the response code.
	})
	m.Get("/signin", signin)

	m.Run()
}

func signin(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, OA.AuthURL(), 302)
}
