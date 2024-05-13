package exceptions

import "errors"

// Produtos errors
var ErrProdutosAlreadyExists = errors.New("produtos: produto already exists")
var ErrProdutosIdIsRequired = errors.New("produtos: produto_id is required and must be at least 3 characters")
var ErrTagIsRequired = errors.New("produtos: tag is required and must be at least 3 characters")
var ErrUrlIsNotValid = errors.New("produtos: url is not a valid URL")
var ErrTagIsNotValid = errors.New("produtos: tag is not valid")
var ErrProdutosNotFound = errors.New("produtos: produto not found")
var ErrBadData = errors.New("produtos: unprocessable json")
var ErrBadRequest = errors.New("produtos: can't update without valid field")
var ErrMissingField = errors.New("produtos: can't create without valid field")
var ErrInternalServer = errors.New("produtos: internal server error")

// Bind errors
var ErrBindDataOnCreateProdutos = errors.New("produtos: error on bind produto request when creating produto")
var ErrBindDataOnUpdateProdutos = errors.New("produtos: error on bind produto request when updating produto")

// DB errors
var ErrCreateProdutosInDB = errors.New("produtos: error creating produto in the database")
var ErrUpdateProdutosInDB = errors.New("produtos: error updating produto in the database")
var ErrDeleteProdutosInDB = errors.New("produtos: error deleting produto in the database")
var ErrGetProdutosInDB = errors.New("produtos: error getting produto in the database")
var ErrListProdutossInDB = errors.New("produtos: error listing produtos in the database")
var ErrProdutosNotEnough = errors.New("produtos: produto not enough")

var ErrReadingEvent = errors.New("produtos: failure to consume the kafka message")
var ErrUnprocessableValidation = errors.New("produtos: failure to parse the event values to validation map")
var ErrUnprocessableJson = errors.New("produtos: failure to parse the event values in domain's entitiy")
var ErrSendEvent = errors.New("produtos: failure to send kafka's event")
var ErrPagamento = errors.New("produtos: pagamento not approved")
