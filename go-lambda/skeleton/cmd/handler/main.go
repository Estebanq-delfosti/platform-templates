package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/config"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/adapter/lambdahandler"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/adapter/memory"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/service"
)

var h *lambdahandler.Handler

func init() {
	cfg := config.Load()
	repo := memory.NewRepository()
	svc := service.NewItemService(repo)
	h = lambdahandler.NewHandler(svc, cfg)
}

func main() {
	lambda.Start(h.Handle)
}
