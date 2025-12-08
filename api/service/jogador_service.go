package service

import (
	"errors"
	"sync"

	"api-campeonato/model"
)

type JogadorService struct {
	mu        sync.Mutex
	jogadores map[int]model.Jogador
	nextID    int
}

func NewJogadorService() *JogadorService {
	return &JogadorService{
		jogadores: make(map[int]model.Jogador),
		nextID:    1,
	}
}

func (s *JogadorService) Create(nome string, idade int, timeID int) model.Jogador {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID
	s.nextID++

	j := model.Jogador{
		ID:     id,
		Nome:   nome,
		Idade:  idade,
		TimeID: timeID,
	}

	s.jogadores[id] = j
	return j
}

func (s *JogadorService) List() []model.Jogador {
	s.mu.Lock()
	defer s.mu.Unlock()

	var lista []model.Jogador
	for _, j := range s.jogadores {
		lista = append(lista, j)
	}
	return lista
}

func (s *JogadorService) GetByID(id int) (model.Jogador, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	j, ok := s.jogadores[id]
	if !ok {
		return model.Jogador{}, errors.New("jogador não encontrado")
	}
	return j, nil
}

func (s *JogadorService) Update(id int, nome string, idade int, timeID int) (model.Jogador, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.jogadores[id]
	if !ok {
		return model.Jogador{}, errors.New("jogador não encontrado")
	}

	j := model.Jogador{
		ID:     id,
		Nome:   nome,
		Idade:  idade,
		TimeID: timeID,
	}

	s.jogadores[id] = j
	return j, nil
}

func (s *JogadorService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.jogadores[id]; !ok {
		return errors.New("jogador não encontrado")
	}

	delete(s.jogadores, id)
	return nil
}
