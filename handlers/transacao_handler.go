package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func GetTransacoes(w http.ResponseWriter, r *http.Request){
	
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