variable "engine" {
  default = "mysql"
}
variable "engine_version" {
  default = "8.0.21"
}
variable "instance_class" {
  default = "db.t2.medium"
}
variable "port" {
  default = 3306
}
variable "size_in_gb" {}
# variable "username" {}
# variable "password" {}
variable "environment" {}
variable "identifier" {}
variable "options" {
  description = "A list of Options to apply."
  type        = any
  default     = []
}
variable "parameters" {
  description = "A list of DB parameters (map) to apply"
  type        = list(map(string))
  default     = []
}
variable "family" {
  description = "The family of the DB parameter group"
  type        = string
  default     = ""
}
variable "major_engine_version" {
  description = "Specifies the major version of the engine that this option group should be associated with"
  type        = string
  default     = ""
}
variable "vpc_id" {}

variable "region" {
  default = "ap-southeast-1"
}
variable "profile" {
  default     = "default"
  description = "aws profile"
}

variable "vault_credentials_path" {
  default = "kv/aws/rds"
}
variable "security_group_id" {
  default     = null
  description = "Security group for RDS. if not set, will create new security group."
}

variable "multi_az" {
  default = false
}