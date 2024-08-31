data "aws_acm_certificate" "main_cert" {
  domain   = "vod.progate.shion.pro"
  statuses = ["ISSUED"]
}
