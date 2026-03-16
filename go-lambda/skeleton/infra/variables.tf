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

variable "github_repo" {
  description = "GitHub repository in owner/repo format (used for OIDC trust policy)"
  type        = string
  default     = "${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}"

  validation {
    condition     = can(regex("^[a-z0-9-]+/[a-z0-9-]+$", var.github_repo))
    error_message = "github_repo must be in owner/repo format using lowercase letters, numbers, and hyphens."
  }
}
