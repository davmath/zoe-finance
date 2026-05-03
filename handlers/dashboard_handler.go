package handlers

import (
	"davmath/zoe-finance/models"
	"davmath/zoe-finance/repository"
	"encoding/json"
	"net/http"
	"time"
)

func getResumoDashboard(w http.ResponseWriter, r *http.Request) {
	var filtro models.FiltroDashboard
	layoutData := "2006-01-02"

	strDataInicio := r.URL.Query().Get("data_inicio")
	if strDataInicio != "" {
		if d, err := time.Parse(layoutData, strDataInicio); err == nil {
			filtro.DataInicio = d
		}
	}

	strDataFim := r.URL.Query().Get("data_fim")
	if strDataFim != "" {
		if d, err := time.Parse(layoutData, strDataFim); err == nil {
			filtro.DataFim = d
		}
	}

	if filtro.DataInicio.IsZero() || filtro.DataFim.IsZero() {
		agora := time.Now()
		// Primeiro dia do mês atual
		filtro.DataInicio = time.Date(agora.Year(), agora.Month(), 1, 0, 0, 0, 0, agora.Location())
		// Último dia do mês atual (adiciona 1 mês, volta 1 segundo)
		filtro.DataFim = filtro.DataInicio.AddDate(0, 1, 0).Add(-time.Second)
	}

	resumo, err := repository.ObterResumo(filtro)
	if err != nil {
		http.Error(w, "Erro ao calcular resumo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resumo)
}

func getDespesasPorCategoria(w http.ResponseWriter, r *http.Request) {
	var filtro models.FiltroDashboard
	layoutData := "2006-01-02"

	if strData := r.URL.Query().Get("data_inicio"); strData != "" {
		if d, err := time.Parse(layoutData, strData); err == nil {
			filtro.DataInicio = d
		}
	}

	if strData := r.URL.Query().Get("data_fim"); strData != "" {
		if d, err := time.Parse(layoutData, strData); err == nil {
			filtro.DataFim = d
		}
	}

	if filtro.DataInicio.IsZero() || filtro.DataFim.IsZero() {
		agora := time.Now()
		filtro.DataInicio = time.Date(agora.Year(), agora.Month(), 1, 0, 0, 0, 0, agora.Location())
		filtro.DataFim = filtro.DataInicio.AddDate(0, 1, 0).Add(-time.Second)
	}

	despesas, err := repository.ObterDespesasPorCategoria(filtro)
	if err != nil {
		http.Error(w, "Erro ao buscar despesas por categoria: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(despesas)
}

func HandleDashboardDespesasCategoria(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDespesasPorCategoria(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func HandleDashboardResumo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getResumoDashboard(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}