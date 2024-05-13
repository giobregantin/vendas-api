package produtos

import (
	"context"

	"github.com/hsxflowers/vendas-api/produtos/domain"

	"github.com/hsxflowers/vendas-api/exceptions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProdutosHandler struct {
	ctx             context.Context
	produtosService domain.Service
}

func NewProdutosHandler(ctx context.Context, produtosService domain.Service) ProdutosHandler {
	return ProdutosHandler{
		ctx,
		produtosService,
	}
}

// Create
//
//	@Summary		Criação de gatos.
//	@Description	Endpoint que permite a criação de gatos.
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.ProdutosRequest	true	"body"
//	@Success		201		{object}	domain.ProdutosResponse
//	@Failure		422		"Unprocessable Json: Payload enviado com erro de syntax do json"
//	@Failure		400		"Erros de validação ou request indevido"
//	@Failure		500			"internal Server Error"
//	@Router			/channels [post]
func (h *ProdutosHandler) Create(c echo.Context) (*domain.CodigoRastreio, error) {
	req := new(domain.ProdutosRequest)

	if err := c.Bind(req); err != nil {
		log.Error("handler_create: error marshal produtos", err)
		return nil, exceptions.New(exceptions.ErrBadData, err)
	}

	codigo, err := h.produtosService.Create(h.ctx, req)
	if err != nil {
		return nil, err
	}

	return codigo, nil
}
