package infra

import (
	"encoding/json"
	"net/http"

	"github.com/jhonathann10/temperature-system/internal/infra/client/viacep"
	"github.com/jhonathann10/temperature-system/internal/infra/client/weatherapi"
	"github.com/jhonathann10/temperature-system/internal/usecase"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	viaCep  viacep.ViaCepInterface
	weather weatherapi.WeatherAPIInterface
}

func NewHandler(viaCep viacep.ViaCepInterface, weather weatherapi.WeatherAPIInterface) *Handler {
	return &Handler{
		viaCep:  viaCep,
		weather: weather,
	}
}

func (h *Handler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cep := r.URL.Query().Get("cep")

	if isCepInvalid(cep) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "invalid zipcode"})
		return
	}

	temperatureUseCase := usecase.NewTemperatureUseCase(h.viaCep, h.weather)
	localidade, err := temperatureUseCase.Execute(cep)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Message})
		return
	}

	errJson := json.NewEncoder(w).Encode(localidade)
	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusInternalServerError)
		return
	}
}

func isCepInvalid(cep string) bool {
	return len(cep) != 8
}
