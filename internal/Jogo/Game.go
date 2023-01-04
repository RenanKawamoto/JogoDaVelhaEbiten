package Jogo

import (
	"image/color"

	"github.com/RenanKawamoto/JogoDaVelhaEbiten/internal/JogoDaVelha"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var LarguraDaTela int = 40
var AlturaDaTela int = 40

var campo JogoDaVelha.Campo
var matriz [3][3]string
var jogador JogoDaVelha.Jogador
var alterarTelaParaResultado bool
var ganhador string

type Game struct {
}

func init() {
	campo = JogoDaVelha.Campo{Largura: 32, Altura: 32, CorDoFundo: color.RGBA{255, 255, 255, 255}, CorDasLinhas: color.RGBA{0, 0, 0, 255}}
	jogador = JogoDaVelha.Jogador{XOuO: "O"}
	matriz = [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	campo.CriarCampo()
	alterarTelaParaResultado = false
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()

	if JogoDaVelha.GanharJogo(matriz) != "" || !JogoDaVelha.AindaNaoFinalizouOJogo(matriz) {
		alterarTelaParaResultado = true
		LarguraDaTela = 80
		ganhador = JogoDaVelha.GanharJogo(matriz)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !alterarTelaParaResultado {
		xMatriz, yMatriz := JogoDaVelha.ConverterPosicaoDoCursor(x, y)
		jogador.Jogar(&matriz, xMatriz, yMatriz)
	}

	if alterarTelaParaResultado && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		LarguraDaTela = 40
		campo.Limpar(&matriz)
		campo.CriarCampo()
		alterarTelaParaResultado = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(4, 4)
	if alterarTelaParaResultado {
		screen.DrawImage(JogoDaVelha.QuemVenceu(ganhador), nil)
	} else {
		screen.DrawImage(campo.Desenhar(matriz), &op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return LarguraDaTela, AlturaDaTela
}
