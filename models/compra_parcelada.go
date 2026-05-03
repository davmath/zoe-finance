package models

import "time"

type CompraParcelada struct {
	ID             int       `json:"id"`
	Descricao      string    `json:"descricao"`
	ValorTotal     float64   `json:"valor_total"`
	QtdParcelas    int       `json:"qtd_parcelas"`
	DataCompra     time.Time `json:"data_compra"`
	IDCartao       int       `json:"id_cartao"`
	IDCategoria    int       `json:"id_categoria"`
	IDResponsavel  int       `json:"id_responsavel"`
	IDSubcategoria *int      `json:"id_subcategoria"`
}

type FiltroCompraParcelada struct {
	Descricao      *string    `json:"descricao"`
	ValorTotalMin  *float64   `json:"valor_total_min"`
	ValorTotalMax  *float64   `json:"valor_total_max"`
	QtdParcelas    *int       `json:"qtd_parcelas"`
	DataInicio     *time.Time `json:"data_inicio"`
	DataFim        *time.Time `json:"data_fim"`
	IDCartao       *int       `json:"id_cartao"`
	IDCategoria    *int       `json:"id_categoria"`
	IDResponsavel  *int       `json:"id_responsavel"`
	IDSubcategoria *int       `json:"id_subcategoria"`
}