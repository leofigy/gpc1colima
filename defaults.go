package main

import(
	"fmt"
	"time"
	"log"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	index = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")),)
)

// data 

type indexData struct {
	Logo string
	Style string
	RequestTime string		
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func indexFunc(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	data := indexData{
		//Logo: "/public/gcp-gopher.svg",
		Logo: "/public/surf.webp",
		Style: "/public/style.css",
		RequestTime: time.Now().Format(time.RFC822),
	}

	if err := index.Execute(w, data); err != nil {
		log.Printf("Error loading template %+v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}