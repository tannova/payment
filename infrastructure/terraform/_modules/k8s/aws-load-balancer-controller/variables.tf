variable "cluster_name" {}

variable "region" {
    default = "ap-southeast-1"
}
variable "profile" {
    default = "default"
}
variable "oidc_provider_arn" {
  description = "EKS oidc_provider_arn"
}