package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
)

func BuscarCartoesCredito() ([]models.CartaoCredito, error) {
	query := `
		SELECT 
			cc.id, 
			cc.nome, 
			cc.dia_fechamento, 
			cc.dia_vencimento, 
			cc.limite, 
			cc.id_responsavel, r.nome 
		FROM finance.tb_cartao_credito cc
		LEFT JOIN finance.tb_responsavel_conta r ON cc.id_responsavel = r.id
		ORDER BY cc.id
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartoes []models.CartaoCredito

	for rows.Next() {
		var c models.CartaoCredito
		err := rows.Scan(
			&c.ID, 
			&c.Nome, 
			&c.DiaFechamento, 
			&c.DiaVencimento, 
			&c.Limite, 
			&c.IDResponsavel,
			&c.NomeResponsavel,
		)
		if err != nil {
			return nil, err
		}
		cartoes = append(cartoes, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cartoes, nil
}

func CriarCartaoCredito(c models.CartaoCredito) (int, error) {
	query := `
		INSERT INTO finance.tb_cartao_credito (
			nome, dia_fechamento, dia_vencimento, limite, id_responsavel
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`

	var idGerado int
	err := database.DB.QueryRow(
		query,
		c.Nome,
		c.DiaFechamento,
		c.DiaVencimento,
		c.Limite,
		c.IDResponsavel,
	).Scan(&idGerado)

	if err != nil {
		return 0, err
	}

	return idGerado, nil
}