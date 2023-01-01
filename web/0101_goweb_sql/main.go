package main

import (
	"net/http"

	"github.com/goweb/app"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)

}
