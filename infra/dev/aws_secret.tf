resource "aws_secretsmanager_secret" "api-main-yaml-credential" {
  name = "api-main-yaml-credential"
}

resource "aws_secretsmanager_secret_version" "api-main-yaml-credential" {
  secret_id = aws_secretsmanager_secret.api-main-yaml-credential.id
  // read from file
  secret_string = file("./api-main-credential.yaml")

  lifecycle {
    create_before_destroy = true
    ignore_changes        = [secret_string]
  }
}

resource "aws_kms_key" "main-key" {
  description = "KMS key for my_secret"
}
