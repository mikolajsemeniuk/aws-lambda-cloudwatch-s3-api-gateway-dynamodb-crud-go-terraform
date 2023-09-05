resource "aws_lambda_function" "lambda_orders" {
  for_each      = { for route in local.routes : route.name => route }
  function_name = "${each.key}-order"
  handler       = each.key
  runtime       = "go1.x"
  role          = aws_iam_role.lambda_role.arn
  s3_key        = aws_s3_object.lambda_binary[each.key].key
  s3_bucket     = aws_s3_bucket.lambda_binaries.bucket
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      },
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_role_policy" "lambda_role_policy" {
  name = "lambda_role_policy"
  role = aws_iam_role.lambda_role.id
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : ["dynamodb:*"],
        "Resource" : "${aws_dynamodb_table.orders.arn}"
      }
    ]
  })
}
