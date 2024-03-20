package matrices

// GenerarMatrizConsecutiva genera una matriz con números consecutivos
func GenerarMatrizConsecutiva(filas, columnas int) [][]int {
	matriz := make([][]int, filas)

	for i := 0; i < filas; i++ {
		matriz[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz[i][j] = i*columnas + j + 1 // Números consecutivos empezando desde 1
		}
	}

	return matriz
}
