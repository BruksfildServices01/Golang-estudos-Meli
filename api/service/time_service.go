package service

import (
	"api-campeonato/model"
	"errors"
	"sync"
)

type TimeService struct {
	mu     sync.Mutex
	times  map[int]model.Time
	nextID int
}

func NewTimeService() *TimeService {
	return &TimeService{
		times:  make(map[int]model.Time),
		nextID: 1,
	}
}

func (s *TimeService) Create(nome string, cidade string) model.Time {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	t := model.Time{
		ID:     id,
		Nome:   nome,
		Cidade: cidade,
	}

	s.times[id] = t
	return t
}

func (s *TimeService) List() []model.Time {
	s.mu.Lock()
	defer s.mu.Unlock()

	var lista []model.Time
	for _, t := range s.times {
		lista = append(lista, t)
	}
	return lista
}

func (s *TimeService) GetByID(id int) (model.Time, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.times[id]
	if !ok {
		return model.Time{}, errors.New("time não encontrado")
	}
	return t, nil
}

func (s *TimeService) Update(id int, nome string, cidade string) (model.Time, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.times[id]
	if !ok {
		return model.Time{}, errors.New("time não encontrado")
	}

	t := model.Time{
		ID:     id,
		Nome:   nome,
		Cidade: cidade,
	}

	s.times[id] = t
	return t, nil
}

func (s *TimeService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.times[id]; !ok {
		return errors.New("time não encontrado")
	}

	delete(s.times, id)
	return nil
}
