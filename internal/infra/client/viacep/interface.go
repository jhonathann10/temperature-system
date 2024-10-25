package viacep

import "github.com/jhonathann10/temperature-system/internal/httperror"

type ViaCepInterface interface {
	GetAddressByCEP(cep string) (*LocalidadeCEP, *httperror.HttpError)
}
