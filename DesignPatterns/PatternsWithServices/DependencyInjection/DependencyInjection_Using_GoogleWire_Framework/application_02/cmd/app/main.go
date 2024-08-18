package main

import (
	"DI-Wire/user"
	"database/sql"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler := user.Wire(db)
	http.Handle("/user", userHandler.FetchByUsername())
	http.ListenAndServe(":8080", nil)
}
