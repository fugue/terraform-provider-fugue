
data "fugue_rule" "example" {
  filter {
    name   = "id"
    values = ["bcc49dd2-dbc8-4599-827c-645e442097e4"] # custom rule ID
  }
}

output "rule_id" {
  value = data.fugue_rule.example.id
}
