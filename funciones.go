package main

import "fmt"

func main() {
	fmt.Print(multiplicar(10, 20))
}

func multiplicar(numero1, numero2 int) int {
	return numero1 * numero2
}
