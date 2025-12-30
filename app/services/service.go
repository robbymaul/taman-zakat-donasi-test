package services

import (
	"context"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/config"
)

type Service struct {
	ctx        context.Context
	config     *config.Config
	repository *repositories.RepositoryContext
}

func NewService(ctx context.Context, r *repositories.RepositoryContext, cfg *config.Config) *Service {

	return &Service{ctx: ctx, config: cfg, repository: r}
}
