data "template_file" "elasticsearch-values" {
  template = file("${path.module}/templates/elasticsearch.yaml")
  vars = {
  }
}

data "template_file" "filebeat-values" {
  template = file("${path.module}/templates/filebeat.yaml")
  vars = {
  }
}

data "template_file" "kibana-values" {
  template = file("${path.module}/templates/kibana.yaml")
  vars = {
    domain   = var.domain
    cert_arn = aws_acm_certificate.domain.arn
  }
}

data "template_file" "generate-cert" {
  template = file("${path.module}/cert/generate.sh")
  vars = {
    profile      = var.profile
    cluster_name = var.cluster_name
  }
}
resource "null_resource" "generate-cert" {
  provisioner "local-exec" {
    command = data.template_file.generate-cert.rendered
  }
}

resource "helm_release" "elasticsearch" {
  repository = "https://Helm.elastic.co"
  chart      = "elasticsearch"
  name       = "elasticsearch"
  version    = "7.11.1"

  namespace        = "elk"
  create_namespace = true
  atomic           = true
  wait             = true
  values = [
    data.template_file.elasticsearch-values.rendered
  ]
  set {
    name  = "replicas"
    value = 3
  }

  depends_on = [
    # kubernetes_secret.elastic-certificates,
    # kubernetes_secret.elastic-certificate-pem,
    # kubernetes_secret.elastic-certificate-crt
    # kubernetes_config_map.elastic-certificates
    null_resource.generate-cert
  ]
}

resource "helm_release" "kibana" {
  repository = "https://Helm.elastic.co"
  chart      = "kibana"
  name       = "kibana"
  version    = "7.11.1"

  namespace        = "elk"
  create_namespace = true
  atomic           = true
  wait             = true
  values = [
    data.template_file.kibana-values.rendered
  ]
  set {
    name  = "resources.requests.cpu"
    value = "500m"
  }
  set {
    name  = "resources.requests.memory"
    value = "1Gi"
  }
  set {
    name  = "resources.limits.cpu"
    value = "500m"
  }
  set {
    name  = "resources.limits.memory"
    value = "1Gi"
  }
  set {
    name  = "service.type"
    value = "NodePort"
  }
  set {
    name  = "ingress.enabled"
    value = "true"
  }
  set {
    name  = "ingress.hosts[0]"
    value = var.domain
  }
  set {
    name  = "ingress.tls[0].hosts[0]"
    value = var.domain
  }
  depends_on = [
    helm_release.elasticsearch,
    null_resource.generate-cert
  ]
}

resource "helm_release" "filebeat" {
  repository = "https://Helm.elastic.co"
  chart      = "filebeat"
  name       = "filebeat"
  version    = "7.11.1"

  namespace        = "elk"
  create_namespace = true
  atomic           = true
  wait             = true
  values = [
    data.template_file.filebeat-values.rendered
  ]
  depends_on = [
    helm_release.elasticsearch,
    null_resource.generate-cert
  ]
}

