package domain

const MIN_LENGTH_CAT_ID = 3

type ProdutosStorage interface {
	VerificaDisponibilidade(produto *Produto) (int, error)
	RemoveQuantidade(produto *Produto) error
}

type ProdutosDatabase interface {
	VerificaDisponibilidade(produto *Produto) (int, error)
	RemoveQuantidade(produto *Produto) error
}

type Service interface {
	Create(produto *ProdutosRequest) (*CodigoRastreio, error)
}

func (u *Produto) ToProdutosDomain() *Produto {
	return &Produto{
		Tipo:       u.Tipo,
		Nome:       u.Nome,
		Valor:      u.Valor,
		Quantidade: u.Quantidade,
	}
}

func (u *Produto) ToProdutosResponse() *ProdutosResponse {
	if u != nil {
		return &ProdutosResponse{
			Tipo:       u.Tipo,
			Nome:       u.Nome,
			Valor:      u.Valor,
			Quantidade: u.Quantidade,
		}
	}
	return nil
}

type Produto struct {
	Tipo       string  `json:"tipo"`
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
}

type ProdutosRequest struct {
	Produtos []Produto `json:"produtos"`
}

type ProdutosResponse struct {
	Tipo       string  `json:"tipo"`
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
}

type ProdutosPagamento struct {
	Tipo  string  `json:"tipo"`
	Valor float64 `json:"valor"`
}

type ConfirmacaoPagamento struct {
	Status bool `json:"status"`
}

type SendEventConfig struct {
	Topic     string
	Key       string
	RequestId string
	AppName   string
	Message   interface{}
}

type CodigoRastreio struct {
	Tipo   string `json:"tipo"`
	Codigo string `json:"codigo"`
}

type ConfirmacaoRastreio struct {
	Tipo   string `json:"tipo"`
	Status bool   `json:"status"`
}
