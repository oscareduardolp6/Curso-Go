package main

import "fmt"

func main() {

	/* Shorhand de declaración y asignación, los ':' es la parte de declaracion, si lo usamos en variables ya declaradas, nos va a causar error */
	dog, cat := "dog", "cat"
	/* Esto sí es valido, porque existe una declaración */
	bird := "robin"
	bird, face := "colibrí", ":D"
	/* Inferencias de datos */
	// var dog, cat = "dog", "cat"
	/* Declaracion y asinacion de 2 variables en una sola línea */
	// var dog, cat string = "dog", "cat"
	/*Declaración y asignacion  en un sola línea */
	// var dog string = "dog"
	/*Declaracion en 2 lineas */
	// var dog string
	// dog = "dog"
	fmt.Println(dog, cat, bird, face)
}
