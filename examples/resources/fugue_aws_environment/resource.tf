
variable "role_arn" {
  type = string
}

data "fugue_aws_types" "all" {
  region = "us-east-1"
}

resource "fugue_aws_environment" "example" {
  name                = "example"
  role_arn            = var.role_arn
  regions             = ["*"]
  compliance_families = ["CIS-AWS_v1.3.0", "CIS-Docker_v1.2.0"]
  resource_types      = data.fugue_aws_types.all.types
}
