//go:build list

package main

import (
	"aws-lambda-cloudwatch-s3-api-gateway-dynamodb-crud-go-terraform/api"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func init() {
	handler = func(c Context, r Event) (Response, error) {
		result, err := svc.Scan(&dynamodb.ScanInput{TableName: &table})
		if err != nil {
			return Response{StatusCode: 400, Body: err.Error()}, nil
		}

		var orders []api.Order
		if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &orders); err != nil {
			return Response{StatusCode: 400, Body: err.Error()}, nil
		}

		body, err := json.Marshal(orders)
		if err != nil {
			return Response{StatusCode: 400, Body: err.Error()}, nil
		}

		return Response{StatusCode: 200, Body: string(body)}, nil
	}
}
