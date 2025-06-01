package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func GetBrasilAPI(cep string, channelBrasilAPI chan<- BrasilAPI) {
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}

	var data BrasilAPI

	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}

	channelBrasilAPI <- data
}
