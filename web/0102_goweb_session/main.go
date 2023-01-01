package main

import (
	"net/http"

	"github.com/goweb/app"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	http.ListenAndServe(":3000", m)

}
