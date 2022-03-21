output "db_instance_address" {
  description = "The address of the RDS instance"
  value = module.db.this_db_instance_address
}

output "db_instance_endpoint" {
  description = "The connection endpoint"
  value = module.db.this_db_instance_endpoint
}

output "db_instance_port" {
  description = "The database port"
  value       = module.db.this_db_instance_port
}
