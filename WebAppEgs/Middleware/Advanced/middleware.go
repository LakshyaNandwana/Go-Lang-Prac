package main

import(
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

//Logging logs all requests with its path adn the time it took to process
func Logging() Middleware{
	//Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc{

		//Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request){
			//Do middleware things
			start := time.Now()
			defer func(){
				log.Println(r.URL.Path, time.Since(start))
			}()

			//Call the next middleware/handler
			f(w, r)
		}
	}
}

func Method(m string) Middleware{
	//Create a new middleware
	return func(f http.HandlerFunc) http.HandlerFunc{

		//Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request){

		//Do middleware things 
			if r.Method != m{
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			//Call the nexxt middleware/handler in chain
			f(w,r)
		}
	}
}

//Chain applies middlewares to a http.handlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, m := range middlewares{
		f =m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "HelloWorld")
}

func main(){
	http.HandleFunc("/", Chain(Hello, Method("GET"),Logging()))
	http.ListenAndServe(":8080", nil)
}