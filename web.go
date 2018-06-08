package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"path"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	tmplFile := path.Join(tmplDir, "index.html")

	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}

func serveEcho(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Echo")
}

func ServeWeb(addr string) error {
	var serverMutex = http.NewServeMux()

	serverMutex.HandleFunc("/", serveHome)
	serverMutex.HandleFunc("/echo", serveEcho)

	return http.ListenAndServe(addr, serverMutex)
}
