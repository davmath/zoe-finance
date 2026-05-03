package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
	"fmt"
)

func BuscarContasBancarias(filtro models.FiltroContaBancaria) ([]models.ContaBancaria, error) {
	query := `
		SELECT 
			cb.id, 
			cb.nome, 
			cb.montante, 
			cb.id_responsavel, r.nome 
		FROM finance.tb_contas_bancarias cb
		LEFT JOIN finance.tb_responsavel_conta r ON cb.id_responsavel = r.id
		WHERE 1=1
	`
	
	var args []interface{}
	paramID := 1

	if filtro.IDResponsavel != nil {
		query += fmt.Sprintf(" AND cb.id_responsavel = $%d", paramID)
		args = append(args, *filtro.IDResponsavel)
		paramID++
	}

	query += " ORDER BY cb.id"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contas []models.ContaBancaria

	for rows.Next() {
		var c models.ContaBancaria
		err := rows.Scan(
			&c.ID,
			&c.Nome,
			&c.Montante,
			&c.IDResponsavel,
			&c.NomeResponsavel,
		)
		if err != nil {
			return nil, err
		}
		contas = append(contas, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return contas, nil
}

func CriarContaBancaria(c models.ContaBancaria) (int, error) {
	query := `
		INSERT INTO finance.tb_contas_bancarias (
			nome, montante, id_responsavel
		) VALUES (
			$1, $2, $3
		) RETURNING id
	`

	var idGerado int
	err := database.DB.QueryRow(
		query,
		c.Nome,
		c.Montante,
		c.IDResponsavel,
	).Scan(&idGerado)

	if err != nil {
		return 0, err
	}

	return idGerado, nil
}