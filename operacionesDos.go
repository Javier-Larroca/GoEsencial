package main

import "fmt"

func main() {
	var frutas = []string{"Manzanas", "Uvas", "Bananas"}

	//var frutas2 []string
	//frutas2 = append(frutas, "Peras")
	//fmt.Println(frutas2)

	frutas = append(frutas, "Peras")

	//fmt.Println(frutas[0])
	fmt.Println(frutas[0:3])
	fmt.Println(len(frutas))
}
