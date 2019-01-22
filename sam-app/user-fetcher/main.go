package main

import (
	// "errors"
	"fmt"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("Got error creating session:")
		return events.APIGatewayProxyResponse{}, err
	}
	svc := dynamodb.New(sess)

	// filt := expression.Name("name").Equal(expression.Value(name))
	proj := expression.NamesList(expression.Name("name"), expression.Name("age"), expression.Name("created_at"))
	// expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		return events.APIGatewayProxyResponse{}, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("users"),
	}
	users, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		return events.APIGatewayProxyResponse{}, err
	}

	jsonBytes, _ := json.Marshal(users)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
