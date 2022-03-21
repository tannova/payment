##############################################################
# Data sources to get VPC, subnets and security group details
##############################################################
data "aws_vpc" "selected" {
  id = var.vpc_id
}

data "aws_subnet_ids" "all" {
  vpc_id = data.aws_vpc.selected.id
}

data vault_generic_secret "rds" {
  path = var.vault_credentials_path
}

module "db" {
  # depends_on = [aws_security_group.rds_sg]
  source = "terraform-aws-modules/rds/aws"

  version = "2.35.0"

  identifier = var.identifier

  engine            = var.engine
  engine_version    = var.engine_version
  instance_class    = var.instance_class
  allocated_storage = var.size_in_gb
  storage_encrypted = false
  port     = var.port

  username = data.vault_generic_secret.rds.data["username"]
  password = data.vault_generic_secret.rds.data["password"]

  vpc_security_group_ids = [local.security_group_id]

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  multi_az = var.multi_az

  # disable backups to create DB faster
  backup_retention_period = 0

  tags = {
    Owner       = "inspirelab"
    Environment = var.environment
  }

  enabled_cloudwatch_logs_exports = ["audit", "general"]

  # DB subnet group
  subnet_ids = data.aws_subnet_ids.all.ids

  # DB parameter group
  family = var.family

  # DB option group
  major_engine_version = var.major_engine_version

  # Snapshot name upon DB deletion
  final_snapshot_identifier = var.identifier

  # Database Deletion Protection
  deletion_protection = false

  parameters = var.parameters

  options = var.options
}

