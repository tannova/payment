include {
  path = find_in_parent_folders("global.hcl")
}
terraform {
  source = "../../_modules/aws/eks"
}
dependency "vpc" {
  config_path = "../vpc"
  mock_outputs = {
    vpc_id                                  = ""
    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}

dependency "sg" {
  config_path = "./sg"
  mock_outputs = {
    security_group_id                       = ""
    mock_outputs_allowed_terraform_commands = ["validate", "plan"]
  }
}

locals {
  vars = yamldecode(file(find_in_parent_folders("variables.yaml")))
}
inputs = {
  environment        = local.vars.environment
  profile            = local.vars.profile
  region             = local.vars.region
  cluster_name       = local.vars.cluster_name
  vpc_id             = dependency.vpc.outputs.vpc_id
  cluster_version    = "1.19"
  security_group_ids = [dependency.sg.outputs.security_group_id]
  map_roles = [
    {
      rolearn  = "arn:aws:iam::901552523807:role/gitlab-runner"
      username = "{{SessionName}}"
      groups   = ["system:masters"]
    },
    {
      rolearn  = "arn:aws:iam::901552523807:role/AWSReservedSSO_Alopay_9c583575b964faa9"
      username = "{{SessionName}}"
      groups   = ["system:masters"]
    }
  ]
  worker_groups = [
    {
      name                          = "${local.vars.cluster_name}-worker-group-1"
      instance_type                 = local.vars.instance_type
      asg_desired_capacity          = 3
      asg_min_size                  = 3
      asg_max_size                  = 3
      root_volume_type              = "gp2"
      additional_security_group_ids = [dependency.sg.outputs.security_group_id]
      key_name                      = local.vars.ssh_key_name
      kubelet_extra_args            = "--node-labels=node.kubernetes.io/app=general",
      tags = [
        {
          key                 = "node.kubernetes.io/app"
          value               = "general"
          propagate_at_launch = "false"
        },
        {
          key                 = "k8s.io/cluster-autoscaler/enabled"
          propagate_at_launch = "false"
          value               = "false"
        },
        {
          key                 = "k8s.io/cluster-autoscaler/${local.vars.cluster_name}"
          propagate_at_launch = "false"
          value               = "owned"
        }
      ]
    }
  ]
}
