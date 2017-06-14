package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	log.Println("Starting backplane-poc app.")

	r := mux.NewRouter()
	r.HandleFunc("/_hc", HealthCheckHandler)
	r.HandleFunc("/", Root)

	// TODO:  Handle incoming port env var.  Don't hard code to 8080

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func Root(w http.ResponseWriter, req *http.Request) {

	html := `<!doctype html>
	           <head>
	             <title>backplane poc</title>
	             <style>
	               body {
	               	background-color:blue;
				    font-family: 'Raleway', serif;
				    font-size:24pt;
				    font-weight:400;
	               	color: white;
	               }
	             </style>
	             <link href="https://fonts.googleapis.com/css?family=Raleway:400" rel="stylesheet">
	           </head>
	           <body>
	             <h1>hello from backplane poc !</h1>
	           </body>
           </html>`

	template, err := template.New("name").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template.Execute(w, nil)
}

// health check...
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {

	// TODO return json / 200 here instead of hello world.

	io.WriteString(w, "hello, world!\n")
}
