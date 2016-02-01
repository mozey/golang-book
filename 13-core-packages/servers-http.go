package main

import (
	"net/http"; "io"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
	<html>
	  <head>
		  <title>Hello World</title>
	  </head>
	  <body>
		  Hello World!
	  </body>
	</html>`,
	)
}
func main() {
	// Handle a URL route (/hello) by calling the given function
	http.HandleFunc("/hello", hello)

	// Handle static files by using FileServer
	http.Handle(
		"/servers/",
		http.StripPrefix(
			"/servers/",
			http.FileServer(http.Dir("servers")),
		),
	)

	http.ListenAndServe(":9000", nil)
}