output "bucket_arn" {
  value = aws_s3_bucket.this.arn
}
output "bucket_regional_domain_name" {
  value = local.s3_domain
}


output "acm_arn" {
  value = local.cloudfront_domain_count == 0 ? "" : module.acm[0].certificate_arn
}
output "origin_access_identity_id" {
  value = local.cloudfront_count == 0 ? "" : aws_cloudfront_origin_access_identity.this[0].id
}

output "origin_access_identity_iam_arn" {
  value = local.cloudfront_count == 0 ? "" : aws_cloudfront_origin_access_identity.this[0].iam_arn
}

output "cloudfront_domain_alias" {
  value = local.cloudfront_count == 0 ? "" : local.cloudfront_domain_count == 0 ? aws_cloudfront_distribution.s3_distribution[0].domain_name : var.cloudfront.domain
}

output "cloudfront_domain" {
  value = local.cloudfront_count == 0 ? "" : aws_cloudfront_distribution.s3_distribution[0].domain_name
}
