package Jogo

import (
	"image/color"

	"github.com/RenanKawamoto/JogoDaVelhaEbiten/internal/JogoDaVelha"
	"github.com/hajimehoshi/ebiten/v2"
)

const LarguraDaTela int = 40
const AlturaDaTela int = 40

var campo = JogoDaVelha.Campo{Largura: 32, Altura: 32, CorDoFundo: color.RGBA{255, 255, 255, 255}, CorDasLinhas: color.RGBA{0, 0, 0, 255}}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 5)
	screen.DrawImage(campo.CriarCampo(), &op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 40, 40
}
