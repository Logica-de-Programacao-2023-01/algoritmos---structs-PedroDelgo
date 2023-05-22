package main

import "fmt"

type Funcionário struct {
	Nome    string
	Salário float64
	Idade   int
}

func (f *Funcionário) AumentarSalário(porcentagem float64) {
	aumento := f.Salário * (porcentagem / 100)
	f.Salário += aumento
}

func (f *Funcionário) DiminuirSalário(porcentagem float64) {
	redução := f.Salário * (porcentagem / 100)
	f.Salário -= redução
}

func (f *Funcionário) TempoDeServiço() int {
	idadeInicial := 18
	idadeAtual := f.Idade
	tempoDeServiço := idadeAtual - idadeInicial
	return tempoDeServiço
}

func main() {
	funcionário := Funcionário{
		Nome:    "João",
		Salário: 5000.0,
		Idade:   30,
	}

	fmt.Printf("Salário inicial do funcionário: R$ %.2f\n", funcionário.Salário)
	funcionário.AumentarSalário(10.0)
	fmt.Printf("Salário após aumento: R$ %.2f\n", funcionário.Salário)
	funcionário.DiminuirSalário(5.0)
	fmt.Printf("Salário após redução: R$ %.2f\n", funcionário.Salário)

	tempoDeServiço := funcionário.TempoDeServiço()
	fmt.Printf("Tempo de serviço do funcionário: %d anos\n", tempoDeServiço)
}
