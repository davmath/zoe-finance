package models

import "time"

// Transacao representa o registro na tabela finance.tb_transacoes
type Transacao struct {
	ID                  int       `json:"id"`
	Descricao           string    `json:"descricao"`
	Valor               float64   `json:"valor"`
	DataTransacao       time.Time `json:"data_transacao"`
	Efetivada           bool      `json:"efetivada"`

	IDCategoria         *int      `json:"id_categoria"`
	NomeCategoria       *string   `json:"categoria_nome,omitempty"`

	IDSubcategoria      *int      `json:"id_subcategoria"`
	NomeSubcategoria    *string   `json:"subcategoria_nome,omitempty"`

	IDResponsavel       *int      `json:"id_responsavel"`
	NomeResponsavel     *string   `json:"responsavel_nome,omitempty"`

	IDContaBancaria     *int      `json:"id_conta_bancaria"`
	NomeContaBancaria   *string   `json:"conta_bancaria_nome,omitempty"`

	IDCartaoCredito     *int      `json:"id_cartao_credito"`
	NomeCartaoCredito   *string   `json:"cartao_credito_nome,omitempty"`

	IDContaDestino      *int      `json:"id_conta_destino"`
	NomeContaDestino    *string   `json:"conta_destino_nome,omitempty"`

	IDCompraParcelada   *int      `json:"id_compra_parcelada"`
	DescCompraParcelada *string   `json:"compra_parcelada_descricao,omitempty"`
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