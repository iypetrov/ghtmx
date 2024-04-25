resource "aws_instance" "backend_ghtmx" {
  ami                    = "ami-0b9932f4918a00c4f"
  instance_type          = "t2.micro"
  subnet_id              = aws_subnet.backend_subnet.id
  vpc_security_group_ids = [aws_security_group.backend_security_group.id]
  iam_instance_profile   = aws_iam_instance_profile.deploy_profile.name
  user_data              = <<EOF
#!/bin/bash

sudo apt-get update
sudo apt install postresql-client

curl -fsSl https://get.docker.com -o docker-init.sh
sh docker-init.sh

sudo docker swarm init
sudo docker service create --name ghtmx --network host iypetrov/ghtmx:latest
EOF

  tags = {
    Name = "Backend EC2"
  }
}

