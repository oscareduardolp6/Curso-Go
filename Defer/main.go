package main

import (
	"fmt"
	"os"
)

func main() {
	/* Aplaza la ejecución para cuándo la función regresa, las funciones en defer trabaja
	en LIFO, los parametros con el defer se pasan al momento qué debería ejecutarse */
	// defer fmt.Println(3)
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(1)
	// fmt.Println(2)
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola Oscar, probando esccribir en el archivo"))
	if err != nil {
		// file.Close()
		fmt.Printf("Error %v", err)
		return
	}
	// file.Close()
}
