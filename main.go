package main

import (
	"fmt"
	
	"github.com/IsabelRamirezs/tareaMatriz/matrices"
	"github.com/IsabelRamirezs/tareaMatriz/requerimientoMatriz"
)

func main() {
	// Generar matriz automáticamente
	filas, columnas := 3, 3
	matriz := matrices.GenerarMatrizConsecutiva(filas, columnas)

	// Imprimir la matriz generada
	fmt.Println("Matriz generada:")
	Print(matriz)

	// Sumar los valores de cada fila
	sumasFilas := requerimientoMatriz.SumarFilas(matriz)

	// Sumar los valores de cada columna
	sumasColumnas := requerimientoMatriz.SumarColumnas(matriz)

	// Imprimir los resultados de la suma de cada fila y columna
	fmt.Println("Resultados de la suma de cada fila:")
	for i, suma := range sumasFilas {
		fmt.Printf("Fila %d: %d\n", i+1, suma)
	}

	fmt.Println("Resultados de la suma de cada columna:")
	for j, suma := range sumasColumnas {
		fmt.Printf("Columna %d: %d\n", j+1, suma)
	}

	// Multiplicar las sumas de filas y columnas
	sumaTotalFilas, sumaTotalColumnas := requerimientoMatriz.SumaTotal(sumasFilas, sumasColumnas)

	// Imprimir los resultados totales
	fmt.Println("Resultado total de la suma de las filas:", sumaTotalFilas)
	fmt.Println("Resultado total de la suma de las columnas:", sumaTotalColumnas)

	// Multiplicar las sumas totales
	resultadoFinal := sumaTotalFilas * sumaTotalColumnas

	// Imprimir el resultado final
	fmt.Println("Resultado final de la multiplicación de sumas de filas y columnas:", resultadoFinal)
}

// Imprimir imprime la matriz
func Print(matriz [][]int) {
	for _, fila := range matriz {
		fmt.Println(fila)
	}
}
