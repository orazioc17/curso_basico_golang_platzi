package main

import (
	pk "curso_golang_basico_platzi/src/mypackage" // Con pk podemos asignarle el nombre q queramos a la importacion
	"fmt"
	"log"
	"strconv"
)

func separacion(nombre string) {
	fmt.Printf("\n\n\n")
	fmt.Printf("##### %s #####\n", nombre)
}

func normalFunction(message string) {
	fmt.Println(message)
}

// Si 2 argumentos son del mismo tipo, la mejor practica es que se declaren juntos y luego el tipo
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func usoFmt() {
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

func operadoresAritmeticos() {
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
}

func variablesConstantesZeroValues() {
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
}

func usoCiclosFor() {
	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n\n")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n\n")

	// for forever
	counterForever := 0
	for {
		counterForever++
		if counterForever >= 100 {
			break
		}
	}
	fmt.Println("Counter Forever")
	fmt.Println(counterForever)
}

func usoIf() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// With or
	if valor1 == 1 || valor2 == 0 {
		fmt.Println("Esta condicion con OR es verdadera")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		/*
		 En Go, nil es el valor cero para punteros, interfaces, mapas, slices, canales y funciones;
		 y corresponde a la representación de un valor no inicializado.
		 Pero ojo, es muy importante no confundir valor “no inicializado” con estado indeterminado,
		 pues nil no es más que otro posible valor válido.
		*/
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)
}

func usoSwitch() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Esto tambien se puede hacer, declarando la variable en la misma linea del switch
	/*
		switch modulo := 4 % 2; modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
		}
	*/

	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	case value < 0:
		fmt.Println("Es menor a cero")
	default:
		fmt.Println("No hay condicion")
	}
}

func usoDeferBreakContinue() {

	// Continue
	fmt.Println("Ejemplo continue")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Break
	fmt.Printf("\nEjemplo break\n")
	counter := 0
	for {
		if counter >= 9 {
			fmt.Println("Saliendo del for forever")
			break
		}
		fmt.Println(counter)
		counter++
	}

	// Defer
	// Defer lo que hace es que la funcion que se use se realice al final, por lo que es util para, por ejemplo
	// terminar una conexion con una base de datos cuando se termine de usar
	fmt.Println()
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
}

func usoArraysSlices() {
	// Array
	fmt.Println("Array")
	var array [4]int
	fmt.Println(array)
	array[0] = 1
	array[1] = 2
	// cap indica la capacidad del array o el slice
	fmt.Println(array, len(array), cap(array))

	// Slice
	fmt.Println("\nSlice")
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice
	// En los corchetes, el primer elemento es inclusivo pero el segundo no, es decir, va hasta antes de ese elemento
	fmt.Println("\nSlicing")
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	fmt.Println("\nAppend")
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nuevo slice
	fmt.Println("\nAppend nuevo slice")
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Estos 3 puntos son para que agregue los elementos de ese slice de forma independiente al antiguo slice
	fmt.Println(slice)
}

func isPalindromo(text string) string {
	var textReversed string
	for i := len(text) - 1; i >= 0; i-- {
		textReversed += string(text[i]) // Aqui se convierte a string porque si no retornaria el codigo ascii
	}

	if text == textReversed {
		return "Es palindromo"
	} else {
		return "No es palindromo"
	}
}

func recorridoArraysSlices() {
	slice := []string{"hola", "que", "hace"}

	for index, value := range slice {
		fmt.Println(index, value)
	}
	// Si solo queremos el indice, al usar solo una variable con range, esta devuelve unicamente el indice

	fmt.Println()
	fmt.Println("ana: ", isPalindromo("ana"))
	fmt.Println("arepera: ", isPalindromo("arepera"))
	fmt.Println("Tienda: ", isPalindromo("Tienda"))

}

func usoMaps() {
	mapa := make(map[string]int)
	fmt.Println("Mapa vacio: ", mapa)

	mapa["Jose"] = 14
	mapa["Pepito"] = 20
	fmt.Println("Mapa con valores: ", mapa)

	for index, value := range mapa {
		fmt.Println(index, value)
	}

	// Encontrar un valor
	value := mapa["Jose"]
	fmt.Println("Edad de Jose: ", value)

	// Saber si un valor existe
	value, ok := mapa["Maria"]                    // Esta llave no existe en el mapa, pero no generara ningun error, sino que retornara un 0, ok retornara true o false dependiendo de si existe o no
	fmt.Println("Valor inexistente: ", value, ok) // "0 false"

}

type car struct {
	brand string
	year  int
}

func usoStruct() {
	// Creando un struct
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var myCar2 car
	fmt.Println(myCar2)
	myCar2.brand = "Ferrari"
	fmt.Println(myCar2)
	myCar2.year = 2022
	fmt.Println(myCar2)
}

func usoModulos() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("Hola Platzi")
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

// Aqui se esta haciendo uso del puntero, pasandole el valor que se encuentra en la direccion de memoria de pc?
func (myPC *pc) duplicateRam() {
	myPC.ram *= 2
}

func usoPunteros() {
	a := 50
	b := &a // ESto seria asignarle a "b" el puntero de "a", es decir, la direccion de memoria donde esta guardado "a"

	fmt.Println(a)
	fmt.Println(b)  // Esto imprimira la direccion de memoria
	fmt.Println(*b) // Y el asterisco accede al valor que esta en esa direccion de memoria

	// Ahora cambiamos el valor al que apunta b
	*b = 100
	fmt.Println("a: ", a) // Y por lo tanto el valor de a cambio tambien

	// Instanciando pc
	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println("myPc: ", myPc)

	myPc.ping()

	fmt.Println("Vieja ram: ", myPc.ram)

	myPc.duplicateRam()

	fmt.Println("Nueva ram: ", myPc.ram)
}

// Practicamente se esta sobreescribiendo la funcion base String que trae por defecto el struct
func (myPC pc) String() string {
	// Recordar que Sprintf no imprime directamente en consola, sino que crea un string con el resultado que de
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func usoStringers() {
	myPC := pc{ram: 16, disk: 100, brand: "msi"}
	fmt.Println(myPC)
}

func main() {
	// Variables, constantes y Zero Values
	fmt.Println("##### Variables, constantes y zero values #####")
	variablesConstantesZeroValues()

	// Operadores Aritmeticos
	separacion("#####Operadores Aritmeticos")
	operadoresAritmeticos()

	// Paquete fmt
	separacion("Paquete fmt")
	usoFmt()

	// Uso de funciones
	separacion("Uso de funciones")
	normalFunction("Esto es un mensaje")
	tripleArgument(3, 5, "Esto es una cadena")

	returnedValue := returnValue(2)
	fmt.Println(returnedValue)

	value1, value2 := doubleReturn(3)
	fmt.Printf("Value 1: %d\nValue2: %d\n", value1, value2)

	// De esta forma se descarte uno de los valores retornados por la funcion
	value1, _ = doubleReturn(10)
	fmt.Printf("Value 1: %d\n", value1)

	// Uso de ciclo for
	separacion("Uso de ciclos for")
	usoCiclosFor()

	// Uso de condicional if
	separacion("Uso de condicional if")
	usoIf()

	// Uso de switch
	separacion("Uso de switch")
	usoSwitch()

	// Uso de defer, break y continue
	separacion("Uso de defer, break y continue")
	usoDeferBreakContinue()

	// Uso de arrays y slices
	separacion("arrays y slices")
	usoArraysSlices()

	// Recorrido de arrays y slices
	separacion("Recorrido de arrays y slices")
	recorridoArraysSlices()

	// Uso de maps
	separacion("Uso de maps")
	usoMaps()

	// Uso de structs (las clases en Go)
	separacion("Uso de structs")
	usoStruct()

	// Uso de modulos
	separacion("Uso de modulos")
	usoModulos()

	// Uso de punteros
	separacion("Uso de punteros")
	usoPunteros()

	// Uso de stringers
	separacion("Uso de stringers")
	usoStringers()
}
