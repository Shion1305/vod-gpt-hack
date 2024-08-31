resource "aws_iam_role" "ecs_task_role" {
  name_prefix = "${local.prefix}-ecs-task-role-"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role" "ecs_task_execution_role" {
  name_prefix = "${local.prefix}-ecs-task-exec-role-"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}


resource "aws_ecs_task_definition" "ecs_task" {
  family                   = "${local.prefix}-example-task-api-server"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  // Replace definition with real application
  container_definitions = jsonencode([
    {
      name = "api-main-container"
      # サンプル用のイメージ
      image     = "${aws_ecr_repository.api-main.repository_url}:latest"
      cpu       = 256
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      environment = [
        {
          name  = "ENV_LOCATION"
          value = "/etc/setting.yaml"
        },
        {
          name  = "SQS_URL"
          value = aws_sqs_queue.transcribe-sqs.url
        }
      ]
      secrets = [
        {
          name      = "ENV_CONTENT"
          valueFrom = aws_secretsmanager_secret.api-main-yaml-credential.arn
        }
      ]
    }
  ])
}

