resource "aws_s3_bucket" "lambda_binaries_bucket" {
  bucket = "bucket-for-lambdas-binaries"
}

locals {
  binaries = [
    "list",
    "find",
    "create",
    "update",
    "remove"
  ]
}

resource "aws_s3_object" "lambda_binary" {
  for_each = toset(local.binaries)
  bucket   = aws_s3_bucket.lambda_binaries_bucket.bucket
  key      = "orders/${each.key}.zip"
  source   = "../bin/${each.key}.zip"
  acl      = "private"
}
