include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../../../_modules/aws/s3"
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  error_document = "index.html"
  environment = local.vars.environment
  profile     = local.vars.profile
  region      = local.vars.region
  bucket      = local.vars.cmgp_bucket
  bucket_acl  = "public-read"
  cloudfront = {
    enabled             = true
    domain              = local.vars.cmgp_domain
    route53_hosted_zone = local.vars.route53_hosted_zone
    cache_policy_id     = local.vars.cloudfront_cache_policy_id
  }
  website_mode = true
  basic_auth = {
    enabled  = true
    username = local.vars.cmgp_username
    password = local.vars.cmgp_password
  }
}
