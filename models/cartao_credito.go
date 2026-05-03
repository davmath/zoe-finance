package models

type CartaoCredito struct {
	ID            int     `json:"id"`
	Nome          string  `json:"nome"`
	DiaFechamento int     `json:"dia_fechamento"`
	DiaVencimento int     `json:"dia_vencimento"`
	Limite        float64 `json:"limite"`
	IDResponsavel int     `json:"id_responsavel"`
}