package main

import "fmt"

type Animal struct {
	Nome    string
	Espécie string
	Idade   int
	Som     string
}

func (a *Animal) ModificarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformações() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Espécie:", a.Espécie)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Som:", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Rex",
		Espécie: "Cachorro",
		Idade:   5,
		Som:     "Au Au",
	}

	animal.ImprimirInformações()

	fmt.Println("Modificando o som do animal...")
	animal.ModificarSom("Woof Woof")

	fmt.Println("Novas informações do animal:")
	animal.ImprimirInformações()
}
