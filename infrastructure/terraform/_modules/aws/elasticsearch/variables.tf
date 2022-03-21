variable "vpc_id" {}
variable "name" {}
variable "environment" {}
variable "instance_type" {
}
variable "instance_count" {
  default = 1
}
variable "region" {
  default = "ap-southeast-1"
}

variable "profile" {
  default     = "default"
  description = "aws profile"
}

variable "elasticsearch_version" {
  default = "7.10"
}


variable "security_group_ids" {
  type    = list(string)
  default = []
}
