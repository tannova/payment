include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../_modules/k8s/datadog"
}
locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}