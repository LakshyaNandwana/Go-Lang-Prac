package client

import (
	"fmt"
)

type HTTPClient interface{
	Get(url string)(int, []byte, error)
}

type HTTPClientFactory interface{
	Create() HTTPClient
}

type ZooClient struct{
	client HTTPClient
}

func NewZooClient(factory HTTPClientFactory) *ZooClient{
	client := ZooClient{
		client: factory.Create(),
	}
	return &client
}

// ReadMessage ...
func (z *ZooClient) ReadMessage(animal string) string {

	statusCode, body, err := z.client.Get(fmt.Sprintf("http://localhost:8080/%s", animal))

	if err !=nil{
		return string("Hi there, the zoo is closed!")
	}

	var singularDictionary = map[string]string{
		"elephants":"elephant",
		"dogs":"dog",
	}

	switch statusCode{
		case 200:
			return string(body)
		case 404:
			return fmt.Sprintf("Hi there, what is %s!", singularDictionary[animal])
		default:
			return string(body)
	}

}
