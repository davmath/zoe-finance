package models

type Subcategoria struct {
	ID               int    `json:"id"`
	
	IDCategoria      int     `json:"id_categoria"`
	NomeCategoria    *string `json:"categoria_nome,omitempty"`

	NomeSubcategoria string `json:"nome_subcategoria"`
}