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
  family        = "${local.prefix}-example-task-api-server"
  network_mode  = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu           = "256"
  memory        = "512"
  task_role_arn = aws_iam_role.ecs_task_role.arn
  execution_role_arn = aws_iam_role.ecs_task_execution_role.arn

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
          value = "/app/setting.yaml"
        },
        {
          name  = "SQS_URL"
          value = aws_sqs_queue.transcribe-sqs.url
        },
        {
          name  = "S3_BUCKET"
          value = aws_s3_bucket.vod-store.bucket
        },
        {
          name  = "REGION"
          value = "us-west-2"
        }
      ]
    }
  ])
}

resource "aws_iam_policy" "secrets_manager_access_policy" {
  name        = "${local.prefix}-secrets-manager-access-policy"
  description = "Allow access to Secrets Manager"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ],
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_secrets_manager" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = aws_iam_policy.secrets_manager_access_policy.arn
}


resource "aws_iam_policy" "ecr_access_policy" {
  name        = "${local.prefix}-ecr-access-policy"
  description = "Allow ECS tasks to pull images from ECR"
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "ecr:GetAuthorizationToken"
        ],
        Resource = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_ecr_access" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = aws_iam_policy.ecr_access_policy.arn
}
