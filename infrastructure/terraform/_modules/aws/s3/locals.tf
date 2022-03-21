locals {
  s3_origin_id            = "S3-${var.bucket}"
  source                  = templatefile("${path.module}/src/lambda-basic-auth.js", { username = var.basic_auth.username, password = var.basic_auth.password })
  basic_auth_count        = var.basic_auth.enabled ? 1 : 0
  cloudfront_count        = var.cloudfront.enabled ? 1 : 0
  cloudfront_domain_count = var.cloudfront.enabled && var.cloudfront.domain != "" ? 1 : 0
  s3_domain = var.website_mode ? aws_s3_bucket.this.website_endpoint : aws_s3_bucket.this.bucket_regional_domain_name
  function_name = "${var.bucket}-basic-auth"
}
