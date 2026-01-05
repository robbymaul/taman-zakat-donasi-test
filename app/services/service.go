package services

import (
	"context"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/config"
	pkgjwt "donasitamanzakattest/pkg/jwt"
	"donasitamanzakattest/pkg/middleware"
	"fmt"
)

type Service struct {
	ctx        context.Context
	config     *config.Config
	repository *repositories.RepositoryContext
}

func NewService(ctx context.Context, r *repositories.RepositoryContext, cfg *config.Config) *Service {

	return &Service{ctx: ctx, config: cfg, repository: r}
}

func (s *Service) GetSessionService() (*models.WpUserSession, error) {
	var wpUserSession *models.WpUserSession
	var err error

	session, ok := s.ctx.Value(middleware.Session).(*pkgjwt.JwtResponse)
	if !ok {
		return nil, fmt.Errorf("no session")
	}

	wpUserSession, err = s.repository.GetSessionRepository(session)
	if err != nil {
		return nil, err
	}

	return wpUserSession, err
}
