
resource "fugue_rule" "rule1" {
  name           = "RDS example rule"
  description    = "RDS instances should not be set as publicly accessible."
  cloud_provider = "AWS"
  severity       = "High"
  resource_type  = "AWS.RDS.Instance"
  rule_text      = <<EOF
default allow = false

allow {
    input.publicly_accessible == false
}
EOF
}
