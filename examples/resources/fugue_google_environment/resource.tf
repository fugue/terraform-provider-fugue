variable "service_account_email" {
  type = string
}

variable "project_id" {
  type = string
}

resource "fugue_google_environment" "example" {
  name                  = "example"
  service_account_email = var.service_account_email
  project_id            = var.project_id
  compliance_families   = ["CIS-Google_v1.1.0"]
}