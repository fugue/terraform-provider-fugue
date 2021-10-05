
resource "fugue_family" "acme_rules" {
  name           = "ACME Rules"
  description    = "Rules to be used for ACME security and compliance"
  recommended    = true
  always_enabled = false
  rule_ids       = [
    "FG_R00277",         # Fugue built-in rule
    "FG_R00229",         # Fugue built-in rule
    fugue_rule.rule1.id, # Custom rule ID, for a rule defined in Terraform
  ]
}
