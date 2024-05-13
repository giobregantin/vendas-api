package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hsxflowers/vendas-api/exceptions"
	"github.com/hsxflowers/vendas-api/produtos/domain"
	"github.com/labstack/gommon/log"
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) VerificaDisponibilidade(ctx context.Context, produto *domain.Produto) (int, error) {
	var quantidade int

	query := "SELECT estoque_prod FROM produtos WHERE nome_prod = $1"
	row := s.db.QueryRowContext(ctx, query, produto.Nome)

	err := row.Scan(&quantidade)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, exceptions.New(exceptions.ErrProdutosNotFound, err)
		}
		log.Error("Error fetching produto from database: ", err)
		return 0, exceptions.New(exceptions.ErrInternalServer, err)
	}

	return quantidade, nil
}

func (s *SQLStore) RemoveQuantidade(ctx context.Context, produto *domain.Produto) error {
	query := "UPDATE produtos SET estoque_prod = estoque_prod - $1 WHERE nome_prod = $2"
	_, err := s.db.ExecContext(ctx, query, produto.Quantidade, produto.Nome)
	if err != nil {
		log.Error("Error updating produto in database: ", err)
		return exceptions.New(exceptions.ErrInternalServer, err)
	}

	return nil
}