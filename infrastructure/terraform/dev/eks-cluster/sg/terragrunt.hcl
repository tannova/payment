include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../../_modules/aws/security-group"
}
dependency "vpc" {
  config_path = "../../vpc"
  mock_outputs = {
    vpc_id                                  = ""
    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}

locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment  = local.vars.environment
  profile      = local.vars.profile
  region       = local.vars.region
  cluster_name = local.vars.cluster_name
  vpc_id       = dependency.vpc.outputs.vpc_id
  ingress_with_cidr_blocks = [
    {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      description = "Allow ssh for internal access."
      cidr_blocks = local.vars.vpc_cidr
    }
  ]
  name = "dev-worker-node-sg"
}
