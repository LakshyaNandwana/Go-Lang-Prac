package main

import (
	"html/template"
	"net/http"
)

type Todo struct{
	Title string
	Done bool
}

type TodoPageData struct{
	PageTitle string
	Todos []Todo
}

func main(){
	//defining .html or .css files for the webPage
	tmp1 := template.Must(template.ParseFiles("static/htmlMain.html")) //Alternate use Template.css  or htmlTemplate.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos : []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done:true},
				{Title: "Task 3", Done:false},
			},
		}
		tmp1.Execute(w, data)
	})
	http.ListenAndServe(":80",nil)
}