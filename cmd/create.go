//go:build create

package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func init() {
	handler = func(c Context, r Event) (Response, error) {
		var input Input
		if err := json.Unmarshal([]byte(r.Body), &input); err != nil {
			return Response{StatusCode: 400, Body: `{ "message": "` + err.Error() + `" }`}, nil
		}

		order := Order{
			Key:       uuid.New().String(),
			Name:      input.Name,
			Amount:    input.Amount,
			Price:     input.Price,
			Completed: input.Completed,
			Created:   time.Now().Format(time.RFC3339),
			Updated:   (time.Time{}).Format(time.RFC3339),
		}

		item, err := dynamodbattribute.MarshalMap(order)
		if err != nil {
			return Response{StatusCode: 400, Body: err.Error()}, nil
		}

		if _, err := svc.PutItem(&dynamodb.PutItemInput{Item: item, TableName: &table}); err != nil {
			return Response{StatusCode: 400}, nil
		}

		return Response{StatusCode: 200, Body: `{ "message": "order added" }`}, nil
	}
}
