resource "aws_elasticsearch_domain" "this" {
  domain_name           = var.name
  elasticsearch_version = var.elasticsearch_version

  cluster_config {
    instance_count = var.instance_count
    instance_type  = var.instance_type
    //To enable multi-az
    zone_awareness_enabled = var.instance_count >= 2 ? true : false
  }

  snapshot_options {
    automated_snapshot_start_hour = 00
  }

  vpc_options {
    subnet_ids         = var.instance_count >= 2 ? data.aws_subnet_ids.selected.ids : [sort(data.aws_subnet_ids.selected.ids)[0]]
    security_group_ids = var.security_group_ids
  }

  tags = {
    Domain      = var.name,
    Environment = var.environment
  }

  ebs_options {
    ebs_enabled = true
    volume_size = 10
    volume_type = "gp2"
  }

  advanced_security_options {
    enabled                        = true
    internal_user_database_enabled = true
    master_user_options {
      master_user_name     = "admin"
      master_user_password = random_password.master_user_password.result
    }
  }
  domain_endpoint_options {
    enforce_https       = true
    tls_security_policy = "Policy-Min-TLS-1-2-2019-07"
  }

  node_to_node_encryption {
    enabled = true
  }
  encrypt_at_rest {
    enabled = true
  }
}

resource "random_password" "master_user_password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-=+[]{}<>?" //Removed  _,@,:
}


data "aws_subnet_ids" "selected" {
  vpc_id = var.vpc_id

  tags = {
    Tier = "Private"
  }
}
