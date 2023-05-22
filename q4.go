package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Nome da Playlist: %s\n", p.Nome)

	var duracaoTotal time.Duration

	for _, musica := range p.Musicas {
		fmt.Printf("Título: %s\n", musica.Titulo)
		fmt.Printf("Artista: %s\n", musica.Artista)
		fmt.Printf("Duração: %s\n", musica.Duracao)

		duracaoTotal += musica.Duracao
	}

	fmt.Printf("Duração Total da Playlist: %s\n", duracaoTotal)
}
