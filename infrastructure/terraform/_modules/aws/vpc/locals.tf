locals {
  private_subnets = var.private_subnets == null ? [cidrsubnet(var.vpc_cidr, 8, 1), cidrsubnet(var.vpc_cidr, 8, 2), cidrsubnet(var.vpc_cidr, 8, 3)] : var.private_subnets
  public_subnets  = var.public_subnets == null ? [cidrsubnet(var.vpc_cidr, 8, 4), cidrsubnet(var.vpc_cidr, 8, 5), cidrsubnet(var.vpc_cidr, 8, 6)] : var.public_subnets
}
