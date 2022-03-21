# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster

data "aws_subnet_ids" "private-subnets" {
  vpc_id = var.vpc_id
  tags = {
    Tier = "Private"
  }
}

data "aws_vpc" "selected" {
  id = var.vpc_id
}

resource "aws_elasticache_subnet_group" "ec_subnet" {
  name       = "${var.name}-ec-subnet-group"
  subnet_ids = data.aws_subnet_ids.private-subnets.ids
}

resource "aws_elasticache_cluster" "elasticcache_cluster" {
  cluster_id           = var.name
  engine               = var.engine
  node_type            = var.node_type
  num_cache_nodes      = var.num_cache_nodes
  parameter_group_name = var.parameter_group_name
  engine_version       = var.engine_version
  port                 = var.port
  tags = {
    Environment = var.environment
  }
  security_group_ids = [local.security_group_id]
  subnet_group_name  = aws_elasticache_subnet_group.ec_subnet.name
}
