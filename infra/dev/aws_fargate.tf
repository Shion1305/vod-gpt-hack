resource "aws_ecs_cluster" "main" {
  name = "${local.prefix}-ecs-cluster"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecs_service" "fargate_service" {
  name            = "fargate-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.ecs_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.public_subnet.id, aws_subnet.public_subnet2.id]
    security_groups  = [aws_security_group.fargate_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.tg.arn
    container_name   = "api-main-container"
    container_port   = 8080
  }
}

resource "aws_lb" "alb" {
  name_prefix        = "${local.prefix}-"
  internal           = false #tfsec:ignore:aws-elb-alb-not-public
  load_balancer_type = "application"
  security_groups    = [aws_security_group.external_sg.id]
  subnets = [
    aws_subnet.public_subnet.id,
    aws_subnet.public_subnet2.id,
  ]
  drop_invalid_header_fields = true

  # NOTE: disable deletion protection on Demo environment
  enable_deletion_protection = true
}

resource "aws_lb_target_group" "tg" {
  name_prefix = "${local.prefix}-"
  port        = 8080
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"
  health_check {
    enabled = true
    path    = "/health"
  }
}

resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_lb.alb.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS-1-2-2017-01"
  certificate_arn   = data.aws_acm_certificate.main_cert.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.tg.arn
  }
}
