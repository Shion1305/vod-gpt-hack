resource "aws_s3_bucket" "vod-store" {
  bucket = "progate-vod-gpt-store"
  tags = {
    Name = "progate-vod-gpt-store"
  }
}
