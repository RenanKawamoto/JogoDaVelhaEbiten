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
