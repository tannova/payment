# resource "kubernetes_secret" "elastic-certificates" {
#   metadata {
#     name      = "elastic-certificates"
#     namespace = "elk"
#   }
#   binary_data = {
#     "elastic-certificates.p12" = var.elastic_certificate_p12
#   }
#   type = "Opaque"
#   depends_on = [
#     kubernetes_namespace.elk
#   ]
# }

# resource "kubernetes_secret" "elastic-certificate-pem" {
#   metadata {
#     name      = "elastic-certificate-pem"
#     namespace = "elk"
#   }
#   data = {
#     "elastic-certificate.pem" = var.elastic_certificate_pem
#   }
#   type = "Opaque"
#   depends_on = [
#     kubernetes_namespace.elk
#   ]
# }

# resource "kubernetes_secret" "elastic-certificate-crt" {
#   metadata {
#     name      = "elastic-certificate-crt"
#     namespace = "elk"
#   }

#   binary_data = {
#     "elastic-certificate.crt" = var.elastic_certificate_crt
#   }
#   type = "Opaque"
#   depends_on = [
#     kubernetes_namespace.elk
#   ]
# }

# resource "kubernetes_config_map" "elastic-certificates" {
#   metadata {
#     name      = "elastic-certificates"
#     namespace = "elk"
#   }
#   binary_data = {
#     "elastic-certificates.p12" = var.elastic_certificate_p12
#     "elastic-certificate.crt"  = var.elastic_certificate_crt
#     "elastic-certificate.pem"  = var.elastic_certificate_pem
#   }
#   depends_on = [
#     kubernetes_namespace.elk
#   ]

# }
# resource "random_password" "password" {
#   length           = 16
#   special          = true
#   override_special = "_%@"
# }

# resource "kubernetes_secret" "elastic-credentials" {
#   metadata {
#     name      = "elastic-credentials"
#     namespace = "elk"
#   }
#   data = {
#     "username" = "elastic"
#     "password" = random_password.password.result
#   }
#   type = "Opaque"
# }

# resource "kubernetes_namespace" "elk" {
#   metadata {
#     name = "elk"
#   }
# }

# resource "random_password" "kibana" {
#   length           = 32
#   special          = true
#   override_special = "_%@"
# }



# resource "kubernetes_secret" "kibana" {
#   metadata {
#     name      = "kibana"
#     namespace = "elk"
#   }
#   data = {
#     "encryptionkey" = random_password.kibana.result
#   }
#   depends_on = [
#     null_resource.generate-cert
#   ]
# }
