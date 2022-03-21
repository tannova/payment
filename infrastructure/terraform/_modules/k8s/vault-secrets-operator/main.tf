data "template_file" "vault-secrets-values" {
  template = file("${path.module}/templates/vault-secrets-values.yaml")
}

data "vault_generic_secret" "vso-token" {
  path = "kv/vault-secrets-operator/"
}

resource "kubernetes_secret" "vault-secrets-operator" {
  metadata {
    name = "vault-secrets-operator"
    namespace = "vault"
  }

  data = {
    "VAULT_TOKEN" = data.vault_generic_secret.vso-token.data["token"]
    "VAULT_TOKEN_LEASE_DURATION" = 86400
  }

  type = "Opaque"
}

resource "helm_release" "vault-secrets-operator" {
  depends_on = [kubernetes_secret.vault-secrets-operator]
  repository = "https://ricoberger.github.io/helm-charts"
  chart = "vault-secrets-operator"
  name = "vault-secrets-operator"
  version = "1.14.4"

  namespace         = "vault"
  create_namespace  = false
  atomic            = true
  wait              = true

  values = [
    data.template_file.vault-secrets-values.rendered
  ]
}
