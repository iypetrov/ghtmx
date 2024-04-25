resource "aws_db_instance" "db_ghtmx" {
  allocated_storage        = 20
  engine                   = "postgres"
  engine_version           = "16.1"
  identifier               = "db-ghtmx"
  instance_class           = "db.t3.micro"
  db_subnet_group_name     = aws_db_subnet_group.db_subnet_group.name
  vpc_security_group_ids   = [aws_security_group.db_security_group.id]
  storage_encrypted        = false
  publicly_accessible      = false
  delete_automated_backups = true
  skip_final_snapshot      = true
  db_name                  = "ipdb"
  username                 = "foo"
  password                 = "foofoofoo"
  apply_immediately        = true
  multi_az                 = false

  tags = {
    Name = "DB PostgreSQL"
  }
}

