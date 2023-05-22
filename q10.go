package main

import "fmt"

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	if len(f.Avaliacoes) == 0 {
		return 0.0
	}

	soma := 0
	for _, avaliacao := range f.Avaliacoes {
		soma += avaliacao
	}

	media := float64(soma) / float64(len(f.Avaliacoes))
	return media
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média das avaliações: %.2f\n", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "A Origem",
		Diretor:    "Christopher Nolan",
		Ano:        2010,
		Avaliacoes: []int{8, 9, 7, 10, 8},
	}

	filme.AdicionarAvaliacao(9)
	filme.RemoverAvaliacao(2)

	filme.ImprimirInformacoes()
}
