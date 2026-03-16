provider "aws" {
  region = var.aws_region
}

locals {
  lambda_arch = var.architecture == "amd64" ? "x86_64" : var.architecture
}

resource "aws_iam_role" "lambda" {
  name = "${var.name}-lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action    = "sts:AssumeRole"
      Effect    = "Allow"
      Principal = { Service = "lambda.amazonaws.com" }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "this" {
  function_name    = var.name
  handler          = "bootstrap"
  runtime          = "provided.al2023"
  architectures    = [local.lambda_arch]
  filename         = "${path.module}/../dist/function.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/function.zip")
  role             = aws_iam_role.lambda.arn
  timeout          = 30
  memory_size      = 256

  environment {
    variables = {
      LOG_LEVEL   = "INFO"
      ENVIRONMENT = var.environment
    }
  }
}

resource "aws_apigatewayv2_api" "this" {
  name          = var.name
  protocol_type = "HTTP"
  target        = aws_lambda_function.this.invoke_arn
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.this.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.this.execution_arn}/*/*"
}
