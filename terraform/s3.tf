resource "aws_s3_bucket" "lambda_binaries" {
  bucket = "bucket-for-lambdas-binaries"
}

resource "aws_s3_object" "lambda_binary" {
  for_each = { for route in local.routes : route.name => route }
  bucket   = aws_s3_bucket.lambda_binaries.bucket
  key      = "orders/${each.key}.zip"
  source   = "../bin/${each.key}.zip"
  acl      = "private"
}
