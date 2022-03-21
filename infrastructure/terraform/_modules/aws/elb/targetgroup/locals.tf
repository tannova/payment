locals {
  target_group_name = try(regex(var.cluster_name, var.name),null) == null ?format("%s-%s", var.cluster_name, var.name) : var.name
}
