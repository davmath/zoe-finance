package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
	"errors"
	"fmt"
)

func BuscarTransacoes(filtro models.FiltroTransacao) ([]models.Transacao, error) {
	query := `
		SELECT 
			t.id, 
			t.descricao, 
			t.valor, 
			t.data_transacao, 
			t.efetivada,
			t.id_categoria, c.nome_categoria,
			t.id_subcategoria, s.nome_subcategoria,
			t.id_responsavel, r.nome,
			t.id_conta_bancaria, cb.nome,
			t.id_cartao_credito, cc.nome,
			t.id_conta_destino, cd.nome,
			t.id_compra_parcelada, cp.descricao
		FROM finance.tb_transacoes t
		LEFT JOIN finance.tb_categorias c ON t.id_categoria = c.id
		LEFT JOIN finance.tb_subcategorias s ON t.id_subcategoria = s.id
		LEFT JOIN finance.tb_responsavel_conta r ON t.id_responsavel = r.id
		LEFT JOIN finance.tb_contas_bancarias cb ON t.id_conta_bancaria = cb.id
		LEFT JOIN finance.tb_cartao_credito cc ON t.id_cartao_credito = cc.id
		LEFT JOIN finance.tb_contas_bancarias cd ON t.id_conta_destino = cd.id
		LEFT JOIN finance.tb_compras_parceladas cp ON t.id_compra_parcelada = cp.id
		WHERE 1=1
	`

	var args []interface{}
	paramID := 1

	if filtro.Descricao != nil {
		query += fmt.Sprintf(" AND t.descricao ILIKE $%d", paramID)
		args = append(args, "%"+*filtro.Descricao+"%")
		paramID++
	}

	if filtro.ValorMin != nil {
		query += fmt.Sprintf(" AND t.valor >= $%d", paramID)
		args = append(args, *filtro.ValorMin)
		paramID++
	}

	if filtro.ValorMax != nil {
		query += fmt.Sprintf(" AND t.valor <= $%d", paramID)
		args = append(args, *filtro.ValorMax)
		paramID++
	}

	if filtro.DataInicio != nil {
		query += fmt.Sprintf(" AND t.data_transacao >= $%d", paramID)
		args = append(args, *filtro.DataInicio)
		paramID++
	}

	if filtro.DataFim != nil {
		query += fmt.Sprintf(" AND t.data_transacao <= $%d", paramID)
		args = append(args, *filtro.DataFim)
		paramID++
	}

	if filtro.IDCategoria != nil {
		query += fmt.Sprintf(" AND t.id_categoria = $%d", paramID)
		args = append(args, *filtro.IDCategoria)
		paramID++
	}

	if filtro.IDResponsavel != nil {
		query += fmt.Sprintf(" AND t.id_responsavel = $%d", paramID)
		args = append(args, *filtro.IDResponsavel)
		paramID++
	}

	if filtro.IDContaBancaria != nil {
		query += fmt.Sprintf(" AND t.id_conta_bancaria = $%d", paramID)
		args = append(args, *filtro.IDContaBancaria)
		paramID++
	}

	if filtro.IDCartaoCredito != nil {
		query += fmt.Sprintf(" AND t.id_cartao_credito = $%d", paramID)
		args = append(args, *filtro.IDCartaoCredito)
		paramID++
	}

	query += " ORDER BY t.data_transacao DESC"

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transacoes []models.Transacao

	for rows.Next() {
		var t models.Transacao
		
		err := rows.Scan(
			&t.ID, 
			&t.Descricao, 
			&t.Valor, 
			&t.DataTransacao, 
			&t.Efetivada,
			&t.IDCategoria, &t.NomeCategoria,
			&t.IDSubcategoria, &t.NomeSubcategoria,
			&t.IDResponsavel, &t.NomeResponsavel,
			&t.IDContaBancaria, &t.NomeContaBancaria,
			&t.IDCartaoCredito, &t.NomeCartaoCredito,
			&t.IDContaDestino, &t.NomeContaDestino,
			&t.IDCompraParcelada, &t.DescCompraParcelada,
		)
		
		if err != nil {
			return nil, err
		}
		transacoes = append(transacoes, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transacoes, nil
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

func AtualizarTransacao(id int, campos map[string]interface{}) error {
	if len(campos) == 0 {
		return nil
	}

	query := "UPDATE finance.tb_transacoes SET "
	var args []interface{}
	paramID := 1

	for coluna, valor := range campos {
		query += fmt.Sprintf("%s = %d, ", coluna, paramID)
		args = append(args, valor)
		paramID++
	}

	query = query[:len(query) - 2]
	query += fmt.Sprintf(" WHERE id = $%d", paramID)
	args = append(args, id)

	_, err := database.DB.Exec(query, args...)
	return err
}

func DeletarTransacao(id int) error {
	query := "DELETE FROM finance.TB_TRANSACOES WHERE id = $1"

	resultado, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}

	if linhasAfetadas == 0 {
		return errors.New("registro_nao_encontrado")
	}

	return nil
}

