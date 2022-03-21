data "template_file" "datadog-values" {
  template = file("${path.module}/templates/datadog.yaml")
}

data vault_generic_secret "datadog" {
  path = "kv/api/datadog"
}

resource "helm_release" "datadog_agent" {
  name       = "datadog-agent"
  chart      = "datadog"
  repository = "https://helm.datadoghq.com"
  version    = "2.10.1"

  namespace        = "monitor"
  create_namespace = true
  atomic           = true
  wait             = true
  values = [
    data.template_file.datadog-values.rendered
  ]

  set_sensitive {
    name  = "datadog.apiKey"
    value = data.vault_generic_secret.datadog.data["api_key"]
  }
}