package main

import (
	// "errors"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	age := request.PathParameters["age"]

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Got error creating session:")
		return events.APIGatewayProxyResponse{}, err
	}
	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
			"Age": {
				N: aws.String(age),
			},
		},
		TableName: aws.String("SampleUsers"),
	}
	result, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Query API call failed:")
		return events.APIGatewayProxyResponse{}, err
	}

	resJsonBytes, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		Body:       string(resJsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
