resource "aws_dynamodb_table" "db-config-dynamodb-table" {
  name           = "db-config"
  billing_mode   = "PROVISIONED" 
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "DefaultKey"

  depends_on = [aws_db_instance.db_ghtmx]

  attribute {
    name = "DefaultKey"
    type = "S"
  }

  global_secondary_index {
    name               = "DBEndpointIndex"
    hash_key           = "DefaultKey"
    write_capacity     = 10
    read_capacity      = 10
    projection_type    = "INCLUDE"
    non_key_attributes = ["DefaultKey"]
  }

  tags = {
    Name = "DB DynamoDB Table"
  }
}

resource "aws_dynamodb_table_item" "item" {
  table_name = aws_dynamodb_table.db-config-dynamodb-table.name
  hash_key   = aws_dynamodb_table.db-config-dynamodb-table.hash_key
  depends_on = [aws_dynamodb_table.db-config-dynamodb-table]

  item = <<ITEM
{
  "DefaultKey": {"S": "STORAGE_ADDR"},
  "Endpoint": {"S": "${aws_db_instance.db_ghtmx.endpoint}"}
}
ITEM
}

