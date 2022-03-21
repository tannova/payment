//Create IAM policy https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html
# data "aws_iam_policy" "load-balancer-policy" {
#   name = "AWSLoadBalancerControllerIAMPolicy"
# }
locals {
  issuer = replace(data.aws_eks_cluster.cluster.identity[0].oidc[0].issuer, "https://", "")
}
resource "aws_iam_role" "eks-load-balancer-role" {
  name               = "${var.cluster_name}-load-balancer"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "${var.oidc_provider_arn}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${local.issuer}:sub": "system:serviceaccount:kube-system:aws-load-balancer-controller",
          "${local.issuer}:aud": "sts.amazonaws.com"
        }
      }
    }
  ]
}
EOF

  inline_policy {
    name   = "AWSLoadBalancerControllerIAMPolicy"
    policy = file("${path.module}/iam_policy.json")
  }

}


# resource "aws_iam_role_policy_attachment" "eks-load-balancer-role-attach" {
#   role = aws_iam_role.eks-load-balancer-role.arn
#   policy_arn = aws_iam_policy.arn
# }

resource "aws_iam_role" "external-dns-role" {
  name               = "${var.cluster_name}-external-dns"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "${var.oidc_provider_arn}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${local.issuer}:sub": "system:serviceaccount:kube-system:external-dns",
          "${local.issuer}:aud": "sts.amazonaws.com"
        }
      }
    }
  ]
}
EOF

  inline_policy {
    name   = "ExternalDNSIAMPolicy"
    policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "route53:ChangeResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ListHostedZones",
        "route53:ListResourceRecordSets"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
EOF
  }

}
