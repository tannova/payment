
data "template_file" "values" {
  template = file("${path.module}/templates/external-dns-values.yaml")
  vars = {
    role_arn = aws_iam_role.external-dns-role.arn
    region   = var.region
  }
}


#https://artifacthub.io/packages/helm/bitnami/external-dns
resource "helm_release" "external-dns" {
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "external-dns"
  name       = "external-dns"
  version    = "4.11.0"

  namespace        = "kube-system"
  create_namespace = true
  atomic           = true
  wait             = true
  values           = [data.template_file.values.rendered]
}
