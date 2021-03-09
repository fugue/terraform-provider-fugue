# Terraform Provider for Fugue

This provider can be used to create and manage environments in Fugue. It is
under active development and should be considered "beta". Resource types are
still being added.

## Requirements

 - Terraform v0.14.x or higher
 - Go 1.14 (to build the provider plugin)

## Build

Run the following command in the root of this repository to build the provider.
This will produce a file named `terraform-provider-fugue` in the same directory.

```shell
make
```

## Install

Run the following to install the Fugue Terraform plugin into `~/.terraform.d/plugins`.

```shell
make install
```

## Example HCL

Take a look at the [example .tf file](./examples/main.tf).

## Resource Types

Create an AWS environment in Fugue using the `fugue_aws_environment` resource:

```hcl
resource "fugue_aws_environment" "test" {
  name = "tf-test-1"
  role_arn = var.role_arn
  regions = ["*"]
  compliance_families = ["FBP"]
  resource_types = data.fugue_aws_types.all.types
}
```

A `fugue_rule_waiver` resource can be used to create a waiver for a specified rule
and resource:

```hcl
resource "fugue_rule_waiver" "waiver1" {
  name = "waive-FG_R00229"
  comment = "This S3 bucket is intentionally public"
  environment_id = fugue_aws_environment.test.id
  rule_id = "FG_R00229"
  resource_type = "AWS.S3.Bucket"
  resource_provider = "aws.us-east-1"
  resource_id = "my-public-s3-bucket"
}
```

You may specify a resource_provider and resource_id of `*` in order to waive
multiple resources for a specific rule.

## Data Sources

Supported AWS resource types may be retrieved using the `fugue_aws_types`
data source. Specify a region when creating the data source.

For example:

```hcl
data "fugue_aws_types" "all" {
  region = "us-east-1"
}
```

You may then retrieve the list of types using `data.fugue_aws_types.all.types`.