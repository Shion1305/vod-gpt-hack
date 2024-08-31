resource "aws_sqs_queue" "transcribe-sqs" {
  name                        = "${local.prefix}-transcribe-sqs"
  fifo_queue                  = true
  content_based_deduplication = true
}
