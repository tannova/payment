resource "kubernetes_role" "role" {
  metadata {
    name      = var.name
    namespace = var.namespace
  }
  dynamic "rule" {
    for_each = var.rules
    content {
      api_groups = rule.value["api_groups"]
      resources  = rule.value["resources"]
      verbs      = rule.value["verbs"]
    }
  }
  depends_on = [
    kubernetes_namespace.this
  ]
}


data "kubernetes_all_namespaces" "allns" {}

locals {
  ns_exist = contains(data.kubernetes_all_namespaces.allns.namespaces, var.name)
}
resource "kubernetes_namespace" "this" {
  count = local.ns_exist ? 0 : 1
  metadata {
    name = var.namespace
  }
}

resource "kubernetes_role_binding" "role_binding" {
  metadata {
    name      = "${var.name}-role-binding"
    namespace = var.namespace
  }
  role_ref {
    name      = var.name
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
  }
  dynamic "subject" {
    for_each = var.subjects
    content {
      kind      = subject.value["kind"]
      name      = subject.value["name"]
      api_group = subject.value["api_group"]
      namespace = subject.value["namespace"]
    }
  }
  depends_on = [
    kubernetes_namespace.this
  ]
}
