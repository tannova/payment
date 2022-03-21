output "ec-configuration-endpoint" {
  value = aws_elasticache_cluster.elasticcache_cluster.configuration_endpoint
}

output "ec-cluster-address" {
  value = aws_elasticache_cluster.elasticcache_cluster.cluster_address
}

output "ec-cache-nodes" {
  value = aws_elasticache_cluster.elasticcache_cluster.cache_nodes
}
