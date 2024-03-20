package requerimientoMatriz

// SumarFilas suma los valores de cada fila de la matriz
func SumarFilas(matriz [][]int) []int {
	sumasFilas := make([]int, len(matriz))

	for i, fila := range matriz {
		for _, valor := range fila {
			sumasFilas[i] += valor
		}
	}

	return sumasFilas
}

// SumarColumnas suma los valores de cada columna de la matriz
func SumarColumnas(matriz [][]int) []int {
	filas, columnas := len(matriz), len(matriz[0])
	sumasColumnas := make([]int, columnas)

	for i := 0; i < columnas; i++ {
		for j := 0; j < filas; j++ {
			sumasColumnas[i] += matriz[j][i]
		}
	}

	return sumasColumnas
}

// SumaTotal suma los valores de las filas y columnas totales
func SumaTotal(sumasFilas, sumasColumnas []int) (int, int) {
	sumaTotalFilas := 0
	sumaTotalColumnas := 0

	for _, suma := range sumasFilas {
		sumaTotalFilas += suma
	}

	for _, suma := range sumasColumnas {
		sumaTotalColumnas += suma
	}

	return sumaTotalFilas, sumaTotalColumnas
}
