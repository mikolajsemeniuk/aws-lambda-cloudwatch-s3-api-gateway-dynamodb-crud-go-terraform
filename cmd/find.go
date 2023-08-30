//go:build find

package main

import (
	"aws-lambda-cloudwatch-s3-api-gateway-dynamodb-crud-go-terraform/api"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func init() {
	handler = func(c Context, r Event) (Response, error) {
		hash := r.QueryStringParameters["hash"]
		if _, err := uuid.Parse(hash); err != nil {
			return Response{Body: `{ "message": "cannot parse hash to uuid" }`, StatusCode: 400}, nil
		}

		sort := r.QueryStringParameters["sort"]
		if _, err := time.Parse(time.RFC3339, sort); err != nil {
			return Response{Body: `{ "message": "cannot parse sort to date time" }`, StatusCode: 400}, nil
		}

		key := map[string]*dynamodb.AttributeValue{"Key": {S: &hash}, "Created": {S: &sort}}
		result, err := svc.GetItem(&dynamodb.GetItemInput{Key: key, TableName: &table})
		if err != nil {
			return Response{Body: err.Error(), StatusCode: 500}, nil
		}

		if result.Item == nil {
			return Response{Body: `{ "message": "not found" }`, StatusCode: 404}, nil
		}

		var order api.Order
		if err := dynamodbattribute.UnmarshalMap(result.Item, &order); err != nil {
			return Response{Body: err.Error(), StatusCode: 500}, nil
		}

		response, err := json.Marshal(order)
		if err != nil {
			return Response{Body: err.Error(), StatusCode: 500}, nil
		}

		return Response{StatusCode: 200, Body: string(response)}, nil
	}
}
