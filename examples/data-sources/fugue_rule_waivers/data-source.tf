data "fugue_waivers" "all_waivers_for_env" {
  filter {
    name   = "environment_id"
    values = ["d2350268-3570-4f3c-a909-6d1bdb4bf397"]
  }
}

output "waiver_ids" {
  value = data.fugue_waiver.all_waivers_for_env.ids
}
