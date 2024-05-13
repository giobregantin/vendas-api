package produtos

import (
	"context"
	_ "log"

	"github.com/hsxflowers/vendas-api/exceptions"
	"github.com/hsxflowers/vendas-api/produtos/domain"
)

type ProdutoRepository struct {
	database domain.ProdutosStorage
}

func NewProdutosRepository(database domain.ProdutosDatabase) *ProdutoRepository {
	return &ProdutoRepository{
		database: database,
	}
}

func (repo *ProdutoRepository) VerificaDisponibilidade(ctx context.Context, produto *domain.Produto) (int, error) {
	quantidade, err := repo.database.VerificaDisponibilidade(ctx, produto)
	if err != nil {
		return 0, exceptions.New(exceptions.ErrInternalServer, err)
	}

	return quantidade, nil
}

func (repo *ProdutoRepository) RemoveQuantidade(ctx context.Context, produto *domain.Produto) error {
	err := repo.database.RemoveQuantidade(ctx, produto)
	if err != nil {
		return exceptions.New(exceptions.ErrInternalServer, err)
	}

	return nil
}
