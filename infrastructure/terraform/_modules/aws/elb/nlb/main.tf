module "acm" {
  source              = "./acm"
  domain              = var.domain
  route53_hosted_zone = var.route53_hosted_zone
  profile             = var.profile
  environment         = var.environment
  region              = var.region
  providers = {
    aws = aws
   }
}

data "aws_subnet_ids" "subnets" {
  vpc_id = var.vpc_id
  tags = {
    Tier = var.internal ? "Private" : "Public"
  }
}

resource "aws_lb" "this" {
  name               = local.name
  internal           = var.internal
  load_balancer_type = "network"
  subnets            = data.aws_subnet_ids.subnets.ids

  enable_deletion_protection = false

  tags = {
    Environment = var.environment
  }
}

resource "aws_lb_listener" "this" {
  load_balancer_arn = aws_lb.this.arn
  port              = var.listener.port
  protocol          = var.listener.protocol
  certificate_arn   = module.acm.certificate_arn
  alpn_policy       = var.listener.alpn_policy

  default_action {
    type             = "forward"
    target_group_arn = var.target_group_arn
  }
}

data "aws_route53_zone" "selected" {
  name         = var.route53_hosted_zone
  private_zone = false
}

resource "aws_route53_record" "this" {
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = var.domain
  type    = "A"

  alias {
    name                   = aws_lb.this.dns_name
    zone_id                = aws_lb.this.zone_id
    evaluate_target_health = true
  }
}
