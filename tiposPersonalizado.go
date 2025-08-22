package main

import "fmt"

func main() {

	//var miCafe Cafe = Cafe{nombre: "Espresso", precio: 19.5, azucar: false, leche: 0}

	var miCafe Cafe = Cafe{"Espresso", 19.5, false, 0}

	fmt.Print(miCafe)
}

type Cafe struct {
	nombre string
	precio float64
	azucar bool
	leche  int
}
