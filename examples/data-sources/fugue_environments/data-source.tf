data "fugue_environments" "all_aws_environments" {
  filter {
    name   = "cloud_provider"
    values = ["aws"]
  }
}

output "environment_ids" {
  value = data.fugue_environment.all_aws_environments.ids
}
