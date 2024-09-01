resource "aws_s3_bucket" "vod-store" {
  bucket = "progate-vod-gpt-store"
  tags = {
    Name = "progate-vod-gpt-store"
  }
}

resource "aws_s3_bucket_policy" "vod_store_policy" {
  bucket = aws_s3_bucket.vod-store.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action    = "s3:GetObject"
        Effect    = "Allow"
        Resource  = "${aws_s3_bucket.vod-store.arn}/*"
        Principal = "*"
      }
    ]
  })
}
