data "aws_acm_certificate" "main_cert" {
  domain   = "vod-gpt.gopher.jp"
  statuses = ["ISSUED"]
  types    = ["AMAZON_ISSUED"]
}
