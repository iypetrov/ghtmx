resource "aws_flow_log" "lam_logs" {
  iam_role_arn    = aws_iam_role.lam_role.arn
  log_destination = aws_cloudwatch_log_group.lam_cloudwatch.arn
  traffic_type    = "ALL"
  vpc_id          = var.vpc_id 
}

