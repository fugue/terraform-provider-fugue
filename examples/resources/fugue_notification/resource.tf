resource "fugue_notification" "notification1" {
  name         = "My Notification"
  emails       = ["foo@example.com", "bar@example.com"]
  environments = ["7c0000e0-d19c-41ff-a927-c9584d0bfdf6", "f305ef6d-10b8-44f7-9200-8ccfffb383d8"]
  events       = ["remediation", "compliance", "drift"]
}
