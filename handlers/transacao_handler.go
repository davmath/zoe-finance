package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func getTransacoes(w http.ResponseWriter, r *http.Request){
	
	if r.Method != http.MethodGet{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	var filtro models.FiltroTransacao

	if desc := query.Get("descricao"); desc != "" {
		filtro.Descricao = &desc
	}

	if vMin := query.Get("valor_minimo"); vMin != ""{
		if val, err := strconv.ParseFloat(vMin, 64); err == nil {
			filtro.ValorMin = &val
		}
	}

	if vMax := query.Get("valor_maximo"); vMax != "" {
		if val, err := strconv.ParseFloat(vMax, 64); err == nil {
			filtro.ValorMax = &val
		}
	}

	if cat := query.Get("id_categoria"); cat != "" {
		if val, err := strconv.Atoi(cat); err == nil {
			filtro.IDCategoria = &val
		}		
	}

	if subcat := query.Get("id_subcategoria"); subcat != "" {
		if val, err := strconv.Atoi(subcat); err == nil {
			filtro.IDSubcategoria = &val
		}
	}

	if resp := query.Get("id_responsavel"); resp != "" {
		if val, err := strconv.Atoi(resp); err == nil {
			filtro.IDResponsavel = &val
		}
	}
	if conta := query.Get("id_conta_bancaria"); conta != "" {
		if val, err := strconv.Atoi(conta); err == nil {
			filtro.IDContaBancaria = &val
		}
	}
	if cartao := query.Get("id_cartao_credito"); cartao != "" {
		if val, err := strconv.Atoi(cartao); err == nil {
			filtro.IDCartaoCredito = &val
		}
	}
	if dest := query.Get("id_conta_destino"); dest != "" {
		if val, err := strconv.Atoi(dest); err == nil {
			filtro.IDContaDestino = &val
		}
	}
	if parc := query.Get("id_compra_parcelada"); parc != "" {
		if val, err := strconv.Atoi(parc); err == nil {
			filtro.IDCompraParcelada = &val
		}
	}

	if efet := query.Get("efetivada"); efet != "" {
		if val, err := strconv.ParseBool(efet); err == nil {
			filtro.Efetivada = &val
		}
	}

	layoutData := "2006-01-02"

	if dataIni := query.Get("data_inicio"); dataIni != "" {
		if val, err := time.Parse(layoutData, dataIni); err == nil {
			filtro.DataInicio = &val
		}
	}

	if dataFim := query.Get("data_fim"); dataFim != "" {
		if val, err := time.Parse(layoutData, dataFim); err == nil {
			filtro.DataFim = &val
		}
	}

	transacoes, err := repository.BuscarTransacoes(filtro)
	if err != nil {
		http.Error(w, "Erro interno: " +err.Error(), http.StatusInternalServerError)
		return
	}

	if transacoes == nil {
		transacoes = []models.Transacao{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transacoes)

}

// createTransacao godoc
// @Summary      Cadastrar nova transação
// @Description  Insere uma nova transação financeira no banco de dados
// @Tags         transacoes
// @Accept       json
// @Produce      json
// @Param        transacao  body      models.Transacao  true  "Dados da Transação"
// @Success      201        {object}  map[string]interface{}
// @Failure      400        {string}  string "JSON inválido"
// @Router       /transacoes [post]
func createTransacao(w http.ResponseWriter, r *http.Request) {
	var novaTransacao models.Transacao

	err := json.NewDecoder(r.Body).Decode(&novaTransacao)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novaTransacao.DataTransacao.IsZero() {
		novaTransacao.DataTransacao = time.Now()
	}

	idGerado, err := repository.CriarTransacao(novaTransacao) 
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Transação criada com sucesso",
		"id": idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}

func patchTransacao(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID é obrigatório.", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idStr)

	var dados map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	delete(dados, "id")

	err = repository.AtualizarTransacao(id, dados)
	if err != nil {
		http.Error(w, "Erro ao atualizar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteTransacao(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Formato de ID inválido", http.StatusBadRequest)
		return
	}

	err = repository.DeletarTransacao(id)
	if err != nil {
		if err.Error() == "registro_nao_encontrado" {
			http.Error(w, "Transação não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao deletar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HandleTransacoes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTransacoes(w, r)
	case http.MethodPost:
		createTransacao(w, r)
	case http.MethodPatch:
		patchTransacao(w, r)
	case http.MethodDelete:
		patchTransacao(w, r)
	default:
		http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
	}
}