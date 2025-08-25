package main

import "fmt"

func main() {
	// var arreglo [2]string
	// arreglo[0] = "Manzanas"
	// arreglo[1] = "Bananas"

	arreglo := [2]string{"Peras", "Mandarina"}

	arreglo[0] = "Manzanas"

	fmt.Print(arreglo)
}
