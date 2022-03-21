include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../../_modules/aws/s3"
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment = local.vars.environment
  profile     = local.vars.profile
  region      = local.vars.region
  bucket      = local.vars.falcon_bucket
  bucket_acl  = "private"
  cloudfront = {
    enabled             = true
    domain              = local.vars.falcon_domain
    route53_hosted_zone = local.vars.route53_hosted_zone
    cache_policy_id     = local.vars.cloudfront_cache_policy_id
  }
}
