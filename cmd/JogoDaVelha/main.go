package main

import (
	"log"

	"github.com/RenanKawamoto/JogoDaVelhaEbiten/internal/Jogo"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Jogo da Velha")
	if err := ebiten.RunGame(&Jogo.Game{}); err != nil {
		log.Fatal(err)
	}
}
