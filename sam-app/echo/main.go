package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"

	"github.com/apex/gateway"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type User struct {
	Name      string
	Age       int
	CreatedAt string
}

func getUsers(c echo.Context) error {
	fmt.Println("===== [GET] Hello, World!")

	sess, err := session.NewSession()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Got error creating session.")
	}
	svc := dynamodb.New(sess)

	proj := expression.NamesList(expression.Name("Name"), expression.Name("Age"), expression.Name("CreatedAt"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Got error building expression.")
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("SampleUsers"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Query API call failed.")
	}
	fmt.Println(result)

	return c.String(http.StatusOK, "[GET] Hello, World!")
}

func createUser(c echo.Context) error {
	fmt.Println("===== [POST] Hello, World!")
	name := c.FormValue("name")
	age := c.FormValue("age")
	fmt.Println(name)
	fmt.Println(age)

	sess, err := session.NewSession()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Got error creating session.")
	}
	svc := dynamodb.New(sess)

	now := time.Now()
	agei, _ := strconv.Atoi(age)
	user := User{
		Name:      name,
		Age:       agei,
		CreatedAt: now.String(),
	}
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Got error marshalling map.")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("SampleUsers"),
	}
	result, err := svc.PutItem(input)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Query API call failed.")
	}
	fmt.Println(result)

	return c.String(http.StatusOK, "[POST] Hello, World!")
}

func deleteUser(c echo.Context) error {
	fmt.Println("===== [DELETE] Hello, World!")
	name := c.Param("name")
	age := c.Param("age")
	fmt.Println(name)
	fmt.Println(age)

	sess, err := session.NewSession()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Got error creating session.")
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
		return c.String(http.StatusInternalServerError, "Query API call failed.")
	}
	fmt.Println(result)

	return c.String(http.StatusOK, "[DELETE] Hello, World!")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/echo_users", getUsers)
	e.POST("/echo_users", createUser)
	e.DELETE("/echo_users/:name/:age", deleteUser)

	gateway.ListenAndServe(":3000", e)
}
