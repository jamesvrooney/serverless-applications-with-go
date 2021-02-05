package main

import "github.com/aws/aws-lambda-go/lambda"

func handler() (string, error) {
	return "Welcome to Serverless world, James", nil
}

func main() {
	lambda.Start(handler)
}
