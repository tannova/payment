terraform {
  required_providers {
    vault = {
      source  = "hashicorp/vault"
      version = "2.20.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "2.2.0"
    }
    template = {
      source  = "hashicorp/template"
      version = "2.2.0"
    }
  }
}
