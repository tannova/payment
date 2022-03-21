terraform {
  required_providers {
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
  }
}

data "vault_generic_secret" "docker-config" {
  path = var.vault_credentials_path
}

data "template_file" "docker-config" {
  template = file("${path.module}/templates/config.json")
  vars = {
    username = data.vault_generic_secret.docker-config.data["username"]
    password = data.vault_generic_secret.docker-config.data["password"]
    auth = data.vault_generic_secret.docker-config.data["auth"]
  }
}

resource "kubernetes_secret" "registry-credentials" {
  metadata {
    name = var.registry_credential_name
    namespace = "default"
  }
  data = {
    ".dockerconfigjson" = data.template_file.docker-config.rendered
  }
  type = "kubernetes.io/dockerconfigjson"
}
