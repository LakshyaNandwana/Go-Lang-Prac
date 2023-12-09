package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello World!")
	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/",fs))
	// Serve the CSS file located in the same directory as the Go file
	//http.Handle("/Template.css", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":80",nil)
}