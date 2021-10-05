
data "fugue_environment" "example" {
  filter {
    name   = "id"
    values = ["7c0000e0-d19c-41ff-a927-c9584d0bfdf6"]
  }
}

output "environment_name" {
  value = data.fugue_environment.example.name
}
