#!/bin/bash

db_addr="$(aws dynamodb get-item --table-name db-config --key '{"DefaultKey": {"S": "STORAGE_ADDR"}}' | jq -r '.Item.Endpoint.S')"
echo "PostgreSQL endpoint: ${db_addr}"
cmd="sudo apt-get update && \
  sudo apt install -y jq && \
  export STORAGE_ADDR=${db_addr} && \
  command -v docker > /dev/null && \
  sudo docker pull iypetrov/ghtmx:latest && \
  sudo docker service update --image iypetrov/ghtmx:latest ghtmx"
while read instance_id; do
  aws ssm send-command \
  --instance-ids "${instance_id}" \
  --document-name "AWS-RunShellScript" \
  --parameters commands="${cmd}" > /dev/null 2>&1
  echo "${instance_id}"
done < <(aws ec2 describe-instances \
  --query "Reservations[*].Instances[*].InstanceId" \
  --filters Name=instance-state-name,Values=running \
  --output text)


