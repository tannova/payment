variable "profile" {

}

variable "region" {

}

variable "bucket" {

}

variable "website_mode" {
  type = bool
  default = false
  description = "Enable s3 staic website host"
}

variable "environment" {

}


variable "index_document" {
  default = "index.html"
}

variable "error_document" {
  default = "error.html"
}

variable "bucket_acl" {
  default     = "private"
  description = "The canned ACL to apply. Valid values are private, public-read, public-read-write, aws-exec-read, authenticated-read, and log-delivery-write. Defaults to private. Conflicts with grant"
}

variable "cors_rules" {
  type = list(object({
    allowed_headers = list(string)
    allowed_methods = list(string)
    allowed_origins = list(string)
    expose_headers  = list(string)
    max_age_seconds = number
  }))
  default = [
    {
      allowed_headers = []
      allowed_methods = ["GET"]
      allowed_origins = ["*"]
      expose_headers  = []
      max_age_seconds = 0
    }
  ]
}


variable "basic_auth" {
  type = object({
    enabled  = bool,
    username = string
    password = string
  })
  default = {
    enabled  = false
    username = ""
    password = ""
  }
}

variable "cloudfront" {
  type = object({
    enabled             = bool,
    domain              = string,
    route53_hosted_zone = string
    cache_policy_id     = string
  })
  default = {
    enabled             = true
    domain              = ""
    route53_hosted_zone = ""
    cache_policy_id     = "658327ea-f89d-4fab-a63d-7e88639e58f6"
  }
}

