package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    time.Time
	Preco   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		// Caso não haja viagens no slice
		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: time.Date(2023, 5, 22, 0, 0, 0, 0, time.UTC), Preco: 200.0},
		{Origem: "Rio de Janeiro", Destino: "Salvador", Data: time.Date(2023, 6, 10, 0, 0, 0, 0, time.UTC), Preco: 350.0},
		{Origem: "São Paulo", Destino: "Brasília", Data: time.Date(2023, 7, 5, 0, 0, 0, 0, time.UTC), Preco: 150.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)
	fmt.Printf("A viagem mais cara é de %s para %s, na data %s, com preço de R$ %.2f.\n",
		viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Data.Format("02/01/2006"), viagemMaisCara.Preco)
}
