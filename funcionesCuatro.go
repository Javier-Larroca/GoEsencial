package main

import "fmt"

func main() {
	alumnos()
	alumnos("Javier", "Priscila", "Nancy")
}

// func alumnos(alumno, edad, direccion string) {
// 	fmt.Println(alumno)
// }

func alumnos(alumno ...string) {
	//fmt.Println(alumno)

	for _, valor := range alumno {
		fmt.Println(valor)
	}

}
