data "fugue_rules" "all_enabled" {
  filter {
    name   = "status"
    values = "ENABLED"
  }
}

output "all_enabled_rule_ids" {
  value = data.fugue_rule.all_enabled.ids
}
