package JogoDaVelha

func ConverterPosicaoDoCursor(x, y int) (int, int) {
	if x >= 0 && x <= 13 && y >= 0 && y <= 11 {
		return 0, 0
	}
	if x >= 13 && x <= 24 && y >= 0 && y <= 11 {
		return 0, 1
	}
	if x >= 26 && x <= 35 && y >= 0 && y <= 11 {
		return 0, 2
	}
	if x >= 0 && x <= 13 && y >= 13 && y <= 22 {
		return 1, 0
	}
	if x >= 13 && x <= 24 && y >= 13 && y <= 22 {
		return 1, 1
	}
	if x >= 26 && x <= 35 && y >= 13 && y <= 22 {
		return 1, 2
	}
	if x >= 0 && x <= 13 && y >= 24 && y <= 33 {
		return 2, 0
	}
	if x >= 13 && x <= 24 && y >= 24 && y <= 33 {
		return 2, 1
	}
	if x >= 26 && x <= 35 && y >= 24 && y <= 33 {
		return 2, 2
	}
	return -1, -1
}

func AindaNaoFinalizouOJogo(matriz [3][3]string) bool {
	contadorDeEspacos := 0
	for _, conteudoDaMatriz := range matriz {
		for _, conteudoDoArray := range conteudoDaMatriz {
			if conteudoDoArray == "" {
				contadorDeEspacos++
			}
		}
	}
	return contadorDeEspacos > 0
}
