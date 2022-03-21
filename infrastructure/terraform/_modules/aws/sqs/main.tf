terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 3.14.0"
    }
  }
}

resource "aws_sqs_queue" "terraform_queue" {
  name                      = var.name
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.terraform_queue_deadletter.arn
    maxReceiveCount     = 4
  })
  fifo_queue = var.fifo_queue
  content_based_deduplication = var.content_based_deduplication

  tags = {
    Environment = var.environment
  }
}