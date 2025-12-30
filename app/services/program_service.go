package services

import (
	"context"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/app/web"
	"donasitamanzakattest/config"
	"donasitamanzakattest/pkg/util"
	"errors"
	"sync"
)

type ProgramStore struct {
	mu    sync.RWMutex
	items map[string]*models.Program
	order []string
}

var (
	programStoreOnce sync.Once
	programStore     *ProgramStore
)

func GetProgramStore() *ProgramStore {
	programStoreOnce.Do(func() {
		store := &ProgramStore{
			items: map[string]*models.Program{},
			order: []string{},
		}

		seedPrograms := []*models.Program{
			{
				Id:       store.generateUniqueID(12),
				Title:    "AYO KITA BANTU SUDAH",
				Owner:    "Taman Zakat",
				Donasi:   6322253,
				Priority: true,
				Category: "Infaq",
				Image:    "https://donasi.tamanzakat.org/wp-content/uploads/2025/11/bantu-sudan.png",
			},
		}

		for _, program := range seedPrograms {
			store.items[program.Id] = program
			store.order = append(store.order, program.Id)
		}

		programStore = store
	})

	return programStore
}

func (s *ProgramStore) generateUniqueID(length int) string {
	for {
		id := util.GenerateRandomString(length)
		if _, exists := s.items[id]; !exists {
			return id
		}
	}
}

func (s *ProgramStore) Create(program *web.ProgramUpsertRequest) *models.Program {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.generateUniqueID(12)
	cat := &models.Program{
		Id:       id,
		Title:    program.Title,
		Owner:    program.Owner,
		Donasi:   0,
		Priority: program.Priority,
		Category: program.Category,
		Image:    program.Image,
	}
	s.items[id] = cat
	s.order = append(s.order, id)
	return cat
}

func (s *ProgramStore) List() []*models.Program {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]*models.Program, 0, len(s.order))
	for _, id := range s.order {
		if c, ok := s.items[id]; ok {
			out = append(out, &models.Program{
				Id:       c.Id,
				Title:    c.Title,
				Owner:    c.Owner,
				Donasi:   c.Donasi,
				Priority: c.Priority,
				Category: c.Category,
				Image:    c.Image,
			})
		}
	}
	return out
}

func (s *ProgramStore) GetByID(id string) (*models.Program, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}
	return &models.Program{
		Id:       c.Id,
		Title:    c.Title,
		Owner:    c.Owner,
		Donasi:   c.Donasi,
		Priority: c.Priority,
		Category: c.Category,
		Image:    c.Image,
	}, true
}

func (s *ProgramStore) Update(id string, program *models.Program) (*models.Program, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.items[program.Id]
	if !ok {
		return nil, false
	}
	c.Title = program.Title
	c.Owner = program.Owner
	c.Donasi = program.Donasi
	c.Priority = program.Priority
	c.Category = program.Category
	c.Image = program.Image
	return &models.Program{
		Id:       c.Id,
		Title:    c.Title,
		Owner:    c.Owner,
		Donasi:   c.Donasi,
		Priority: c.Priority,
		Category: c.Category,
		Image:    c.Image,
	}, true
}

func (s *ProgramStore) Delete(id string) (*models.Program, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}

	deleted := &models.Program{
		Id:       c.Id,
		Title:    c.Title,
		Owner:    c.Owner,
		Donasi:   c.Donasi,
		Priority: c.Priority,
		Category: c.Category,
		Image:    c.Image,
	}
	delete(s.items, id)

	for i := 0; i < len(s.order); i++ {
		if s.order[i] == id {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}

	return deleted, true
}

type ProgramService struct {
	service *Service
	store   *ProgramStore
}

func NewProgramService(ctx context.Context, r *repositories.RepositoryContext, cfg *config.Config) *ProgramService {
	return &ProgramService{
		service: NewService(ctx, r, cfg),
		store:   GetProgramStore(),
	}
}

func (s *ProgramService) Create(payload *web.ProgramUpsertRequest) (*models.Program, error) {

	return s.store.Create(payload), nil
}

func (s *ProgramService) List() ([]*models.Program, error) {
	return s.store.List(), nil
}

func (s *ProgramService) GetByID(id string) (*models.Program, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.GetByID(id)
	if !ok {
		return nil, errors.New("program not found")
	}
	return c, nil
}

func (s *ProgramService) Update(id string, payload *models.Program) (*models.Program, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.Update(id, payload)
	if !ok {
		return nil, errors.New("program not found")
	}
	return c, nil
}

func (s *ProgramService) Delete(id string) (*models.Program, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.Delete(id)
	if !ok {
		return nil, errors.New("program not found")
	}
	return c, nil
}
