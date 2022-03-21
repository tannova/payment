resource "aws_s3_bucket" "this" {
  bucket = var.bucket
  acl    = var.bucket_acl

  dynamic "website" {
    for_each = var.website_mode ? [var.website_mode] : []
    content {
      index_document = var.index_document
    }
  }
  dynamic "cors_rule" {
    for_each = var.cors_rules
    content {
      allowed_headers = cors_rule.value["allowed_headers"]
      allowed_methods = cors_rule.value["allowed_methods"]
      allowed_origins = cors_rule.value["allowed_origins"]
      expose_headers  = cors_rule.value["expose_headers"]
      max_age_seconds = cors_rule.value["max_age_seconds"]
    }
  }
  tags = {
    "Environment" = var.environment
  }
}


resource "aws_cloudfront_origin_access_identity" "this" {
  count   = local.cloudfront_count
  comment = "Managed by terraform. Bucket ${var.bucket}"
}


data "aws_iam_policy_document" "s3_policy" {
  count = local.cloudfront_count
  statement {
    # S3 policy allow cloudfront access private S3 objcet.
    # Allow access for normal mode, not website_modes
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.this.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.this[0].iam_arn]
    }
  }
  dynamic "statement" {
    # S3 policy allows cloudfront with matched Referer header can access object.
    # It's created when website_mode is enabled
    for_each = var.website_mode ? [var.website_mode] : []
    content {
      actions   = ["s3:GetObject"]
      resources = ["${aws_s3_bucket.this.arn}/*"]
      principals {
        type        = "*"
        identifiers = ["*"]
      }
      condition {
        test     = "StringLike"
        values   = [aws_cloudfront_origin_access_identity.this[0].id]
        variable = "aws:Referer"
      }
    }
  }
}

resource "aws_s3_bucket_policy" "this" {
  # Assign policy to S3 bucket
  count  = local.cloudfront_count
  bucket = aws_s3_bucket.this.id
  policy = data.aws_iam_policy_document.s3_policy[0].json
}


resource "aws_cloudfront_function" "basic_auth" {
  # Basic authentication function, only create if var.basic_auth.enabled = true
  count   = local.basic_auth_count
  name    = local.function_name
  runtime = "cloudfront-js-1.0"
  comment = "Managed by terraform. Bucket ${var.bucket}"
  publish = true
  code    = local.source
}

resource "aws_cloudfront_distribution" "s3_distribution" {
  count   = local.cloudfront_count
  enabled = true
  origin {
    domain_name = local.s3_domain
    origin_id   = local.s3_origin_id
    dynamic "s3_origin_config" {
      # For s3 static resource only, website_mode disabled
      for_each = var.website_mode ? [] : [true]
      content {
        origin_access_identity = aws_cloudfront_origin_access_identity.this[0].cloudfront_access_identity_path
      }
    }
    dynamic "custom_origin_config" {
      # For website_mode enabled
      for_each = var.website_mode ? [true] : []
      content {
        http_port              = 80
        https_port             = 443
        origin_protocol_policy = "http-only"
        origin_ssl_protocols   = ["TLSv1.1", "TLSv1.2", "SSLv3"]
      }
    }
    dynamic "custom_header" {
      # This header supports access pirvate S3,
      # we will create a policy in S3 only allow access object if matching Referer header
      for_each = var.website_mode ? [true] : []
      content {
        name  = "Referer"
        value = aws_cloudfront_origin_access_identity.this[0].id
      }
    }
  }
  tags = {
    "Environment" = var.environment
  }
  aliases = local.cloudfront_domain_count > 0 ? [var.cloudfront.domain] : []
  viewer_certificate {
    acm_certificate_arn            = local.cloudfront_domain_count > 0 ? module.acm[0].certificate_arn : null
    cloudfront_default_certificate = local.cloudfront_domain_count > 0 ? false : true
    ssl_support_method             = local.cloudfront_domain_count > 0 ? "sni-only" : null
  }

  custom_error_response {
    error_caching_min_ttl = 300
    error_code            = 404
    response_code         = 200
    response_page_path    = "/${var.error_document}"
  }

  custom_error_response {
    error_caching_min_ttl = 300
    error_code            = 403
    response_code         = 200
    response_page_path    = "/${var.error_document}"
  }
  default_cache_behavior {
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = local.s3_origin_id
    cache_policy_id        = var.cloudfront.cache_policy_id
    dynamic "function_association" {
      # Setup basic authenticaiton function for cloudfront
      for_each = local.basic_auth_count == 0 ? [] : [var.basic_auth]
      content {
        event_type   = "viewer-request"
        function_arn = aws_cloudfront_function.basic_auth[0].arn
      }
    }

  }

  ordered_cache_behavior {
    # Disable cache for / index
    path_pattern           = "/"
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD", "OPTIONS"]
    target_origin_id       = local.s3_origin_id
    cache_policy_id        = "4135ea2d-6df8-44a3-9df3-4b5a84be39ad"
    viewer_protocol_policy = "redirect-to-https"
    dynamic "function_association" {
      # Setup basic authenticaiton function for cloudfront
      for_each = local.basic_auth_count == 0 ? [] : [var.basic_auth]
      content {
        event_type   = "viewer-request"
        function_arn = aws_cloudfront_function.basic_auth[0].arn
      }
    }
  }
  ordered_cache_behavior {
    # Disable cache for html files, for versioning support
    path_pattern           = "*.html"
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD", "OPTIONS"]
    target_origin_id       = local.s3_origin_id
    cache_policy_id        = "4135ea2d-6df8-44a3-9df3-4b5a84be39ad"
    viewer_protocol_policy = "redirect-to-https"
    dynamic "function_association" {
      # Setup basic authenticaiton function for cloudfront
      for_each = local.basic_auth_count == 0 ? [] : [var.basic_auth]
      content {
        event_type   = "viewer-request"
        function_arn = aws_cloudfront_function.basic_auth[0].arn
      }
    }
  }
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
}



data "aws_route53_zone" "selected" {
  name         = var.cloudfront.route53_hosted_zone
  private_zone = false
}

# Setup router 53 for cloudfront alias domain
resource "aws_route53_record" "cloudfront_domain" {
  count   = local.cloudfront_domain_count
  zone_id = data.aws_route53_zone.selected.zone_id
  name    = var.cloudfront.domain
  type    = "A"
  alias {
    name                   = aws_cloudfront_distribution.s3_distribution[0].domain_name
    zone_id                = aws_cloudfront_distribution.s3_distribution[0].hosted_zone_id
    evaluate_target_health = true
  }
}


# Create domain and SSL certificate
module "acm" {
  count               = local.cloudfront_domain_count
  source              = "./acm"
  domain              = var.cloudfront.domain
  route53_hosted_zone = var.cloudfront.route53_hosted_zone
  profile             = var.profile
  environment         = var.environment
  # Cloudfront only supports ACM from us-east-1 region
  region = "us-east-1"
  providers = {
    aws = aws.us-east-1
  }
}


