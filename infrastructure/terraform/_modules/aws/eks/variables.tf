variable "environment" {}
variable "region" {}
variable "cluster_name" {}

variable "profile" {
  default     = "default"
  description = " aws profile"
}

variable "cluster_version" {
  default     = "1.19"
  description = "Kubernetes version to use for the EKS cluster."
}

variable "ssh_key_name" {
  default     = ""
  description = "The key pair name that should be used for the instances in the autoscaling group"
}

variable "security_group_ids" {
  default     = []
  description = "additional_security_group_ids for worker node"
  type = list(string)
}

variable "vpc_id" {
  
}

variable "map_roles" {
  description = "Additional IAM roles to add to the aws-auth configmap. See examples/basic/variables.tf for example format."
  type = list(object({
    rolearn  = string
    username = string
    groups   = list(string)
  }))
  default = []
}

variable "map_users" {
  description = "Additional IAM users to add to the aws-auth configmap. See examples/basic/variables.tf for example format."
  type = list(object({
    userarn  = string
    username = string
    groups   = list(string)
  }))
  default = []
}
variable "worker_groups" {
  description = <<EOF
  A list of maps defining worker group configurations to be defined using AWS Launch Configurations. 
  See workers_group_defaults for valid keys.
  https://github.com/terraform-aws-modules/terraform-aws-eks/blob/master/local.tf
  EOF
  type        = list(any)
  default     = []
}