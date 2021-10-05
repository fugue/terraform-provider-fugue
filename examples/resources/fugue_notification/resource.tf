
resource "fugue_notification" "security_emails" {
  name         = "Compliance notifications"
  emails       = ["security-team@example.com", "compliance-team@example.com"]
  environments = ["7c0000e0-d19c-41ff-a927-c9584d0bfdf6", "f305ef6d-10b8-44f7-9200-8ccfffb383d8"]
  events       = ["compliance"]
}

resource "fugue_notification" "drift_emails" {
  name         = "Drift notifications"
  emails       = ["joe@acme.com"]
  environments = ["7c0000e0-d19c-41ff-a927-c9584d0bfdf6"]
  events       = ["drift"]
}

resource "fugue_notification" "sns_topic_example" {
  name         = "Drift notifications"
  topic_arn    = "arn:aws:sns:us-east-2:0123456789:your-topic-arn"
  environments = ["7c0000e0-d19c-41ff-a927-c9584d0bfdf6"]
  events       = ["compliance", "drift"]
}
