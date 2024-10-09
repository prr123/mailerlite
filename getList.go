package main

import (
	"context"
	"log"
	"os"
	"github.com/mailerlite/mailerlite-go"
)


func main() {

	APITok, err  := os.ReadFile("/home/peter/install/ml_api_token.txt")
	if err != nil {log.Fatalf(" error reading token: %v\n", err)}

	client := mailerlite.NewClient(string(APITok))

	ctx := context.TODO()

	listOptions := &mailerlite.ListSubscriberOptions{
		Limit:  200,
		Page:   0, 
		Filters: &[]mailerlite.Filter{{
			Name:  "status", 
			Value: "active",
		}},
	}

	subscribers, _, err := client.Subscriber.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(subscribers.Meta.Total)
}
