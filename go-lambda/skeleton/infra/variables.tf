variable "name" {
  description = "Name of the Lambda function"
  type        = string
  default     = "${{ values.name }}"
}

variable "aws_region" {
  description = "AWS region to deploy to"
  type        = string
  default     = "${{ values.awsRegion }}"
}

variable "environment" {
  description = "Deployment environment (e.g. production, staging)"
  type        = string
  default     = "production"
}

variable "architecture" {
  description = "Lambda architecture: arm64 or amd64 (mapped to x86_64 for AWS)"
  type        = string
  default     = "${{ values.architecture }}"
}
