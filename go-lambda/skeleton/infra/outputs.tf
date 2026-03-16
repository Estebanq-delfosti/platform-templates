output "function_arn" {
  description = "ARN of the Lambda function"
  value       = aws_lambda_function.this.arn
}

output "api_endpoint" {
  description = "HTTP API endpoint URL"
  value       = aws_apigatewayv2_api.this.api_endpoint
}
