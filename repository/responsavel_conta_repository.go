package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
)

func BuscarResponsaveisConta() ([]models.ResponsavelConta, error) {
	query := "SELECT id, nome FROM finance.tb_responsavel_conta ORDER BY ID"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var responsaveis []models.ResponsavelConta

	for rows.Next() {
		var r models.ResponsavelConta
		err := rows.Scan(&r.ID, &r.Nome)
		if err != nil {
			return nil, err
		}
		responsaveis = append(responsaveis, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return responsaveis, nil
	
}

func CriarResponsavel(r models.ResponsavelConta) (int, error) {
	query := "INSERT INTO finance.tb_responsavel_conta (nome) VALUES ($1) RETURNING id"
	
	var idGerado int
	err := database.DB.QueryRow(query, r.Nome).Scan(&idGerado)
	if err != nil {
		return 0, err
	}

	return idGerado, nil
}
