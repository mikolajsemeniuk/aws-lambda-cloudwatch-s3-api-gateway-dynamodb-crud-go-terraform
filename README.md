# AWS-lambda-cloudwatch-s3-api-gateway-dynamodb-crud-go-terraform

Provision your resources with following commands

```sh
terraform -chdir=terraform init
terraform -chdir=terraform plan
terraform -chdir=terraform apply -auto-approve
terraform -chdir=terraform destroy -auto-approve
```

Build binaries

```sh
# available tags: doc, list, find, create, remove
export tag=create
GOOS=linux GOARCH=amd64 go build -tags "$tag" -ldflags="-s -w" -o bin/$tag ./cmd
zip -j bin/$tag.zip bin/$tag
```

Test Lambda

```json
{
  "pathParameters": {},
  "headers": {},
  "queryStringParameters": {
    "hash": "6ceb32cb-efdd-47a3-acc9-e7a5e4cd945a",
    "sort": "2023-08-30T11:59:14Z"
  },
  "httpMethod": "GET",
  "body": "{\"Name\": \"new\", \"Amount\": 3, \"Price\": 1.5, \"Completed\": true }"
}
```
