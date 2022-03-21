
variable "region" {
  default = "ap-southeast-1"
}
variable "profile" {
  default     = "default"
  description = "aws profile"
}
variable "environment" {

}
variable "name" {

}
variable "description" {
  default = "Security group description."
}

variable "vpc_id" {

}

variable "ingress_with_cidr_blocks" {
  description = "List of ingress rules to create where 'cidr_blocks' is used"
  type        = list(map(string))
  default     = []
}

variable "ingress_with_source_security_group_id" {
  description = "List of ingress rules to create where 'source_security_group_id' is used"
  type        = list(map(string))
  default     = []
}

variable "egress_with_cidr_blocks" {
  default = [{
    from_port = "-1"
    to_port ="-1"
    protocol ="-1"
    cidr_blocks = "0.0.0.0/0"
    description = "Allow nodes all egress to the Internet."
  }]
  description = "List of egress rules to create where 'cidr_blocks' is used"
}
