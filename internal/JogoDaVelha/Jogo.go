package JogoDaVelha

func GanharJogo(matriz [3][3]string) string {
	for indexDaMatriz, linhaDaMatriz := range matriz {
		if linhaDaMatriz[0] == linhaDaMatriz[1] && linhaDaMatriz[1] == linhaDaMatriz[2] {
			return linhaDaMatriz[0]
		}
		if matriz[0][indexDaMatriz] == matriz[1][indexDaMatriz] && matriz[0][indexDaMatriz] == matriz[2][indexDaMatriz] {
			return matriz[0][indexDaMatriz]
		}
	}

	if matriz[0][0] == matriz[1][1] && matriz[1][1] == matriz[2][2] {
		return matriz[0][0]
	}

	if matriz[0][2] == matriz[1][1] && matriz[1][1] == matriz[2][0] {
		return matriz[1][1]
	}

	return ""
}
