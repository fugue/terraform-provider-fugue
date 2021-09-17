data "fugue_rule" "rule1" {
  filter {
    name   = "name"
    values = ["Rule Name"]
  }
}

output "rule_id" {
  value = data.fugue_rule.rule1.id
}
