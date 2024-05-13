package produtos

import (
	"github.com/hsxflowers/vendas-api/exceptions"
	"github.com/hsxflowers/vendas-api/produtos/domain"
	"github.com/labstack/gommon/log"
)

type Service struct {
	repo     domain.ProdutosDatabase
	consumer ServiceConsumer
}

func NewProdutosService(repo domain.ProdutosStorage, consumer ServiceConsumer) domain.Service {
	return &Service{
		repo:     repo,
		consumer: consumer,
	}
}

func (s *Service) Create(req *domain.ProdutosRequest) (*domain.CodigoRastreio, error) {
	log.Debug("[POST - Create] - Request processed by service ")

	for _, prod := range req.Produtos {
		produto := prod.ToProdutosDomain()

		quantidade, err := s.repo.VerificaDisponibilidade(produto)
		switch {
		case err != nil:
			return nil, err
		case quantidade == 0:
			log.Error("create_produto_service: produto not available")
			return nil, exceptions.New(exceptions.ErrProdutosNotFound, nil)
		case quantidade < produto.Quantidade:
			log.Error("create_produto_service: produto not available")
			return nil, exceptions.New(exceptions.ErrProdutosNotEnough, nil)
		}
	}

	res, err := s.GetPagamento(req)
	if err != nil {
		return nil, err
	}

	if res.Status {
		for _, prod := range req.Produtos {
			produto := prod.ToProdutosDomain()

			err := s.repo.RemoveQuantidade(produto)
			if err != nil {
				return nil, err
			}
		}
	} else {
		return nil, exceptions.New(exceptions.ErrPagamento, nil)
	}

	confirmacaoRastreio := s.toRastreio(&req.Produtos[0])

	confirmacao := s.SendEvent("req_envio", &domain.SendEventConfig{
		Key:       "rastreio",
		RequestId: "rastreio",
		Message:   confirmacaoRastreio,
	})
	if confirmacao != nil {
		return nil, confirmacao
	}

	codigo, err := s.ReadEventRastreio("cod_rastreio")
	if err != nil {
		return nil, err
	}

	return codigo, nil
}

func (s *Service) toPagamento(produto *domain.ProdutosRequest) *domain.ProdutosPagamento {
	var valor float64
	for _, prod := range produto.Produtos {
		valor = prod.Valor * float64(prod.Quantidade)
	}

	return &domain.ProdutosPagamento{
		Tipo:  produto.Produtos[0].Tipo,
		Valor: valor,
	}
}

func (s *Service) toProdutoResponse(produto *domain.Produto) *domain.ProdutosResponse {
	return &domain.ProdutosResponse{
		Tipo:       produto.Tipo,
		Nome:       produto.Nome,
		Valor:      produto.Valor,
		Quantidade: produto.Quantidade,
	}
}

func (s *Service) ReadEventPagamento(topico string) (*domain.ConfirmacaoPagamento, error) {
	msg, err := s.consumer.ReadMessagePagamento(topico)
	return msg, err
}

func (s *Service) ReadEventRastreio(topico string) (*domain.CodigoRastreio, error) {
	msg, err := s.consumer.ReadMessageRastreio(topico)
	return msg, err
}

func (s *Service) SendEvent(topico string, config *domain.SendEventConfig) error {
	return s.consumer.SendEvent(topico, config.Key, config.RequestId, config.Message)
}

func (s *Service) GetPagamento(req *domain.ProdutosRequest) (*domain.ConfirmacaoPagamento, error) {
	pagamento := s.toPagamento(req)

	err := s.SendEvent("req_cobra", &domain.SendEventConfig{
		Key:       "pagamento",
		RequestId: "pagamento",
		Message:   pagamento,
	})
	if err != nil {
		return nil, err
	}

	return s.ReadEventPagamento("status_pgto")

}

func (s *Service) toRastreio(produto *domain.Produto) *domain.ConfirmacaoRastreio {
	return &domain.ConfirmacaoRastreio{
		Tipo:   produto.Tipo,
		Status: true,
	}
}
