module github.com/${{ values.repoOwner }}/${{ values.name }}

go 1.22

require (
	github.com/go-chi/chi/v5 v5.2.5
	github.com/go-chi/cors v1.2.2
	github.com/google/uuid v1.6.0
	go.uber.org/zap v1.27.0
)

require go.uber.org/multierr v1.10.0 // indirect
