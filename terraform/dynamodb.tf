resource "aws_dynamodb_table" "orders" {
  name         = "orders"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Key"
  range_key    = "Created"

  attribute {
    name = "Key"
    type = "S"
  }

  attribute {
    name = "Created"
    type = "S"
  }

  attribute {
    name = "Name"
    type = "S"
  }

  attribute {
    name = "Amount"
    type = "N"
  }

  attribute {
    name = "Price"
    type = "N"
  }

  local_secondary_index {
    name            = "NameIndex"
    range_key       = "Name"
    projection_type = "ALL"
  }

  local_secondary_index {
    name            = "AmountIndex"
    range_key       = "Amount"
    projection_type = "ALL"
  }

  local_secondary_index {
    name            = "PriceIndex"
    range_key       = "Price"
    projection_type = "ALL"
  }
}
