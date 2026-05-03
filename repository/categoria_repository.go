package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
)

func BuscarCategorias() ([]models.Categoria, error) {
	query := "SELECT id, nome_categoria, tipo FROM finance.tb_categorias ORDER BY id"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []models.Categoria

	for rows.Next() {
		var c models.Categoria
		err := rows.Scan(&c.ID, &c.NomeCategoria, &c.Tipo)
		if err != nil {
			return nil, err
		}
		categorias = append(categorias, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categorias, nil
}

func CriarCategoria(c models.Categoria) (int, error) {
	query := "INSERT INTO finance.tb_categorias (nome_categoria, tipo) VALUES ($1, $2) RETURNING id"

	var idGerado int
	err := database.DB.QueryRow(query, c.NomeCategoria, c.Tipo).Scan(&idGerado)
	if err != nil {
		return 0, err
	}

	return idGerado, nil
}