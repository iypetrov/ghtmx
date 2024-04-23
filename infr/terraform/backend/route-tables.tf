resource "aws_route_table" "backend_route_table" {
  vpc_id = var.vpc_id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = var.gateway_id
  }
  tags = {
    Name = "Backend Route Table"
  }
}

resource "aws_route_table_association" "backend_route_table_association" {
  depends_on     = [aws_subnet.backend_subnet, aws_route_table.backend_route_table]
  subnet_id      = aws_subnet.backend_subnet.id
  route_table_id = aws_route_table.backend_route_table.id
}