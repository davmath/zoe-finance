package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
	"fmt"
)

func BuscarComprasParceladas(filtro models.FiltroCompraParcelada) ([]models.CompraParcelada, error) {
	query := `
		SELECT 
			cp.id, cp.descricao, cp.valor_total, cp.qtd_parcelas, cp.data_compra, 
			cp.id_cartao, cc.nome,
			cp.id_categoria, c.nome_categoria,
			cp.id_responsavel, r.nome,
			cp.id_subcategoria, s.nome_subcategoria
		FROM finance.tb_compras_parceladas cp
		LEFT JOIN finance.tb_cartao_credito cc ON cp.id_cartao = cc.id
		LEFT JOIN finance.tb_categorias c ON cp.id_categoria = c.id
		LEFT JOIN finance.tb_responsavel_conta r ON cp.id_responsavel = r.id
		LEFT JOIN finance.tb_subcategorias s ON cp.id_subcategoria = s.id
		WHERE 1=1
	`
	
	var args []interface{}
	paramID := 1

	if filtro.Descricao != nil {
		query += fmt.Sprintf(" AND cp.descricao ILIKE $%d", paramID)
		args = append(args, "%"+*filtro.Descricao+"%")
		paramID++
	}

	if filtro.ValorTotalMin != nil {
		query += fmt.Sprintf(" AND cp.valor_total >= $%d", paramID)
		args = append(args, *filtro.ValorTotalMin)
		paramID++
	}

	if filtro.ValorTotalMax != nil {
		query += fmt.Sprintf(" AND cp.valor_total <= $%d", paramID)
		args = append(args, *filtro.ValorTotalMax)
		paramID++
	}

	if filtro.QtdParcelas != nil {
		query += fmt.Sprintf(" AND cp.qtd_parcelas = $%d", paramID)
		args = append(args, *filtro.QtdParcelas)
		paramID++
	}

	if filtro.DataInicio != nil {
		query += fmt.Sprintf(" AND cp.data_compra >= $%d", paramID)
		args = append(args, *filtro.DataInicio)
		paramID++
	}

	if filtro.DataFim != nil {
		query += fmt.Sprintf(" AND cp.data_compra <= $%d", paramID)
		args = append(args, *filtro.DataFim)
		paramID++
	}

	if filtro.IDCartao != nil {
		query += fmt.Sprintf(" AND cp.id_cartao = $%d", paramID)
		args = append(args, *filtro.IDCartao)
		paramID++
	}

	if filtro.IDCategoria != nil {
		query += fmt.Sprintf(" AND cp.id_categoria = $%d", paramID)
		args = append(args, *filtro.IDCategoria)
		paramID++
	}

	if filtro.IDResponsavel != nil {
		query += fmt.Sprintf(" AND cp.id_responsavel = $%d", paramID)
		args = append(args, *filtro.IDResponsavel)
		paramID++
	}

	if filtro.IDSubcategoria != nil {
		query += fmt.Sprintf(" AND cp.id_subcategoria = $%d", paramID)
		args = append(args, *filtro.IDSubcategoria)
		paramID++
	}

	query += " ORDER BY cp.data_compra DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compras []models.CompraParcelada

	for rows.Next() {
		var cp models.CompraParcelada
		err := rows.Scan(
			&cp.ID, &cp.Descricao, &cp.ValorTotal, &cp.QtdParcelas, &cp.DataCompra, 
			&cp.IDCartao, &cp.NomeCartao,
			&cp.IDCategoria, &cp.NomeCategoria,
			&cp.IDResponsavel, &cp.NomeResponsavel,
			&cp.IDSubcategoria, &cp.NomeSubcategoria,
		)
		if err != nil {
			return nil, err
		}
		compras = append(compras, cp)
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