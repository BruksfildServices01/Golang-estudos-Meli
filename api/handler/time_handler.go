package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"api-campeonato/service"
)

type TimeHandler struct {
	Service *service.TimeService
}

func NewTimeHandler(s *service.TimeService) *TimeHandler {
	return &TimeHandler{Service: s}
}

func (h *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/times" {
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

	if strings.HasPrefix(r.URL.Path, "/times/") {
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

// @Summary      Cria um novo time
// @Description  Cria um time com nome e cidade
// @Tags         times
// @Accept       json
// @Produce      json
// @Param        body  body      model.Time  true  "Dados do time"
// @Success      201   {object}  model.Time
// @Failure      400   {string}  string  "json inválido"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /times [post]
func (h *TimeHandler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nome   string `json:"nome"`
		Cidade string `json:"cidade"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	t := h.Service.Create(input.Nome, input.Cidade)
	writeJSON(w, http.StatusCreated, t)
}

// @Summary      Lista todos os times
// @Description  Retorna todos os times em memória
// @Tags         times
// @Produce      json
// @Success      200  {array}   model.Time
// @Failure      405  {string}  string  "método não permitido"
// @Router       /times [get]
func (h *TimeHandler) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	lista := h.Service.List()
	writeJSON(w, http.StatusOK, lista)
}

// @Summary      Busca time por ID
// @Description  Retorna um time específico
// @Tags         times
// @Produce      json
// @Param        id   path      int  true  "ID do time"
// @Success      200  {object}  model.Time
// @Failure      404  {string}  string  "time não encontrado"
// @Router       /times/{id} [get]
func (h *TimeHandler) getByID(w http.ResponseWriter, r *http.Request, id int) {
	t, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "time não encontrado", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// @Summary      Atualiza um time
// @Description  Atualiza nome e cidade de um time existente
// @Tags         times
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "ID do time"
// @Param        body  body      model.Time true  "Dados do time"
// @Success      200   {object}  model.Time
// @Failure      400   {string}  string  "json inválido"
// @Failure      404   {string}  string  "time não encontrado"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /times/{id} [put]
func (h *TimeHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodPut {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nome   string `json:"nome"`
		Cidade string `json:"cidade"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	t, err := h.Service.Update(id, input.Nome, input.Cidade)
	if err != nil {
		http.Error(w, "time não encontrado", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// @Summary      Deleta um time
// @Description  Remove um time pelo ID
// @Tags         times
// @Param        id  path  int  true  "ID do time"
// @Success      204  {string}  string  "No Content"
// @Failure      404  {string}  string  "time não encontrado"
// @Failure      405  {string}  string  "método não permitido"
// @Router       /times/{id} [delete]
func (h *TimeHandler) delete(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "time não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
