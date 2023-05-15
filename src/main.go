package main

import (
	"fmt"
	"main/XWeb"
	"net/http"
)

func main() {
	r := XWeb.New()

	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	r.Run(":9999")
}
