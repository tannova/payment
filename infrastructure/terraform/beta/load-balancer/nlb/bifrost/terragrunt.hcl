include {
  path = find_in_parent_folders("global.hcl")
}

terraform {
  source = "../../../../_modules/aws/elb/nlb"
}

dependency "vpc" {
  config_path = "../../../vpc"

  mock_outputs = {
    vpc_id = ""

    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}

dependency "targetgroup" {
  config_path = "../../targetgroup/bifrost"

  mock_outputs = {
    arn = ""

    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment         = local.vars.environment
  profile             = local.vars.profile
  region              = local.vars.region
  cluster_name        = local.vars.cluster_name
  vpc_id              = dependency.vpc.outputs.vpc_id
  internal            = local.vars.internal_access
  domain              = local.vars.bifrost_domain
  route53_hosted_zone = local.vars.route53_hosted_zone
  name                = "bifrost"
  target_group_arn    = dependency.targetgroup.outputs.arn
}
