package main

import(
	"html/template"
	"net/http"
	"fmt"
)

type ContactDetails struct{
	Email string
	Subject string
	Message string
}

func main(){
	tmp1 := template.Must(template.ParseFiles("forms/forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{
			tmp1.Execute(w, nil)
			return
		}
		details := ContactDetails{
			Email: r.FormValue("email"),
			Subject : r.FormValue("subject"),
			Message : r.FormValue("message"),
		}

		
		_ = details
		fmt.Println(details) //Printing the details on the screen

		tmp1.Execute(w, struct{Success bool}{true})
	})

	http.ListenAndServe(":8080",nil)
}