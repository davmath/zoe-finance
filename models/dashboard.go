package models

import "time"

type ResumoDashboard struct {
	SaldoAtualContas float64 `json:"saldo_atual_contas"`
	TotalReceitas    float64 `json:"total_receitas"`
	TotalDespesas    float64 `json:"total_despesas"`
	BalancoPeriodo   float64 `json:"balanco_periodo"`
}

type FiltroDashboard struct {
	DataInicio time.Time
	DataFim    time.Time
}

type DespesaPorCategoria struct {
	IDCategoria   int     `json:"id_categoria"`
	NomeCategoria string  `json:"nome_categoria"`
	Total         float64 `json:"total"`
}