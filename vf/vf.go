package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func startServer() {
	m := martini.Classic()
	m.Use(martini.Static("./assets"))
	http.ListenAndServe(":8090", m)
}

func main() {
	go startServer()
	select {}
}
