# Installation steps for Vault-secrets-operator


0. install vault via eks-bootstrap-apps
1. unseal vault, login vault via generated root token, then create vault admin token.
2. vault login with admin token, manually create vault policy for vault-secrets-operator

```
path "kv/*" {
  capabilities = ["read"]
}

path "kv/data/*" {
  capabilities = ["read"]
}
```
3. manually create vault token for vault-secrets-operator

```shell script
$ vault token create -policy vault-secrets-operator
```
4. kv put vault-secret-operator token into path "kv/vault-secrets-operator", under key "token".

5. run vault-bootstrap-apps/terragrunt apply

