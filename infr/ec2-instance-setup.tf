resource "aws_instance" "ec2_ghtmx" {
  ami           = var.ami
  instance_type = var.instance_type
}
