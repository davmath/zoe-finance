package models

type ContaBancaria struct {
	ID              int     `json:"id"`
	Nome            string  `json:"nome"`
	Montante        float64 `json:"montante"`
	
	IDResponsavel   int     `json:"id_responsavel"`
	NomeResponsavel *string `json:"responsavel_nome,omitempty"`
}

type FiltroContaBancaria struct {
	IDResponsavel *int `json:"id_responsavel"`
}