package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// take in a payload and something with it 

func HandleRequest(event MyEvent) (string, error) {
	if event.Username  == ""{
		return "", fmt.Errorf("username cannot be empty")
	}


	return fmt.Sprintf("Successfully called by: %s", event.Username), nil
}

func main() {
	lambda.Start(HandleRequest)
	// fmt.Println("Hello world")
}