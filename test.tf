# orca-iac ignore-block
resource "aws_s3_bucket" "data" {
  # bucket is public
  # bucket is not encrypted
  # bucket does not have access logs
  # bucket does not have versioning
  bucket        = "test-data"
  acl           = "public"
  force_destroy = true
}