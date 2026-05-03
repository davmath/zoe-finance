package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
)

func getCartoesCredito(w http.ResponseWriter, _ *http.Request) {
	cartoes, err := repository.BuscarCartoesCredito()
	if err != nil {
		http.Error(w, "Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if cartoes == nil {
		cartoes = []models.CartaoCredito{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cartoes)
}

func createCartaoCredito(w http.ResponseWriter, r *http.Request) {
	var novoCartao models.CartaoCredito

	err := json.NewDecoder(r.Body).Decode(&novoCartao)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novoCartao.Nome == "" {
		http.Error(w, "O campo 'nome' é obrigatório", http.StatusBadRequest)
		return
	}

	idGerado, err := repository.CriarCartaoCredito(novoCartao)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Cartão de crédito criado com sucesso",
		"id":       idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}

func HandleCartoesCredito(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCartoesCredito(w, r)
	case http.MethodPost:
		createCartaoCredito(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}
