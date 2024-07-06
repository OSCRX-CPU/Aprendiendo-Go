package main

import "fmt"

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
}
