module github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}

go 1.22

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/google/uuid v1.6.0
)
