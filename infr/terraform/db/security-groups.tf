resource "aws_security_group" "db_security_group" {
  egress = [
    {
      cidr_blocks      = ["10.0.1.0/24"]
      description      = ""
      from_port        = 8080 
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
      to_port          = 8080 
    }
  ]
  ingress = [
    {
      cidr_blocks      = ["10.0.2.0/24", "10.0.3.0/24"]
      description      = ""
      from_port        = 5432 
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
      to_port          = 5432
    }
  ]
  vpc_id = var.vpc_id
  tags = {
    Name = "DB Security Group"
  }
}

