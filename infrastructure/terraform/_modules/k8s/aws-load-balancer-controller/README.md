# AWS Load Balancer Controller
https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html

Skip Step 1
```
Determine whether you have an existing IAM OIDC provider for your cluster.
View your cluster's OIDC provider URL.
```

Start from Step 2
```
Download an IAM policy for the AWS Load Balancer Controller that allows it to make calls to AWS APIs on your behalf. You can view the policy document on GitHub.
```

## NOTED
If you want to have ingress ALB then you need to config sevice.Type: NodePort