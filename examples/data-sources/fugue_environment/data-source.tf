data "fugue_environment" "environment1" {
  filter {
    name   = "name"
    values = ["Environment Name"]
  }
  filter {
    name   = "cloud_provider"
    values = ["aws"]
  }
}

output "environment_id" {
  value = data.fugue_environment.environment1.id
}
