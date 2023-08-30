//go:build update

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
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

		var body Input
		if err := json.Unmarshal([]byte(r.Body), &body); err != nil {
			return Response{StatusCode: 400, Body: `{ "message": "` + err.Error() + `" }`}, nil
		}

		key := map[string]*dynamodb.AttributeValue{"Key": {S: &hash}, "Created": {S: &sort}}
		values := map[string]*dynamodb.AttributeValue{
			":name":      {S: &body.Name},
			":amount":    {N: aws.String(fmt.Sprintf("%v", body.Amount))},
			":price":     {N: aws.String(fmt.Sprintf("%v", body.Price))},
			":completed": {BOOL: &body.Completed},
			":updated":   {S: aws.String(time.Now().Format(time.RFC3339))},
		}
		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: values,
			ExpressionAttributeNames: map[string]*string{
				"#N": aws.String("Name"),
				"#A": aws.String("Amount"),
				"#P": aws.String("Price"),
				"#C": aws.String("Completed"),
				"#U": aws.String("Updated"),
			},
			TableName:        aws.String(table),
			Key:              key,
			ReturnValues:     aws.String("ALL_OLD"),
			UpdateExpression: aws.String("SET #N = :name, #A = :amount, #P = :price, #C = :completed, #U = :updated"),
		}

		result, err := svc.UpdateItem(input)
		if err != nil {
			return Response{Body: err.Error(), StatusCode: 500}, nil
		}

		if result.Attributes == nil {
			return Response{Body: `{ "message": "not found" }`, StatusCode: 404}, nil
		}

		return Response{StatusCode: 200, Body: `{ "message": "order updated" }`}, nil
	}
}
