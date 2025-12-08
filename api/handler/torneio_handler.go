package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"api-campeonato/service"
)

// TorneioHandler lida com as requisições HTTP de torneios
type TorneioHandler struct {
	Service *service.TorneioService
}

// NewTorneioHandler cria um novo handler de torneio
func NewTorneioHandler(s *service.TorneioService) *TorneioHandler {
	return &TorneioHandler{
		Service: s,
	}
}

// writeJSON é uma ajudinha para responder JSON
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// parseIDFromPath extrai o ID da URL, ex: "/torneios/123"
func parseIDFromPath(path string) (int, error) {
	// Exemplo de path: "/torneios/123"
	parts := strings.Split(strings.Trim(path, "/"), "/")
	// parts[0] = "torneios"
	// parts[1] = "123"
	if len(parts) < 2 {
		return 0, strconv.ErrSyntax
	}
	return strconv.Atoi(parts[1])
}

// ServeHTTP implementa a interface http.Handler
func (h *TorneioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/torneios" {

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

	if strings.HasPrefix(r.URL.Path, "/torneios/") {
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

	// se chegou aqui, rota não encontrada
	http.NotFound(w, r)
}

// ===== Métodos privados do handler =====

// @Summary      Cria um novo torneio
// @Description  Cria um torneio com nome e ano
// @Tags         torneios
// @Accept       json
// @Produce      json
// @Param        body  body      model.Torneio  true  "Dados do torneio"
// @Success      201   {object}  model.Torneio
// @Failure      400   {string}  string  "json inválido"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /torneios [post]
func (h *TorneioHandler) create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Ler JSON
	var input struct {
		Nome string `json:"nome"`
		Ano  int    `json:"ano"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	// Chamar o service
	t := h.Service.Create(input.Nome, input.Ano)

	// Responder
	writeJSON(w, http.StatusCreated, t)
}

// @Summary      Lista todos os torneios
// @Description  Retorna todos os torneios em memória
// @Tags         torneios
// @Produce      json
// @Success      200  {array}   model.Torneio
// @Failure      405  {string}  string  "método não permitido"
// @Router       /torneios [get]
func (h *TorneioHandler) list(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	lista := h.Service.List()
	writeJSON(w, http.StatusOK, lista)
}

// @Summary      Busca torneio por ID
// @Description  Retorna um torneio específico
// @Tags         torneios
// @Produce      json
// @Param        id   path      int  true  "ID do torneio"
// @Success      200  {object}  model.Torneio
// @Failure      404  {string}  string  "torneio não encontrado"
// @Router       /torneios/{id} [get]
func (h *TorneioHandler) getByID(w http.ResponseWriter, r *http.Request, id int) {
	t, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "torneio não encontrado", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// @Summary      Atualiza um torneio
// @Description  Atualiza nome e ano de um torneio existente
// @Tags         torneios
// @Accept       json
// @Produce      json
// @Param        id    path      int           true  "ID do torneio"
// @Param        body  body      model.Torneio true  "Dados do torneio"
// @Success      200   {object}  model.Torneio
// @Failure      400   {string}  string  "json inválido"
// @Failure      404   {string}  string  "torneio não encontrado"
// @Failure      405   {string}  string  "método não permitido"
// @Router       /torneios/{id} [put]
func (h *TorneioHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodPut {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nome string `json:"nome"`
		Ano  int    `json:"ano"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	t, err := h.Service.Update(id, input.Nome, input.Ano)
	if err != nil {
		http.Error(w, "torneio não encontrado", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// @Summary      Deleta um torneio
// @Description  Remove um torneio pelo ID
// @Tags         torneios
// @Param        id  path  int  true  "ID do torneio"
// @Success      204  {string}  string  "No Content"
// @Failure      404  {string}  string  "torneio não encontrado"
// @Failure      405  {string}  string  "método não permitido"
// @Router       /torneios/{id} [delete]
func (h *TorneioHandler) delete(w http.ResponseWriter, r *http.Request, id int) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "torneio não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
