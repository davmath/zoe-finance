CREATE OR REPLACE FUNCTION finance.fn_gerar_parcelas_transacao()
RETURNS TRIGGER AS $$
DECLARE
	v_valor_parcela DECIMAL(10,2);
	v_valor_ultima_parcela DECIMAL(10,2);
	v_data_vencimento TIMESTAMP;
	i INT;
BEGIN
	v_valor_parcela := ROUND(NEW.VALOR_TOTAL / NEW.QTD_PARCELAS, 2);
	v_valor_ultima_parcela := NEW.VALOR_TOTAL - (v_valor_parcela * (NEW.QTD_PARCELAS - 1));
	
	FOR i IN 1..NEW.QTD_PARCELAS LOOP
		v_data_vencimento := NEW.DATA_COMPRA + ((i - 1) || ' month')::INTERVAL;
		INSERT INTO finance.TB_TRANSACOES (
            DESCRICAO,
            VALOR,
            DATA_TRANSACAO,
            ID_CATEGORIA,
			ID_SUBCATEGORIA,
            ID_RESPONSAVEL,
            ID_CARTAO_CREDITO,
            ID_COMPRA_PARCELADA,
            EFETIVADA
        ) VALUES (
            NEW.DESCRICAO || ' (' || i || '/' || NEW.QTD_PARCELAS || ')',
            CASE WHEN i = NEW.QTD_PARCELAS THEN v_valor_ultima_parcela ELSE v_valor_parcela END,
            v_data_vencimento,
            NEW.ID_CATEGORIA,
			NEW.ID_SUBCATEGORIA,
            NEW.ID_RESPONSAVEL,
            NEW.ID_CARTAO,
            NEW.ID, -- ID da compra parcelada que acabou de ser gerado
            FALSE -- Parcelas futuras nascem como não pagas
        );
    END LOOP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trg_apos_inserir_compra_parcelada
AFTER INSERT ON finance.TB_COMPRAS_PARCELADAS
FOR EACH ROW
EXECUTE FUNCTION finance.fn_gerar_parcelas_transacao();

select DISTINCT id_compra_parcelada from finance.TB_TRANSACOES WHERE ID_RESPONSAVEL = 3; -- where efetivada is false ORDER BY ID;
select * from finance.TB_COMPRAS_PARCELADAS;
select * from finance.TB_CARTAO_CREDITO;

UPDATE FINANCE.TB_COMPRAS_PARCELADAS SET ID_RESPONSAVEL = 3 where id = 5;

ALTER TABLE finance.TB_COMPRAS_PARCELADAS
ADD COLUMN ID_SUBCATEGORIA INT REFERENCES finance.TB_SUBCATEGORIAS(ID);

ALTER TABLE finance.TB_TRANSACOES 
DROP CONSTRAINT chk_origem_destino_diferentes;

ALTER TABLE finance.TB_TRANSACOES 
ADD CONSTRAINT chk_origem_destino_diferentes 
CHECK (
    ID_CONTA_BANCARIA IS NULL OR 
    ID_CONTA_DESTINO IS NULL OR 
    ID_CONTA_BANCARIA != ID_CONTA_DESTINO
);

INSERT INTO finance.TB_COMPRAS_PARCELADAS (
    DESCRICAO, 
    VALOR_TOTAL, 
    QTD_PARCELAS, 
    DATA_COMPRA, 
    ID_CARTAO, 
    ID_CATEGORIA,
	ID_SUBCATEGORIA,
    ID_RESPONSAVEL
) VALUES (
    'Manual Saúde - Cabelo', 
    553.14, 
    6, 
    '2025-12-28 23:00:00', 
    1, 
    7,
	31,
    2
);

select * from finance.tb_categorias;
select * from finance.tb_subcategorias;
select id, descricao, efetivada from finance.tb_transacoes where id_compra_parcelada = 10 order by id asc;
select * from finance.TB_COMPRAS_PARCELADAS order by data_compra asc;
select * from finance.tb_cartao_credito;

update finance.tb_transacoes set efetivada = True where id in (56, 57, 58, 59);