resource "aws_acm_certificate" "domain" {
  domain_name       = var.domain
  validation_method = "DNS"
   tags = {
    Environment = var.environment
  }
}


data "aws_route53_zone" "selected" {
  name         =  var.route53_hosted_zone
  private_zone = false
}

resource "aws_route53_record" "domain" {
  for_each = {
    for dvo in aws_acm_certificate.domain.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = data.aws_route53_zone.selected.zone_id
}

resource "aws_acm_certificate_validation" "domain" {
  certificate_arn         = aws_acm_certificate.domain.arn 
  validation_record_fqdns = [for record in aws_route53_record.domain : record.fqdn]
}

data "template_file" "vault-values" {
  template = file("${path.module}/templates/values.yaml")
  vars = {
    domain   = var.domain
    cert_arn = aws_acm_certificate.domain.arn 
  }
}

resource "helm_release" "vault" {
  repository = "https://helm.releases.hashicorp.com"
  chart      = "vault"
  name       = "vault"
  version    = "0.8.0"

  namespace        = "vault"
  create_namespace = true
  atomic           = true
  wait             = true

  values = [
    data.template_file.vault-values.rendered
  ]
  set {
    name  = "injector.enabled"
    value = false
  }
  set {
    name  = "server.resources.requests.memory"
    value = "1Gi"
  }
  set {
    name  = "server.resources.requests.cpu"
    value = "500m"
  }
  set {
    name  = "server.resources.limits.memory"
    value = "1Gi"
  }
  set {
    name  = "server.resources.limits.cpu"
    value = "500m"
  }
  set {
    name  = "ui.enabled"
    value = true
  }
  set {
    name  = "server.service.type"
    value = "NodePort"
  }
  set {
    name  = "server.standalone.enabled"
    value = true
  }
  set {
    name  = "server.ingress.enabled"
    value = true
  }
  set {
    name  = "server.ingress.hosts[0].host"
    value = var.domain
  }
  set {
    name  = "server.ingress.tls[0].hosts[0]"
    value = var.domain
  }
}
