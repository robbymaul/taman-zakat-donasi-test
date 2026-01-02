package services

import (
	"context"
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/config"
	"net/http"
)

type IPostsService interface {
	PublicListPostsService() ([]*models.ListWpPost, error)
	PublicGetPostsService(postName string) (*models.DetailWpPost, error)
}

type PostsService struct {
	*Service
}

func NewPostsService(ctx context.Context, repo *repositories.RepositoryContext, cfg *config.Config) IPostsService {
	return &PostsService{
		Service: NewService(ctx, repo, cfg),
	}
}

func (s *PostsService) PublicListPostsService() ([]*models.ListWpPost, error) {
	var posts []*models.ListWpPost
	var err error

	posts, err = s.repository.PublicListPostsRepository()
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "list posts").WithStatusCode(http.StatusInternalServerError)
	}

	return posts, err
}

func (s *PostsService) PublicGetPostsService(postName string) (*models.DetailWpPost, error) {
	var posts *models.DetailWpPost
	var err error

	posts, err = s.repository.PublicGetPostsRepository(postName)
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "get posts").WithStatusCode(http.StatusInternalServerError)
	}

	return posts, err
}
