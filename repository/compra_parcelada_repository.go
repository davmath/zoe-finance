package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
	"fmt"
)

// BuscarComprasParceladas retorna os registros baseados nos filtros
func BuscarComprasParceladas(filtro models.FiltroCompraParcelada) ([]models.CompraParcelada, error) {
	query := `
		SELECT 
			id, descricao, valor_total, qtd_parcelas, data_compra, 
			id_cartao, id_categoria, id_responsavel, id_subcategoria 
		FROM finance.tb_compras_parceladas 
		WHERE 1=1
	`
	
	var args []interface{}
	paramID := 1

	if filtro.Descricao != nil {
		query += fmt.Sprintf(" AND descricao ILIKE $%d", paramID)
		args = append(args, "%"+*filtro.Descricao+"%")
		paramID++
	}

	if filtro.ValorTotalMin != nil {
		query += fmt.Sprintf(" AND valor_total >= $%d", paramID)
		args = append(args, *filtro.ValorTotalMin)
		paramID++
	}

	if filtro.ValorTotalMax != nil {
		query += fmt.Sprintf(" AND valor_total <= $%d", paramID)
		args = append(args, *filtro.ValorTotalMax)
		paramID++
	}

	if filtro.QtdParcelas != nil {
		query += fmt.Sprintf(" AND qtd_parcelas = $%d", paramID)
		args = append(args, *filtro.QtdParcelas)
		paramID++
	}

	if filtro.DataInicio != nil {
		query += fmt.Sprintf(" AND data_compra >= $%d", paramID)
		args = append(args, *filtro.DataInicio)
		paramID++
	}

	if filtro.DataFim != nil {
		query += fmt.Sprintf(" AND data_compra <= $%d", paramID)
		args = append(args, *filtro.DataFim)
		paramID++
	}

	if filtro.IDCartao != nil {
		query += fmt.Sprintf(" AND id_cartao = $%d", paramID)
		args = append(args, *filtro.IDCartao)
		paramID++
	}

	if filtro.IDCategoria != nil {
		query += fmt.Sprintf(" AND id_categoria = $%d", paramID)
		args = append(args, *filtro.IDCategoria)
		paramID++
	}

	if filtro.IDResponsavel != nil {
		query += fmt.Sprintf(" AND id_responsavel = $%d", paramID)
		args = append(args, *filtro.IDResponsavel)
		paramID++
	}

	if filtro.IDSubcategoria != nil {
		query += fmt.Sprintf(" AND id_subcategoria = $%d", paramID)
		args = append(args, *filtro.IDSubcategoria)
		paramID++
	}

	query += " ORDER BY data_compra DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compras []models.CompraParcelada

	for rows.Next() {
		var c models.CompraParcelada
		err := rows.Scan(
			&c.ID, &c.Descricao, &c.ValorTotal, &c.QtdParcelas, 
			&c.DataCompra, &c.IDCartao, &c.IDCategoria, 
			&c.IDResponsavel, &c.IDSubcategoria,
		)
		if err != nil {
			return nil, err
		}
		compras = append(compras, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return compras, nil
}

func CriarCompraParcelada(c models.CompraParcelada) (int, error) {
	query := `
		INSERT INTO finance.tb_compras_parceladas (
			descricao, valor_total, qtd_parcelas, data_compra, 
			id_cartao, id_categoria, id_responsavel, id_subcategoria
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING id
	`

	var idGerado int
	err := database.DB.QueryRow(
		query,
		c.Descricao,
		c.ValorTotal,
		c.QtdParcelas,
		c.DataCompra,
		c.IDCartao,
		c.IDCategoria,
		c.IDResponsavel,
		c.IDSubcategoria,
	).Scan(&idGerado)

	if err != nil {
		return 0, err
	}

	return idGerado, nil
}