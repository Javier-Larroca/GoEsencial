package main

import "fmt"

func main() {

	puntos := 100

	if puntos <= 10 {
		fmt.Print("Puntos incorrectos")
	} else if puntos == 100 {
		fmt.Print("Puntos correctos")
	} else {
		fmt.Print("Tus puntos son ", puntos)
	}

}
