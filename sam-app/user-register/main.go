package main

import (
	// "errors"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type UserParam struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	Name      string
	Age       int
	CreatedAt string
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	reqBody := request.Body
	reqjsonBytes := ([]byte)(reqBody)
	userReq := new(UserParam)
	if err := json.Unmarshal(reqjsonBytes, userReq); err != nil {
		fmt.Println("Got error json marshalling.", err)
		return events.APIGatewayProxyResponse{}, err
	}

	now := time.Now()
	user := User{
		Name:      userReq.Name,
		Age:       userReq.Age,
		CreatedAt: now.String(),
	}
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		return events.APIGatewayProxyResponse{}, err
	}

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Got error creating session:")
		return events.APIGatewayProxyResponse{}, err
	}
	svc := dynamodb.New(sess)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("SampleUsers"),
	}
	result, err := svc.PutItem(input)
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
