provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

resource "aws_cloudfront_distribution" "domain_front" {
  enabled = "${var.cloudfront_enabled}"

  origin {
    domain_name = "${var.cloudfront_origin}"
    origin_id   = "hidensneak-${var.cloudfront_origin}"

    custom_origin_config {
      http_port              = 14380
      https_port             = 443
      origin_protocol_policy = "match-viewer"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
  }

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "hidensneak-${var.cloudfront_origin}"

    forwarded_values {
      headers = ["*"]

      cookies {
        forward = "all"
      }

      query_string = true
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0

    smooth_streaming = false
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  price_class = "PriceClass_All"

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  lifecycle {
    create_before_destroy = true
  }
}
