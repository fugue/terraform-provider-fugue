
data "fugue_rule_waiver" "example" {
  filter {
    name   = "id"
    values = ["8e74d107-bf6c-4b56-8c1b-895484aa825a"] # waiver ID
  }
}

output "waiver_name" {
  value = data.fugue_rule_waiver.example.name
}
