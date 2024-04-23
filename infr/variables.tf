# General Variables

variable "region" {
  description = "Default region for provider"
  type        = string
  default     = "eu-west-2"
}

# EC2 Variables

variable "ami" {
  description = "Amazon machine image to use for EC2 instance"
  type        = string
  default     = "ami-0b9932f4918a00c4f"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}
