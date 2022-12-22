package Jogo

import (
	"fmt"
	"image/color"

	"github.com/RenanKawamoto/JogoDaVelhaEbiten/internal/JogoDaVelha"
	"github.com/hajimehoshi/ebiten/v2"
)

const LarguraDaTela int = 40
const AlturaDaTela int = 40

var campo JogoDaVelha.Campo
var matriz [3][3]string
var jogador JogoDaVelha.Jogador

type Game struct {
}

func init() {
	campo = JogoDaVelha.Campo{Largura: 32, Altura: 32, CorDoFundo: color.RGBA{255, 255, 255, 255}, CorDasLinhas: color.RGBA{0, 0, 0, 255}}
	jogador = JogoDaVelha.Jogador{XOuO: "O"}
	matriz = [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	campo.CriarCampo()
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	fmt.Println(x, y)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		xMatriz, yMatriz := jogador.ConverterPosicaoDoCursor(x, y)
		jogador.Jogar(&matriz, xMatriz, yMatriz)
		fmt.Println(xMatriz, yMatriz)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(4, 2)
	screen.DrawImage(campo.Desenhar(matriz), &op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 40, 40
}
