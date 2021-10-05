
data "fugue_rules" "enabled_custom_rules" {
  filter {
    name   = "status"
    values = "ENABLED"
  }
}

output "enabled_custom_rule_ids" {
  value = data.fugue_rules.enabled_custom_rules.ids
}
