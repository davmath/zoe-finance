package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
	"fmt"
)

func BuscarTransacoes(filtro models.FiltroTransacao) ([]models.Transacao, error) {
	query := `
			SELECT
				ID,
				DESCRICAO,
				VALOR,
				DATA_TRANSACAO,
				ID_CATEGORIA,
				ID_SUBCATEGORIA,
				ID_RESPONSAVEL,
				ID_CONTA_BANCARIA,
				ID_CARTAO_CREDITO,
				ID_CONTA_DESTINO,
				ID_COMPRA_PARCELADA,
				EFETIVADA
			FROM finance.TB_TRANSACOES
			WHERE 1=1
	`

	var args []interface{}
	paramID := 1

	if filtro.Descricao != nil {
		query += fmt.Sprintf(" AND DESCRICAO ILIKE $%d", paramID)
		args = append(args, "%"+*filtro.Descricao+"%")
		paramID++
	}

	if filtro.ValorMin != nil {
		query += fmt.Sprintf(" AND VALOR >= $%d", paramID)
		args = append(args, *filtro.ValorMin)
		paramID++
	}

	if filtro.ValorMax != nil {
		query += fmt.Sprintf(" AND VALOR <= $%d", paramID)
		args = append(args, *filtro.ValorMax)
		paramID++
	}

	if filtro.DataInicio != nil {
		query += fmt.Sprintf(" AND DATA_TRANSACAO >= $%d", paramID)
		args = append(args, *filtro.DataInicio)
		paramID++
	}

	if filtro.DataFim != nil {
		query += fmt.Sprintf(" AND DATA_TRANSACAO <= $%d", paramID)
		args = append(args, *filtro.DataFim)
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
	if filtro.Efetivada != nil {
		query += fmt.Sprintf(" AND efetivada = $%d", paramID)
		args = append(args, *filtro.Efetivada)
		paramID++
	}

	if filtro.IDContaBancaria != nil {
		query += fmt.Sprintf(" AND ID_CONTA_BANCARIA = $%d", paramID)
		args = append(args, *filtro.IDContaBancaria)
		paramID++
	}

	if filtro.IDCartaoCredito != nil {
		query += fmt.Sprintf(" AND ID_CARTAO_CREDITO = $%d", paramID)
		args = append(args, *filtro.IDCartaoCredito)
		paramID++
	}

	if filtro.IDContaDestino != nil {
		query += fmt.Sprintf(" AND ID_CONTA_DESTINO = $%d", paramID)
		args = append(args, *filtro.IDContaDestino)
		paramID++
	}

	if filtro.IDCompraParcelada != nil {
		query += fmt.Sprintf(" AND ID_COMPRA_PARCELADA = $%d", paramID)
		args = append(args, *filtro.IDCompraParcelada)
		paramID++
	}

	query += " ORDER BY DATA_TRANSACAO DESC"

	return executarConsulta(query, args)
}

func executarConsulta(query string, args []interface{}) ([]models.Transacao, error) {
	var transacoes []models.Transacao

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return  nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Transacao

		err := rows.Scan(
			&t.ID,
			&t.Descricao,
			&t.Valor,
			&t.DataTransacao,
			&t.IDCategoria,
			&t.IDSubcategoria,
			&t.IDResponsavel,
			&t.IDContaBancaria,
			&t.IDCartaoCredito,
			&t.IDContaDestino,
			&t.IDCompraParcelada,
			&t.Efetivada,
		)

		if err != nil{
			return nil, err
		}

		transacoes = append(transacoes, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transacoes, nil
}

func CriarTransacao(t models.Transacao) (int, error) {
	query := `
			INSERT INTO finance.tb_transacoes (
				descricao,
				valor,
				data_transacao,
				id_categoria,
				id_subcategoria,
				id_responsavel,
				id_conta_bancaria,
				id_cartao_credito,
				id_conta_destino,
				id_compra_parcelada,
				efetivada
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 
			) RETURNING id
	`

	var idGerado int

	err := database.DB.QueryRow(
		query,
		t.Descricao,
		t.Valor,
		t.DataTransacao,
		t.IDCategoria,
		t.IDSubcategoria,
		t.IDResponsavel,
		t.IDContaBancaria,
		t.IDCartaoCredito,
		t.IDContaDestino,
		t.IDCompraParcelada,
		t.Efetivada,
	).Scan(&idGerado)

	if err != nil {
		return 0, err
	}

	return idGerado, nil
}