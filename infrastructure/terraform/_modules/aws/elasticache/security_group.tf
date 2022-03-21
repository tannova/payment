resource "aws_security_group" "this" {
  count       = local.create ? 1 : 0
  name_prefix = "ec-${var.name}"
  vpc_id      = var.vpc_id

  ingress {
    from_port = var.port
    to_port   = var.port
    protocol  = "tcp"

    cidr_blocks = [
      data.aws_vpc.selected.cidr_block,
    ]
  }

  tags = {
    "Environment" = var.environment
    "Name"        = "ec-${var.name}"
  }
}

locals {
  create            = var.security_group_id == null
  security_group_id = local.create ? aws_security_group.this[0].id : var.security_group_id
}
