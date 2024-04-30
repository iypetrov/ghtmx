provider "aws" {
  region = "eu-west-2"
}

terraform {
  backend "s3" {
    bucket = "ghtmx-tf-state"
    key    = "infr/terraform.tfstate"
    region = "eu-west-2"
  }
}

module "db" {
  source = ".//db"
  vpc_id = aws_vpc.vpc_ghtmx.id
  db_username = var.db_username
  db_password = var.db_password
}

module "backend" {
  source = ".//backend"
  vpc_id = aws_vpc.vpc_ghtmx.id
}

module "lam" {
  source = ".//lam"
  vpc_id = aws_vpc.vpc_ghtmx.id
}

