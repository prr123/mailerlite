package main

import (
	"context"
	"log"
	"os"

	"github.com/mailerlite/mailerlite-go"
)


func main() {

	APITok, err  := os.ReadFile("/home/peter/install/ml_api_token.txt")
    if err != nil {log.Fatalf("error -- reading token: %v\n", err)}

	client := mailerlite.NewClient(string(APITok))

	ctx := context.TODO()

	subscriber := &mailerlite.Subscriber{
		Email: "peter@spacescabo.com",
		Fields: map[string]interface{}{
			"city": "Valencia",
		},
	}

	newSubscriber, _, err := client.Subscriber.Create(ctx, subscriber)
	if err != nil {
		log.Fatalf("error -- create subsciber: %v\n", err)
	}

	log.Printf("success added: %s\n", newSubscriber.Data.Email)
}
