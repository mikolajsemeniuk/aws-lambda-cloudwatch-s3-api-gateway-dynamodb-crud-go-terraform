//go:build remove

package main

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
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
		values := "ALL_OLD"
		input := &dynamodb.DeleteItemInput{Key: key, TableName: &table, ReturnValues: &values}

		result, err := svc.DeleteItem(input)
		if err != nil {
			return Response{Body: err.Error(), StatusCode: 500}, nil
		}

		if result.Attributes == nil {
			return Response{Body: `{ "message": "not found" }`, StatusCode: 404}, nil
		}

		return Response{StatusCode: 204, Body: `{ "message": "order deleted" }`}, nil
	}
}
