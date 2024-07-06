package main

import (
	"fmt"
)

func main() {
	// Declaracion de Constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de Variables Enteras
	Base := 12
	var altura int = 14
	var area int

	fmt.Println(Base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	// Area cuadrado
	const basecuadrado = 10
	areacuadrado := basecuadrado * basecuadrado
	fmt.Println("Area cuadrado:", areacuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Area Rectangulo

	baseRectangulo := 5
	alturaRectangulo := 4

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El area del Rectangulo es:", areaRectangulo)

	//Circulo : AreaCirculo = pi por radio al cuadrado
	const PI float64 = 3.14 //constant
	var radioCirculo float64 = 10

	AreaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es:", AreaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)
}
