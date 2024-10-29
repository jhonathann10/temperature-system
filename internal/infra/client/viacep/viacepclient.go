package viacep

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonathann10/temperature-system/internal/httperror"
)

type LocalidadeCEP struct {
	Localidade string `json:"localidade"`
}

type ViaCEPClient struct {
	BaseURL string
}

func NewViaCEPClient(baseURL string) *ViaCEPClient {
	return &ViaCEPClient{
		BaseURL: baseURL,
	}
}

func (v *ViaCEPClient) GetAddressByCEP(cep string) (*LocalidadeCEP, *httperror.HttpError) {
	localidade := &LocalidadeCEP{}
	resp, err := http.Get(fmt.Sprintf("%s/%s/json/", v.BaseURL, cep))
	if err != nil {
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(localidade)
	if err != nil {
		return nil, &httperror.HttpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return localidade, nil
}
