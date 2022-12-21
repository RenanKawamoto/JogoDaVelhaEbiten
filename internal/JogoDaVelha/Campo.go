package JogoDaVelha

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const linhaHorizontal = true
const linhaVertical = false
const descotoAreaQueALinhaOcupa = 1

type Campo struct {
	Largura      int
	Altura       int
	CorDoFundo   color.RGBA
	CorDasLinhas color.RGBA
	Imagem       *ebiten.Image
}

func (c *Campo) criarLinha(verticalOuHorizontal bool) *ebiten.Image {
	if verticalOuHorizontal {
		linha := ebiten.NewImage(c.Largura, 1)
		linha.Fill(c.CorDasLinhas)
		return linha
	}
	linha := ebiten.NewImage(1, c.Altura)
	linha.Fill(c.CorDasLinhas)
	return linha
}

func (c *Campo) posicionarLinha(verticalOuHorizontal bool, posicao float64) *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	if verticalOuHorizontal {
		op.GeoM.Translate(0, posicao)
		return &op
	}
	op.GeoM.Translate(posicao, 0)
	return &op
}

func (c *Campo) CriarCampo() *ebiten.Image {
	c.Imagem = ebiten.NewImage(c.Largura, c.Altura)
	c.Imagem.Fill(c.CorDoFundo)

	linhaVertical1 := c.criarLinha(linhaVertical)
	linhaVertical2 := c.criarLinha(linhaVertical)
	linhaHorizontal1 := c.criarLinha(linhaHorizontal)
	linhaHorizontal2 := c.criarLinha(linhaHorizontal)

	c.Imagem.DrawImage(linhaHorizontal1, c.posicionarLinha(linhaHorizontal, float64(c.Largura/3)))
	c.Imagem.DrawImage(linhaHorizontal2, c.posicionarLinha(linhaHorizontal, float64(c.Largura-c.Largura/3-descotoAreaQueALinhaOcupa)))
	c.Imagem.DrawImage(linhaVertical1, c.posicionarLinha(linhaVertical, float64(c.Altura/3)))
	c.Imagem.DrawImage(linhaVertical2, c.posicionarLinha(linhaVertical, float64(c.Altura-c.Altura/3-descotoAreaQueALinhaOcupa)))

	return c.Imagem
}

func (c *Campo) Desenhar(sliceDoJogo []bool) *ebiten.Image {

}
