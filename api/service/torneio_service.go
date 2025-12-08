package service

import (
	"errors"
	"math/rand"
	"sync"

	"api-campeonato/model"
)

// TorneioService é responsável por gerenciar torneios
type TorneioService struct {
	mu       sync.Mutex
	torneios map[int]model.Torneio
}

// NewTorneioService cria um novo serviço de torneios
func NewTorneioService() *TorneioService {
	return &TorneioService{
		torneios: make(map[int]model.Torneio),
	}
}

// Create cria um novo torneio
func (s *TorneioService) Create(nome string, ano int) model.Torneio {
	s.mu.Lock()
	defer s.mu.Unlock()

	// gera um ID aleatório simples
	id := rand.Intn(1000000)

	t := model.Torneio{
		ID:   id,
		Nome: nome,
		Ano:  ano,
	}

	s.torneios[id] = t
	return t
}

// List retorna todos os torneios
func (s *TorneioService) List() []model.Torneio {
	s.mu.Lock()
	defer s.mu.Unlock()

	var lista []model.Torneio
	for _, t := range s.torneios {
		lista = append(lista, t)
	}
	return lista
}

// GetByID retorna um torneio pelo ID
func (s *TorneioService) GetByID(id int) (model.Torneio, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.torneios[id]
	if !ok {
		return model.Torneio{}, errors.New("torneio não encontrado")
	}
	return t, nil
}

// Update atualiza um torneio existente
func (s *TorneioService) Update(id int, nome string, ano int) (model.Torneio, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.torneios[id]
	if !ok {
		return model.Torneio{}, errors.New("torneio não encontrado")
	}

	t := model.Torneio{
		ID:   id,
		Nome: nome,
		Ano:  ano,
	}

	s.torneios[id] = t
	return t, nil
}

// Delete remove um torneio pelo ID
func (s *TorneioService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.torneios[id]; !ok {
		return errors.New("torneio não encontrado")
	}

	delete(s.torneios, id)
	return nil
}
