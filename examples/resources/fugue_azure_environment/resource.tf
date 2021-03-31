
variable "tenant_id" {
  description = "Azure tenant ID"
  type        = string
}

variable "subscription_id" {
  description = "Azure subscription ID"
  type        = string
}

variable "application_id" {
  description = "Azure application ID"
  type        = string
}

variable "client_secret" {
  description = "Azure client secret"
  type        = string
  sensitive   = true
}

resource "fugue_azure_environment" "example" {
  name                   = "example"
  tenant_id              = var.tenant_id
  subscription_id        = var.subscription_id
  application_id         = var.application_id
  client_secret          = var.client_secret
  compliance_families    = ["CISAZURE"]
  survey_resource_groups = ["*"]
}
