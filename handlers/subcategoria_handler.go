package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
)

func getSubcategorias(w http.ResponseWriter, _ *http.Request) {
	subcategorias, err := repository.BuscarSubcategorias()
	if err != nil {
		http.Error(w, "Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if subcategorias == nil {
		subcategorias = []models.Subcategoria{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subcategorias)
}

func createSubcategoria(w http.ResponseWriter, r *http.Request) {
	var novaSubcategoria models.Subcategoria

	err := json.NewDecoder(r.Body).Decode(&novaSubcategoria)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novaSubcategoria.NomeSubcategoria == "" || novaSubcategoria.IDCategoria == 0 {
		http.Error(w, "Os campos 'nome_subcategoria' e 'id_categoria' são obrigatórios", http.StatusBadRequest)
		return
	}

	idGerado, err := repository.CriarSubcategoria(novaSubcategoria)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Subcategoria criada com sucesso",
		"id":       idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}

func HandleSubcategorias(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getSubcategorias(w, r)
	case http.MethodPost:
		createSubcategoria(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}