package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println(pi, pi2)
	fmt.Println("pi: ", pi)

	// Declaracion de variables enteras
	base := 12 // Asi se pueden crear y asignar directamente variables sin tener que usar mas palabras reservadas
	fmt.Println(base)
	var altura int = 14 // Y asi se inicializa con "var"
	var area int // Y esta es la ultima forma, creando la variable sin inicializarla
	area = 3
	fmt.Println(altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("\nZero values")
	fmt.Println(a,b,c,d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)
}
