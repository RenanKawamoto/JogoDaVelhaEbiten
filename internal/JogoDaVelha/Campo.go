package JogoDaVelha

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func (c *Campo) CriarCampo() {
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
}

func (c *Campo) desenharForma(conteudo string) *ebiten.Image {
	if conteudo == "O" {
		img, _, _ := ebitenutil.NewImageFromFile("../../resources/imgs/Circulo.png")
		return img
	} else if conteudo == "X" {
		img, _, _ := ebitenutil.NewImageFromFile("../../resources/imgs/X.png")
		return img
	}
	img := ebiten.NewImage(1, 1)
	return img
}

func (c *Campo) posicionarForma(indexDaMatriz, indexDoArray int) *ebiten.DrawImageOptions {
	op := ebiten.DrawImageOptions{}
	switch indexDaMatriz {
	case 0:
		switch indexDoArray {
		case 0:
			op.GeoM.Translate(1, 1)
			return &op
		case 1:
			op.GeoM.Translate(12, 1)
			return &op
		case 2:
			op.GeoM.Translate(23, 1)
			return &op
		}
	case 1:
		switch indexDoArray {
		case 0:
			op.GeoM.Translate(1, 12)
			return &op
		case 1:
			op.GeoM.Translate(12, 12)
			return &op
		case 2:
			op.GeoM.Translate(23, 12)
			return &op
		}
	case 2:
		switch indexDoArray {
		case 0:
			op.GeoM.Translate(1, 23)
			return &op
		case 1:
			op.GeoM.Translate(12, 23)
			return &op
		case 2:
			op.GeoM.Translate(23, 23)
			return &op
		}
	}
	return &op
}

func (c *Campo) Desenhar(matrizDoJogo [3][3]string) *ebiten.Image {
	for indexDaMatriz, conteudoDaMatriz := range matrizDoJogo {
		for indexDoArray, conteudoDoArray := range conteudoDaMatriz {
			c.Imagem.DrawImage(c.desenharForma(conteudoDoArray), c.posicionarForma(indexDaMatriz, indexDoArray))
		}
	}
	return c.Imagem
}

func (c *Campo) Limpar(matrizDoJogo *[3][3]string) {
	*matrizDoJogo = [3][3]string{}
}
