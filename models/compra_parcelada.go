package models

import "time"

type CompraParcelada struct {
	ID               int       `json:"id"`
	Descricao        string    `json:"descricao"`
	ValorTotal       float64   `json:"valor_total"`
	QtdParcelas      int       `json:"qtd_parcelas"`
	DataCompra       time.Time `json:"data_compra"`
	
	IDCartao         int       `json:"id_cartao"`
	NomeCartao       *string   `json:"cartao_nome,omitempty"`
	
	IDCategoria      int       `json:"id_categoria"`
	NomeCategoria    *string   `json:"categoria_nome,omitempty"`
	
	IDResponsavel    int       `json:"id_responsavel"`
	NomeResponsavel  *string   `json:"responsavel_nome,omitempty"`
	
	IDSubcategoria   *int      `json:"id_subcategoria"`
	NomeSubcategoria *string   `json:"subcategoria_nome,omitempty"`
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