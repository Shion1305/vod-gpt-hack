variable "aws_access_key" {
  type = string
}

variable "aws_secret_key" {
  type = string
}

variable "aws_session_token" {
  type = string
}

locals {
  prefix = "pg"
  tags = {
    "Terraform" = "true"
    "Project"   = "progate_20240831_hackathon"
  }
}
