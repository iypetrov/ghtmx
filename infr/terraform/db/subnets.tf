resource "aws_subnet" "db_subnet" {
  vpc_id                  = var.vpc_id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "eu-west-2a"
  map_public_ip_on_launch = true

  tags = {
    Name = "DB Subnet"
  }
}

resource "aws_subnet" "db_tmp_subnet" {
  vpc_id                  = var.vpc_id
  cidr_block              = "10.0.3.0/24"
  availability_zone       = "eu-west-2b"
  map_public_ip_on_launch = true

  tags = {
    Name = "DB Tmp Subnet"
  }
}

resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "db_subnet_group"
  subnet_ids = [aws_subnet.db_subnet.id, aws_subnet.db_tmp_subnet.id]

  tags = {
    Name = "DB Subnet Group"
  }
}
