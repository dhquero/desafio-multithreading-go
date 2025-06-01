package main

import (
	"fmt"
	"time"

	"github.com/dhquero/desafio-multithreading-go/api"
)

func main() {
	var searchCep string

	fmt.Print("CEP: ")
	fmt.Scanln(&searchCep)

	channelViaCEP := make(chan api.ViaCEP)
	channelBrasilAPI := make(chan api.BrasilAPI)

	go api.GetViaCEP(searchCep, channelViaCEP)
	go api.GetBrasilAPI(searchCep, channelBrasilAPI)

	select {
	case dataCEP := <-channelViaCEP:
		fmt.Printf("Via CEP. CEP: %s. Estado: %s. Cidade: %s. Bairro: %s. Rua: %s\n", dataCEP.Cep, dataCEP.Uf, dataCEP.Localidade, dataCEP.Bairro, dataCEP.Logradouro)
	case dataCEP := <-channelBrasilAPI:
		fmt.Printf("Brasil API. CEP: %s. Estado: %s. Cidade: %s. Bairro: %s. Rua: %s\n", dataCEP.Cep, dataCEP.State, dataCEP.City, dataCEP.Neighborhood, dataCEP.Street)
	case <-time.After(time.Second * 1):
		println("Tempo limite de processamento atingido")
	}
}
