include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../_modules/k8s/aws-load-balancer-controller"
}

dependency "cluster" {
  config_path = "..//eks-cluster"

  mock_outputs = {
    cluster_name = "eks"

    kubectl_config = "---"

    region = "ap-southeast-1"

    profile = ""

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
  oidc_provider_arn= dependency.cluster.outputs.oidc_provider_arn
}