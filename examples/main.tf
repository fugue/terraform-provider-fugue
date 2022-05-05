terraform {
  required_providers {
    fugue = {
      version = "0.0.10"
      source  = "fugue/fugue"
    }
  }
}

provider "fugue" {}

data "fugue_aws_types" "all" {
  region = "us-east-1"
}

variable "role_arn" {
  type = string
}

resource "fugue_aws_environment" "test" {
  name                = "tf-test-1"
  role_arn            = var.role_arn
  regions             = ["*"]
  compliance_families = ["FBP"]
  resource_types      = data.fugue_aws_types.all.types
}

output "aws_env_id" {
  value = fugue_aws_environment.test.id
}

# Waives "IAM root user access key should not exist" as an example
resource "fugue_rule_waiver" "waiver1" {
  name              = "waive-FG_R00004"
  comment           = "This is an example waiver!"
  environment_id    = fugue_aws_environment.test.id
  rule_id           = "FG_R00004"
  resource_type     = "AWS.IAM.CredentialReport"
  resource_provider = "*"
  resource_id       = "*"
}
