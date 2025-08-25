package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Pais struct {
	NOMBRE    string
	POBLACION float32
}

type Paises struct {
	Paises []Pais `json:"paises"`
}

func main() {
	data, err := os.Open("list.json")
	if err != nil {
		fmt.Print(err)
	}

	byteValue, _ := ioutil.ReadAll(data)

	var paises Paises

	err = json.Unmarshal(byteValue, &paises)

	if err != nil {
		fmt.Print("Error: ", err)
	}

	for i := 0; i < len(paises.Paises); i++ {
		fmt.Println("* " + paises.Paises[i].NOMBRE)
	}
}
