package main

import (
	"net/http"
)

func main() {
	s := "Hello, world!"
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		html := `
    <html>
    <head>
        <title>Hello</title>
        <script type="text/javascript" src="/assets/hello.js"></script>
        <link href="/assets/hello.css" rel="stylesheet" />
    </head>
    <body>
        <span class="hello">` + s + `</span>
    </body>
    </html>
    `
		res.Header().Set("Content-Type", "text/html") //config HTML header
		res.Write([]byte(html))                       // response to web browser
	})
	http.Handle(
		"/assets/", // run file server when client access to /assets/ path
		http.StripPrefix( // when run the file server, delets /assets/ in URL path cus directory was designated
			"/assets/", http.FileServer(http.Dir("assets"))), // in webserver, view sub files of the assets directory
	)
	http.ListenAndServe(":80", nil) // run webserver in 80 port
}
