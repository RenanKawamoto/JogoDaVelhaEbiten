package JogoDaVelha

import (
	"image/color"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func QuemVenceu(vencedor string) *ebiten.Image {
	img := ebiten.NewImage(80, 80)
	caminhoDoArquivoDeFonte := "../../resources/fonts/dogica.ttf"
	arquivoDaFonteEmBytes, _ := ioutil.ReadFile(caminhoDoArquivoDeFonte)
	dogicaTTF, _ := opentype.Parse(arquivoDaFonteEmBytes)
	face, _ := opentype.NewFace(dogicaTTF, &opentype.FaceOptions{
		Size:    8,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if vencedor == "X" {
		text.Draw(img, "Jogador\n\n   X   \n\nVenceu", face, 12, 8, color.White)
	} else if vencedor == "O" {
		text.Draw(img, "Jogador\n\n   O   \n\nVenceu", face, 12, 8, color.White)
	} else {
		text.Draw(img, " Deu\n\nVelha", face, 20, 8, color.White)
	}
	return img
}
