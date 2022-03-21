module "security-group" {
  source                                = "terraform-aws-modules/security-group/aws"
  version = "3.18.0"
  name                                  = var.name
  ingress_with_cidr_blocks              = var.ingress_with_cidr_blocks
  ingress_with_source_security_group_id = var.ingress_with_source_security_group_id
  tags = {
    "Environment" = var.environment
  }
  egress_with_cidr_blocks = var.egress_with_cidr_blocks
  vpc_id                  = var.vpc_id
}
