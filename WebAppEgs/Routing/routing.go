/*
Routing in GoLang using gorilla/mux
Run the following command in the terminal to get the package
Create a go.mod file inside the project folder
go mod init routing.go
downloads the dependency package
go get -u github.com/gorilla/mux
*/

package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r)
}
