
variable "environment" {}
variable "region" {}
variable "cluster_name" {}


variable "vpc_cidr" {
  default = "10.0.0.0/16"
  description = "CIDR uses in EKS cluster"
}

variable "profile" {
  default = "default"
  description = "aws profile"
}

variable "private_subnets" {
  default = null
  type = list(string)
}

variable "public_subnets" {
  default = null
  type = list(string)
}