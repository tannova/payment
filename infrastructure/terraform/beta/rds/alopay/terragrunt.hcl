include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../../_modules/aws/rds"
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
dependencies {
  paths = ["../../vault"]
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment          = local.vars.environment
  profile              = local.vars.profile
  region               = local.vars.region
  port                 = 3306
  size_in_gb           = 10
  engine               = "mysql"
  engine_version       = "5.7.19"
  major_engine_version = "5.7"
  identifier           = "alopay-beta"
  parameters = [
    {
      name  = "character_set_client"
      value = "utf8"
    },
    {
      name  = "character_set_server"
      value = "utf8"
    }
  ]
  vault_credentials_path = "kv/aws/rds"
  options                = []
  family                 = "mysql5.7"
  vpc_id                 = dependency.vpc.outputs.vpc_id
  security_group_id      = dependency.sg.outputs.security_group_id
}
