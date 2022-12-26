package main

import (
	"net/http"

	"github.com/goweb/myapp3/app"
)

func main() {

	http.ListenAndServe(":3000", app.NewHandler())
}
