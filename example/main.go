package main

import (
	"fmt"

	musicbrainz_wrapper "github.com/Victiniiiii/musicbrainz_wrapper/pkg"
)

func main() {
	songs := []string{
		"Tamirci Çırağı",
		"Despacito",
		"La Vie en Rose",
		"昨日よりもっと好き",
		"güzel günler göreceğiz çocuklar",
		"Ahmet Beyin Ceketi",
		"Bohemian Rhapsody",
		"Gangnam Style",
		"Hallelujah",
		"Lágrimas Negras",
		"Feliz Navidad",
		"Merci Chérie",
		"Dragostea Din Tei",
		"Non, je ne regrette rien",
		"O Sole Mio",
		"Ai Se Eu Te Pego",
		"Kimi no Na wa",
		"Bésame Mucho",
		"Volare",
		"Kazma",
		"Sakura Sakura",
		"El Perdón",
		"La Cumparsita",
		"Für Elise",
		"Kommt ein Vogel geflogen",
		"Yeniden de sevebiliriz akdeniz",
		"Viva La Vida",
	}

	for _, s := range songs {
		artist, genre, lang := musicbrainz_wrapper.DetectMetadata(s)
		fmt.Printf("Song: %q, Artist: %s, Genre: %s, Language: %s\n", s, artist, genre, lang)
	}
}
