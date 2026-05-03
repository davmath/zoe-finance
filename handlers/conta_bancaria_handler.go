package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

func getContasBancarias(w http.ResponseWriter, r *http.Request) {
	var filtro models.FiltroContaBancaria

	if resp := r.URL.Query().Get("id_responsavel"); resp != "" {
		if val, err := strconv.Atoi(resp); err == nil {
			filtro.IDResponsavel = &val
		}
	}

	contas, err := repository.BuscarContasBancarias(filtro)
	if err != nil {
		http.Error(w, "Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if contas == nil {
		contas = []models.ContaBancaria{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contas)
}

func createContaBancaria(w http.ResponseWriter, r *http.Request) {
	var novaConta models.ContaBancaria

	err := json.NewDecoder(r.Body).Decode(&novaConta)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novaConta.Nome == "" {
		http.Error(w, "O campo 'nome' é obrigatório", http.StatusBadRequest)
		return
	}

	idGerado, err := repository.CriarContaBancaria(novaConta)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Conta bancária criada com sucesso",
		"id":       idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}

func HandleContasBancarias(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getContasBancarias(w, r)
	case http.MethodPost:
		createContaBancaria(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}