package repository

import (
	"davmath/zoe-finance/database"
	"davmath/zoe-finance/models"
)

func ObterResumo(filtro models.FiltroDashboard) (models.ResumoDashboard, error) {
	var resumo models.ResumoDashboard

	querySaldo := `SELECT COALESCE(SUM(montante), 0) FROM finance.tb_contas_bancarias`
	err := database.DB.QueryRow(querySaldo).Scan(&resumo.SaldoAtualContas)
	if err != nil {
		return resumo, err
	}


	queryFluxo := `
		SELECT 
			COALESCE(SUM(CASE WHEN c.tipo = 'Receita' THEN t.valor ELSE 0 END), 0) as total_receitas,
			COALESCE(SUM(CASE WHEN c.tipo = 'Despesa' THEN t.valor ELSE 0 END), 0) as total_despesas
		FROM finance.tb_transacoes t
		JOIN finance.tb_categorias c ON t.id_categoria = c.id
		WHERE t.data_transacao >= $1 AND t.data_transacao <= $2
	`

	err = database.DB.QueryRow(queryFluxo, filtro.DataInicio, filtro.DataFim).Scan(
		&resumo.TotalReceitas,
		&resumo.TotalDespesas,
	)
	if err != nil {
		return resumo, err
	}

	resumo.BalancoPeriodo = resumo.TotalReceitas - resumo.TotalDespesas

	return resumo, nil
}

func ObterDespesasPorCategoria(filtro models.FiltroDashboard) ([]models.DespesaPorCategoria, error) {
	query := `
		SELECT 
			c.id, 
			c.nome_categoria, 
			COALESCE(SUM(t.valor), 0) as total
		FROM finance.tb_transacoes t
		JOIN finance.tb_categorias c ON t.id_categoria = c.id
		WHERE c.tipo = 'Despesa' 
		  AND t.data_transacao >= $1 
		  AND t.data_transacao <= $2
		GROUP BY c.id, c.nome_categoria
		ORDER BY total DESC
	`

	rows, err := database.DB.Query(query, filtro.DataInicio, filtro.DataFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var despesas []models.DespesaPorCategoria

	for rows.Next() {
		var d models.DespesaPorCategoria
		err := rows.Scan(&d.IDCategoria, &d.NomeCategoria, &d.Total)
		if err != nil {
			return nil, err
		}
		despesas = append(despesas, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if despesas == nil {
		despesas = []models.DespesaPorCategoria{}
	}

	return despesas, nil
}