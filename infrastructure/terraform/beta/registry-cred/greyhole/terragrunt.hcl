include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../../_modules/k8s/registry-cred"
}

dependencies {
  paths = ["../vault"]
}

locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment  = local.vars.environment
  profile      = local.vars.profile
  region       = local.vars.region
  cluster_name = local.vars.cluster_name
  vault_credentials_path = "kv/docker-config/greyhole"
  registry_credential_name = "greyhole"
  kubeconfig   = ""
}
