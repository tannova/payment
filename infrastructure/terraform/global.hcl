locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}

remote_state {
  backend = "s3"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite"
  }
  config = {
    # Need to change if we apply for another account
    bucket = "alopay-eks-tf-state"
    region =  local.vars.region
    encrypt = true
    key    = "${path_relative_to_include()}/terraform.tfstate"
    profile = local.vars.profile
  }
}
