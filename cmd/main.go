package main

import (
	"aws-lambda-cloudwatch-s3-api-gateway-dynamodb-crud-go-terraform/api"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Response events.APIGatewayProxyResponse
type Event events.APIGatewayProxyRequest
type Context context.Context
type Order api.Order
type Input api.OrderInput

var conn *session.Session
var svc *dynamodb.DynamoDB
var table = "orders"
var handler func(Context, Event) (Response, error)

func main() {
	conn = session.Must(session.NewSession())
	svc = dynamodb.New(conn)
	lambda.Start(handler)
}
