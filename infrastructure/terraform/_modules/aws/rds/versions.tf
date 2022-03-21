provider "aws" {
  version = ">= 2.28.1"
  region  = var.region
  profile = var.profile
}

provider "vault" {
  //Recommend use system environment for vault provider to protect vault secret token
  //Please set VAULT_ADDR, VAULT_TOKEN env before running this module
}