
data "aws_subnet_ids" "private-subnets" {
  vpc_id = var.vpc_id
  tags = {
    Tier = "Private"
  }
}
module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  version         = "17.1.0"
  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version
  subnets         = data.aws_subnet_ids.private-subnets.ids

  tags = {
    Environment = var.environment
  }

  enable_irsa = true

  vpc_id = var.vpc_id

  worker_groups = var.worker_groups

  map_users = var.map_users
  map_roles = var.map_roles
  
}

data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}
