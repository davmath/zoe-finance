package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
)

func getResponsavelConta(w http.ResponseWriter, _ *http.Request) {
	responsaveis, err := repository.BuscarResponsaveisConta()
	if err != nil {
		http.Error(w, "Erro interno"+err.Error(), http.StatusInternalServerError)
		return
	}

	if responsaveis == nil {
		responsaveis = []models.ResponsavelConta{}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responsaveis)

}

func createResponsavelConta(w http.ResponseWriter, r *http.Request) {
	var novoResponsavel models.ResponsavelConta

	err := json.NewDecoder(r.Body).Decode(&novoResponsavel)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novoResponsavel.Nome == "" {
		http.Error(w, "O campo 'nome' é obrigatório", http.StatusBadRequest)
		return
	}

	idGerado, err := repository.CriarResponsavel(novoResponsavel)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Responsável criado com sucesso",
		"id":       idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201
	json.NewEncoder(w).Encode(resposta)
}

func HandleResponsavelConta(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getResponsavelConta(w, r)
	case http.MethodPost:
		createResponsavelConta(w, r)
	default:
		http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
	}
}