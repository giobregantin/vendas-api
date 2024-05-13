package exceptions

import (
	"errors"
	"net/http"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleException(err error) ErrResponse {
	customErr, ok := err.(*Error)
	if !ok {
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	switch err != nil {
	case
		errors.Is(customErr.CustomErr, ErrProdutosIdIsRequired),
		errors.Is(customErr.CustomErr, ErrTagIsRequired),
		errors.Is(customErr.CustomErr, ErrUrlIsNotValid),
		errors.Is(customErr.CustomErr, ErrTagIsNotValid),
		errors.Is(customErr.CustomErr, ErrBadRequest):
		return ErrResponse{
			Code:    http.StatusBadRequest,
			Message: customErr.CustomErr.Error(),
		}
	case
		errors.Is(customErr.CustomErr, ErrCreateProdutosInDB),
		errors.Is(customErr.CustomErr, ErrGetProdutosInDB),
		errors.Is(customErr.CustomErr, ErrListProdutossInDB),
		errors.Is(customErr.CustomErr, ErrUpdateProdutosInDB),
		errors.Is(customErr.CustomErr, ErrDeleteProdutosInDB),
		errors.Is(customErr.CustomErr, ErrBindDataOnCreateProdutos),
		errors.Is(customErr.CustomErr, ErrBindDataOnUpdateProdutos),
		errors.Is(customErr.CustomErr, ErrProdutosNotEnough):
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: customErr.CustomErr.Error(),
		}
	case
		errors.Is(customErr.CustomErr, ErrBadData):
		return ErrResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: customErr.CustomErr.Error(),
		}
	case
		errors.Is(customErr.CustomErr, ErrProdutosNotFound):
		return ErrResponse{
			Code:    http.StatusNotFound,
			Message: customErr.CustomErr.Error(),
		}
	default:
		return ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		}
	}
}
