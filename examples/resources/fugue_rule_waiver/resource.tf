
resource "fugue_rule_waiver" "example" {
  name = "waive-FG_R00229"
  comment = "This S3 bucket is intentionally public"
  environment_id = fugue_aws_environment.test.id
  rule_id = "FG_R00229"
  resource_type = "AWS.S3.Bucket"
  resource_provider = "aws.us-east-1"
  resource_id = "my-public-s3-bucket"
}
