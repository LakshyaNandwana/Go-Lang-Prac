package main

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/walmartdigital/gomock-tutorial-code/pkg/client"
)

type RestyClient struct{
	client *resty.Client
}

func NewRestyClient() *RestyClient{
	r := RestyClient{
		client: resty.New(),
	}
	return &r
}

func(r RestyClient) Get(url string) (int, []byte, error){
	resp, err := r.client.R().Get(url)
	body := resp.Body()
	return resp.StatusCode(), body, err
}

type RestyClientFactory struct{}


func (f RestyClientFactory) Create() client.HTTPClient{
	r := NewRestyClient()
	return *r
}




func monkeys(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love monkeys!")
}

func dogs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love dogs!")
}

func main() {
	http.HandleFunc("/monkeys", monkeys)
	http.HandleFunc("/dogs", dogs)
	go http.ListenAndServe(":8080", nil)

	zoo := client.NewZooClient(RestyClientFactory{})
	fmt.Println(zoo.ReadMessage("monkeys"))
	fmt.Println(zoo.ReadMessage("dogs"))
}
