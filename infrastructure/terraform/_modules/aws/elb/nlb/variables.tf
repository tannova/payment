variable "domain" {

}

variable "route53_hosted_zone" {

}

variable "profile" {

}

variable "environment" {

}

variable "cluster_name" {

}

variable "region" {
  default = "ap-southeast-1"

}

variable "name" {

}

variable "vpc_id" {

}


variable "internal" {
  type = bool

}

variable "listener" {
  type = object({
    port        = string
    protocol    = string
    alpn_policy = string
  })
  default = {
    port        = "443"
    protocol    = "TLS"
    alpn_policy = "HTTP2Preferred"
  }

}
variable "target_group_arn" {
  
}