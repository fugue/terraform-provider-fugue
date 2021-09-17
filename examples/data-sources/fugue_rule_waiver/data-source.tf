data "fugue_rule_waiver" "waiver1" {
  filter {
    name   = "name"
    values = ["Waiver Name"]
  }
  filter {
    name   = "environment_id"
    values = ["52604038-7989-497b-9b64-926299fc97f0"]
  }
  filter {
    name   = "rule_id"
    values = ["c4f2aeeb-8642-4410-bee8-55b2c9116067"]
  }
}

output "waiver_id" {
  value = data.fugue_waiver.waiver1.id
}
