include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../..//_modules/k8s/vault"
}


dependency "load-balancer"{
  config_path = "../load-balancer"
  mock_outputs = {
    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}

locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment = local.vars.environment
  profile = local.vars.profile
  region = local.vars.region
  cluster_name = local.vars.cluster_name
  domain = local.vars.vault_domain
  route53_hosted_zone = local.vars.route53_hosted_zone
}