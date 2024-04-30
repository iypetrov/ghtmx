resource "aws_cloudwatch_log_group" "lam_cloudwatch" {
  name = "lam-cloudwatch"
}

resource "aws_cloudwatch_query_definition" "lam_default_query" {
  name = "lam_default_query"

  log_group_names = [
    "lam-cloudwatch"
  ]

  query_string = <<EOF
fields @timestamp, srcAddr, srcPort, dstAddr, dstPort, bytes
EOF
}
