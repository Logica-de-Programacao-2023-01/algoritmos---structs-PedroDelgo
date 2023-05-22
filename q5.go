package main

import (
	"fmt"
	"strings"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao string
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	resultado := make([]Playlist, 0)

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if strings.Contains(strings.ToLower(musica.Titulo), strings.ToLower(titulo)) {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	playlist1 := Playlist{
		Nome: "Playlist 1",
		Musicas: []Musica{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: "3:30"},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: "4:00"},
		},
	}

	playlist2 := Playlist{
		Nome: "Playlist 2",
		Musicas: []Musica{
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: "3:45"},
			{Titulo: "Música 4", Artista: "Artista 4", Duracao: "5:15"},
		},
	}

	playlist3 := Playlist{
		Nome: "Playlist 3",
		Musicas: []Musica{
			{Titulo: "Música 5", Artista: "Artista 5", Duracao: "4:20"},
			{Titulo: "Música 6", Artista: "Artista 6", Duracao: "3:10"},
		},
	}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	titulo := "música 2"
	resultado := buscarPlaylistsPorTitulo(titulo, playlists)

	if len(resultado) > 0 {
		fmt.Printf("Playlists com o título '%s':\n", titulo)
		for _, playlist := range resultado {
			fmt.Println(playlist.Nome)
		}
	} else {
		fmt.Printf("Nenhuma playlist encontrada com o título '%s'\n", titulo)
	}
}
