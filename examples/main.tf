terraform {
  required_providers {
    fugue = {
      version = "0.1"
      source  = "fugue.co/co/fugue"
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
