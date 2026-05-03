package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
)

func getCategorias(w http.ResponseWriter, _ *http.Request) {
	categorias, err := repository.BuscarCategorias()
	if err != nil {
		http.Error(w, "Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if categorias == nil {
		categorias = []models.Categoria{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categorias)
}

func createCategoria(w http.ResponseWriter, r *http.Request) {
	var novaCategoria models.Categoria

	err := json.NewDecoder(r.Body).Decode(&novaCategoria)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novaCategoria.NomeCategoria == "" || novaCategoria.Tipo == "" {
		http.Error(w, "Os campos 'nome_categoria' e 'tipo' são obrigatórios", http.StatusBadRequest)
		return
	}

	idGerado, err := repository.CriarCategoria(novaCategoria)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Categoria criada com sucesso",
		"id":       idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}

func HandleCategorias(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCategorias(w, r)
	case http.MethodPost:
		createCategoria(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}