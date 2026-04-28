package models

import "time"

type Transacao struct {
	ID                int        `json:"id"`
	Descricao         string     `json:"descricao"`
	Valor             float64    `json:"valor"`
	DataTransacao     time.Time  `json:"data_transacao"`
	IDCategoria       int        `json:"id_categoria"`
	IDSubcategoria    *int       `json:"id_subcategoria"`
	IDResponsavel     int        `json:"id_responsavel"`
	IDContaBancaria   *int       `json:"id_conta_bancaria"`
	IDCartaoCredito   *int       `json:"id_cartao_credito"`
	IDContaDestino    *int       `json:"id_conta_destino"`
	IDCompraParcelada *int       `json:"id_compra_parcelada"`
	Efetivada         bool       `json:"efetivada"`
}

type FiltroTransacao struct {
	Descricao         *string    `json:"descricao"`
	ValorMin          *float64   `json:"valor_min"`
	ValorMax          *float64   `json:"valor_max"`
	DataInicio        *time.Time `json:"data_inicio"`
	DataFim           *time.Time `json:"data_fim"`
	IDCategoria       *int       `json:"id_categoria"`
	IDSubcategoria    *int       `json:"id_subcategoria"`
	IDResponsavel     *int       `json:"id_responsavel"`
	IDContaBancaria   *int       `json:"id_conta_bancaria"`
	IDCartaoCredito   *int       `json:"id_cartao_credito"`
	IDContaDestino    *int       `json:"id_conta_destino"`
	IDCompraParcelada *int       `json:"id_compra_parcelada"`
	Efetivada         *bool      `json:"efetivada"`
}