package services

import (
	"context"
	"donasitamanzakattest/app/helpers"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/config"
	"donasitamanzakattest/pkg/util"
	"errors"
	"net/http"
	"sync"
)

type CategoryStore struct {
	mu    sync.RWMutex
	items map[string]*models.Category
	order []string
}

var (
	categoryStoreOnce sync.Once
	categoryStore     *CategoryStore
)

func GetCategoryStore() *CategoryStore {
	categoryStoreOnce.Do(func() {
		store := &CategoryStore{
			items: map[string]*models.Category{},
			order: []string{},
		}

		seedNames := []string{"Palestina", "Infaq", "Zakat", "Wakaf", "Ramadan", "Qurban"}
		for _, name := range seedNames {
			id := store.generateUniqueID(12)
			cat := &models.Category{Id: id, Name: name}
			store.items[id] = cat
			store.order = append(store.order, id)
		}

		categoryStore = store
	})

	return categoryStore
}

func (s *CategoryStore) generateUniqueID(length int) string {
	for {
		id := util.GenerateRandomString(length)
		if _, exists := s.items[id]; !exists {
			return id
		}
	}
}

func (s *CategoryStore) Create(name string) *models.Category {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.generateUniqueID(12)
	cat := &models.Category{Id: id, Name: name}
	s.items[id] = cat
	s.order = append(s.order, id)
	return cat
}

func (s *CategoryStore) List() []*models.Category {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]*models.Category, 0, len(s.order))
	for _, id := range s.order {
		if c, ok := s.items[id]; ok {
			out = append(out, &models.Category{Id: c.Id, Name: c.Name})
		}
	}
	return out
}

func (s *CategoryStore) GetByID(id string) (*models.Category, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}
	return &models.Category{Id: c.Id, Name: c.Name}, true
}

func (s *CategoryStore) Update(id string, name string) (*models.Category, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}
	c.Name = name
	return &models.Category{Id: c.Id, Name: c.Name}, true
}

func (s *CategoryStore) Delete(id string) (*models.Category, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}

	deleted := &models.Category{Id: c.Id, Name: c.Name}
	delete(s.items, id)

	for i := 0; i < len(s.order); i++ {
		if s.order[i] == id {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}

	return deleted, true
}

type CategoryService struct {
	service *Service
	store   *CategoryStore
}

func NewCategoryService(ctx context.Context, r *repositories.RepositoryContext, cfg *config.Config) *CategoryService {
	return &CategoryService{
		service: NewService(ctx, r, cfg),
		store:   GetCategoryStore(),
	}
}

func (s *CategoryService) Create(name string) (*models.Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	return s.store.Create(name), nil
}

func (s *CategoryService) List() ([]*models.Category, error) {
	return s.store.List(), nil
}

func (s *CategoryService) GetByID(id string) (*models.Category, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.GetByID(id)
	if !ok {
		return nil, errors.New("category not found")
	}
	return c, nil
}

func (s *CategoryService) Update(id string, name string) (*models.Category, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	c, ok := s.store.Update(id, name)
	if !ok {
		return nil, errors.New("category not found")
	}
	return c, nil
}

func (s *CategoryService) Delete(id string) (*models.Category, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.Delete(id)
	if !ok {
		return nil, errors.New("category not found")
	}
	return c, nil
}

func (s *CategoryService) PublicListCategoriesService() ([]*models.WpDjaCategory, error) {
	var err error
	var categories []*models.WpDjaCategory

	categories, err = s.service.repository.PublicListCategoriesRepositories()
	if err != nil {
		return nil, helpers.NewErrorTrace(err, "list categories").WithStatusCode(http.StatusInternalServerError)
	}

	return categories, err
}
