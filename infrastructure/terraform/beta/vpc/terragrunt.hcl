include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../_modules/aws/vpc"
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment  = local.vars.environment
  profile      = local.vars.profile
  region       = local.vars.region
  cluster_name = local.vars.cluster_name
  vpc_cidr     = local.vars.vpc_cidr
}
