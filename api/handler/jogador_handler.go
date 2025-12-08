package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"api-campeonato/service"
)

type JogadorHandler struct {
	Service *service.JogadorService
}

func NewJogadorHandler(s *service.JogadorService) *JogadorHandler {
	return &JogadorHandler{Service: s}
}

func (h *JogadorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/jogadores" {
		switch r.Method {
		case http.MethodPost:
			h.create(w, r)
		case http.MethodGet:
			h.list(w, r)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/jogadores/") {
		id, err := parseIDFromPath(r.URL.Path)
		if err != nil {
			http.Error(w, "id inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.getByID(w, r, id)
		case http.MethodPut:
			h.update(w, r, id)
		case http.MethodDelete:
			h.delete(w, r, id)
		default:
			http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

// @Summary      Cria um novo jogador
// @Description  Cria um jogador vinculado a um time
// @Tags         jogadores
// @Accept       json
// @Produce      json
// @Param        body  body      model.Jogador  true  "Dados do jogador"
// @Success      201   {object}  model.Jogador
// @Failure      400   {string}  string  "json inválido"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /jogadores [post]
func (h *JogadorHandler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nome   string `json:"nome"`
		Idade  int    `json:"idade"`
		TimeID int    `json:"time_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	j := h.Service.Create(input.Nome, input.Idade, input.TimeID)
	writeJSON(w, http.StatusCreated, j)
}

// @Summary      Lista todos os jogadores
// @Description  Retorna todos os jogadores em memória
// @Tags         jogadores
// @Produce      json
// @Success      200  {array}   model.Jogador
// @Failure      405  {string}  string  "método não permitido"
// @Router       /jogadores [get]
func (h *JogadorHandler) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	lista := h.Service.List()
	writeJSON(w, http.StatusOK, lista)
}

// @Summary      Busca jogador por ID
// @Description  Retorna um jogador específico
// @Tags         jogadores
// @Produce      json
// @Param        id   path      int  true  "ID do jogador"
// @Success      200  {object}  model.Jogador
// @Failure      404  {string}  string  "jogador não encontrado"
// @Router       /jogadores/{id} [get]
func (h *JogadorHandler) getByID(w http.ResponseWriter, r *http.Request, id int) {
	j, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "jogador não encontrado", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, j)
}

// @Summary      Atualiza um jogador
// @Description  Atualiza dados de um jogador existente
// @Tags         jogadores
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "ID do jogador"
// @Param        body  body      model.Jogador true  "Dados do jogador"
// @Success      200   {object}  model.Jogador
// @Failure      400   {string}  string  "json inválido"
// @Failure      404   {string}  string  "jogador não encontrado"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /jogadores/{id} [put]
func (h *JogadorHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodPut {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nome   string `json:"nome"`
		Idade  int    `json:"idade"`
		TimeID int    `json:"time_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	j, err := h.Service.Update(id, input.Nome, input.Idade, input.TimeID)
	if err != nil {
		http.Error(w, "jogador não encontrado", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, j)
}

// @Summary      Deleta um jogador
// @Description  Remove um jogador pelo ID
// @Tags         jogadores
// @Param        id  path  int  true  "ID do jogador"
// @Success      204  {string}  string  "No Content"
// @Failure      404  {string}  string  "jogador não encontrado"
// @Failure      405  {string}  string  "método não permitido"
// @Router       /jogadores/{id} [delete]
func (h *JogadorHandler) delete(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "jogador não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
