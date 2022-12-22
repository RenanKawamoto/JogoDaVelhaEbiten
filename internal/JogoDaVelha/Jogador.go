package JogoDaVelha

type Jogador struct {
	XOuO string
}

func (j *Jogador) trocarJogador() {
	if j.XOuO == "X" {
		j.XOuO = "O"
		return
	}
	j.XOuO = "X"
}

func (j *Jogador) Jogar(matriz *[3][3]string, linha, coluna int) {
	for indexMatriz, conteudoDaMatriz := range matriz {
		for indexArray, conteudoDoArray := range conteudoDaMatriz {
			if indexMatriz == linha && indexArray == coluna && conteudoDoArray != "X" && conteudoDoArray != "O" {
				matriz[linha][coluna] = j.XOuO
				j.trocarJogador()
				return
			}
		}
	}
}

func (j *Jogador) ConverterPosicaoDoCursor(x, y int) (int, int) {
	if x >= 0 && x <= 13 && y >= 0 && y <= 11 {
		return 0, 0
	}
	if x >= 13 && x <= 24 && y >= 0 && y <= 11 {
		return 0, 1
	}
	if x >= 26 && x <= 35 && y >= 0 && y <= 11 {
		return 0, 2
	}
	return 2, 2
}
