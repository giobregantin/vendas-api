package produtos

import (
	_ "log"

	"github.com/hsxflowers/vendas-api/produtos/domain"
	"github.com/hsxflowers/vendas-api/exceptions"
	"github.com/labstack/gommon/log"
)

type ProdutoRepository struct {
	database domain.ProdutosStorage
}

func NewProdutosRepository(database domain.ProdutosDatabase) *ProdutoRepository {
	return &ProdutoRepository{
		database: database,
	}
}

func (repo *ProdutoRepository) VerificaDisponibilidade(produto *domain.Produto) (int, error) {
	quantidade, err := repo.database.VerificaDisponibilidade(produto)
	if err != nil {
		log.Error("cat_repo: error on create produto in the database", err)
		return 0, exceptions.New(exceptions.ErrInternalServer, err)
	}

	return quantidade, nil
}

func (repo *ProdutoRepository) RemoveQuantidade(produto *domain.Produto) error {
	err := repo.database.RemoveQuantidade(produto)
	if err != nil {
		log.Error("cat_repo: error on create produto in the database", err)
		return exceptions.New(exceptions.ErrInternalServer, err)
	}

	return nil
}
