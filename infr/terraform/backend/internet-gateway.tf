resource "aws_internet_gateway" "backend_internet_gateway" {
  vpc_id = var.vpc_id
  tags = {
    Name = "Backend Internet Gateway"
  }
}
