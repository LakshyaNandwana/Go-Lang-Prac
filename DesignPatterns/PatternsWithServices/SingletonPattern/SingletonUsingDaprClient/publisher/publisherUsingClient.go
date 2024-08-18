package publisher

import (
	"context"
	"log"
	"sync"

	dapr "github.com/dapr/go-sdk/client"
)

type PubDetails struct {
	client dapr.Client
}

var (
	instance *PubDetails
	once     sync.Once
)

func GetInstance() *PubDetails {
	once.Do(func() {
		client, err := dapr.NewClient()
		if err != nil {
			log.Fatalf("Failed to create Dapr client: %v", err)
		}
		instance = &PubDetails{
			client: client,
		}
	})
	return instance
}

func PublishEvent(payload string) {

	ctx := context.Background()
	GetInstance().client.PublishEvent(ctx, "pubsubName", "TopicName", payload)

}
