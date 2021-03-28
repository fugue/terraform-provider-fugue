variable "tenant_id" {
  description = "Azure tenant_id"
  type        = string
  sensitive   = true
}

variable "secret" {
  description = "Azure secret"
  type        = string
  sensitive   = true
}

resource "fugue_azure_environment" "test_azure_env" {
  name                   = "tf-azure-test-1"
  application_id         = "xxxxx"
  subscription_id        = "xxxxxxx"
  secret                 = var.secret
  tenant_id              = var.tenant_id
  compliance_families    = ["CISAZURE"]
  survey_resource_groups = ["*"]
}

output "aws_env_id" {
  value = fugue_azure_environment.test_azure_env.id
}
