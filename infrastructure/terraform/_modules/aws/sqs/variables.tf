variable "name" {}
variable "environment" {}
variable "fifo_queue" {
  default = false
}
variable "content_based_deduplication" {
  default = false
}