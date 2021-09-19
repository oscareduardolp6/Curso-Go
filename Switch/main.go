package main

import "fmt"

func main() {
	emoji := "gato"
	switch emoji {
	case "gato":
		fmt.Println("es un gato")
	case "dog":
		fmt.Println("Es un perro ")
	default:
		fmt.Println("no sé que sea")
	}

	/* Agrupar switches */
	switch emoji {
	case "gato", "perro":
		fmt.Println("Es un animal")
	}

	/* Usar comparaciones */
	emoji2 := "perro"
	switch {
	case emoji == "gato" && emoji2 == "perro":

	}
}
