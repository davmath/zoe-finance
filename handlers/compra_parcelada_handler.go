package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func getComprasParceladas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var filtro models.FiltroCompraParcelada

	if desc := query.Get("descricao"); desc != "" {
		filtro.Descricao = &desc
	}

	if vMin := query.Get("valor_total_min"); vMin != "" {
		if val, err := strconv.ParseFloat(vMin, 64); err == nil {
			filtro.ValorTotalMin = &val
		}
	}
	if vMax := query.Get("valor_total_max"); vMax != "" {
		if val, err := strconv.ParseFloat(vMax, 64); err == nil {
			filtro.ValorTotalMax = &val
		}
	}

	if qtd := query.Get("qtd_parcelas"); qtd != "" {
		if val, err := strconv.Atoi(qtd); err == nil {
			filtro.QtdParcelas = &val
		}
	}
	if cartao := query.Get("id_cartao"); cartao != "" {
		if val, err := strconv.Atoi(cartao); err == nil {
			filtro.IDCartao = &val
		}
	}
	if cat := query.Get("id_categoria"); cat != "" {
		if val, err := strconv.Atoi(cat); err == nil {
			filtro.IDCategoria = &val
		}
	}
	if resp := query.Get("id_responsavel"); resp != "" {
		if val, err := strconv.Atoi(resp); err == nil {
			filtro.IDResponsavel = &val
		}
	}
	if subcat := query.Get("id_subcategoria"); subcat != "" {
		if val, err := strconv.Atoi(subcat); err == nil {
			filtro.IDSubcategoria = &val
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
			val = val.Add((23 * time.Hour) + (59 * time.Minute) + (59 * time.Second))
			filtro.DataFim = &val
		}
	}

	compras, err := repository.BuscarComprasParceladas(filtro)
	if err != nil {
		http.Error(w, "Erro interno: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if compras == nil {
		compras = []models.CompraParcelada{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(compras)
}

func createCompraParcelada(w http.ResponseWriter, r *http.Request) {
	var novaCompra models.CompraParcelada

	err := json.NewDecoder(r.Body).Decode(&novaCompra)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if novaCompra.DataCompra.IsZero() {
		novaCompra.DataCompra = time.Now()
	}

	idGerado, err := repository.CriarCompraParcelada(novaCompra)
	if err != nil {
		http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resposta := map[string]interface{}{
		"mensagem": "Compra parcelada criada com sucesso",
		"id":       idGerado,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposta)
}

func HandleComprasParceladas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getComprasParceladas(w, r)
	case http.MethodPost:
		createCompraParcelada(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}