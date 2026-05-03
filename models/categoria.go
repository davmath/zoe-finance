package models

// Categoria representa o registro na tabela finance.tb_categorias
type Categoria struct {
	ID            int    `json:"id"`
	NomeCategoria string `json:"nome_categoria"`
	Tipo          string `json:"tipo"`
}