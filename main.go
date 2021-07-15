package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("Hello %s %s!", name.FirstName, name.LastName)
	return fmt.Sprintf("Hello %s %s!", name.FirstName, name.LastName), nil
}

func main() {
	log.Printf("lambda started!")
	lambda.Start(HandleRequest)
	log.Printf("lambda finished!!")
}
