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
	var area int        // Y esta es la ultima forma, creando la variable sin inicializarla
	area = 3
	fmt.Println(altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("\nZero values")
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es: ", areaCuadrado)

	var x float64 = 10
	var y float64 = 50

	// Suma
	var result float64 = x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = x - y
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	// Division
	result = x / y
	fmt.Println("Division: ", result)

	// Modulo/Residuo
	result = 50 % 10
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("##### Paquete fmt #####")

	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)

	// Uso de printf
	nombre := "Platzi"
	cursos := 500

	// La buena practica es agregar el tipo de dato correspondiente, pero si no se tiene cereza del tipo de dato que se 
	// va a imprimir, se puede usar $v
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Conocer el tipo de dato de una variable
	fmt.Printf("helloMessahe: %T\n", helloMessage)
	fmt.Printf("Cursos: %T\n", cursos)
	
}
