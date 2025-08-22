package main

import "fmt"

func main() {
	const estatico bool = true
	var dinamico bool = true

	//estatico = false
	dinamico = false

	fmt.Print(estatico, dinamico)

}
