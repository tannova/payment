variable "region" {
  default = "ap-southeast-1"
}
variable "profile" {
  default = "default"
}

variable "environment" {

}
variable "cluster_name" {
  
}

variable "rules" {
  type = list(object({
    api_groups = list(string)
    resources  = list(string)
    verbs      = list(string)
  }))
  default = []
}

variable "subjects" {
  type = list(object({
    kind      = string
    name      = string
    api_group = string
    namespace = string 
  }))
}

variable "name" {
  
}
variable "namespace" {
  
}