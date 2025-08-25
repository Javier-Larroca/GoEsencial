package main

import (
	"fmt"
	"time"
)

func main() {

	tiempo := time.Now()

	switch diaDeHoy := tiempo.Weekday(); diaDeHoy {
	case 0:
		fmt.Println("Domingo")
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	default:
		fmt.Println("Error, no más días en la semana.")
	}

}
