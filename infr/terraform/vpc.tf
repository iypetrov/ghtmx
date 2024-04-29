resource "aws_vpc" "vpc_ghtmx" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "VPC: ghtmx"
  }
}
