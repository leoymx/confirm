package main

import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
)

func main() {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["avatar_url"] = "https://avatars.githubusercontent.com/u/5652582?v=3"
		ctx.HTML(200, "index") // 200 is the response code.
	})

	m.Run()
}
