resource "aws_sqs_queue" "transcribe-sqs" {
  name_prefix                 = "${local.prefix}-transcribe-sqs"
  fifo_queue                  = true
  content_based_deduplication = true
}
