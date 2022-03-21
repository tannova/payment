include {
  path = find_in_parent_folders("global.hcl")
}

terraform {
  source = "../../../_modules/aws/elasticache"
}

dependency "vpc" {
  config_path = "../../vpc"

  mock_outputs = {
    vpc_id = ""

    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}
dependency "sg" {
  config_path = "../sg"
  mock_outputs = {
    security_group_id                       = ""
    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment = local.vars.environment
  profile     = local.vars.profile
  region      = local.vars.region
  vpc_id      = dependency.vpc.outputs.vpc_id

  // elasticache
  name                 = "alopay-shield-beta"
  port                 = 6379
  engine               = "redis"
  num_cache_nodes      = 1
  engine_version       = "6.x"
  parameter_group_name = "default.redis6.x"
  node_type            = "cache.t2.micro"
  security_group_id    = dependency.sg.outputs.security_group_id
}
