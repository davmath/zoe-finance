package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
)

func BuscarSubcategorias() ([]models.Subcategoria, error) {
	query := `
		SELECT 
			s.id, 
			s.id_categoria, c.nome_categoria,
			s.nome_subcategoria 
		FROM finance.tb_subcategorias s
		LEFT JOIN finance.tb_categorias c ON s.id_categoria = c.id
		ORDER BY s.id
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategorias []models.Subcategoria

	for rows.Next() {
		var s models.Subcategoria
		
		err := rows.Scan(
			&s.ID, 
			&s.IDCategoria, 
			&s.NomeCategoria,
			&s.NomeSubcategoria,
		)
		
		if err != nil {
			return nil, err
		}
		subcategorias = append(subcategorias, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subcategorias, nil
}

func CriarSubcategoria(s models.Subcategoria) (int, error) {
	query := "INSERT INTO finance.tb_subcategorias (id_categoria, nome_subcategoria) VALUES ($1, $2) RETURNING id"

	var idGerado int
	err := database.DB.QueryRow(query, s.IDCategoria, s.NomeSubcategoria).Scan(&idGerado)
	if err != nil {
		return 0, err
	}

	return idGerado, nil
}