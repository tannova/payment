variable "vpc_id" {}
variable "name" {}
variable "engine" {}
variable "num_cache_nodes" {}
variable "environment" {}
variable "engine_version" {}
variable "port" {}
variable "parameter_group_name" {}
variable "node_type" {}
# variable "elasticache_subnet_group_name" {}
variable "region" {
  default = "ap-southeast-1"
}
variable "profile" {
  default     = "default"
  description = "aws profile"
}

variable "security_group_id" {
  default     = null
  description = "Security group for RDS. if not set, will create new security group."
}
