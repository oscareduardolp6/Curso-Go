package main

import "fmt"

func main() {
	fruit := "Manzana"
	var p *string
	p = &fruit
	// & indica que debe devolver la dirección de memoria
	fmt.Printf("Tipo %T, Valor: %s, Direccion: %v", fruit, fruit, &fruit)
	/* El operador de desreferenciación '*' nos va a traer el valor de la variable a la que apunta el apuntador,
	o sea, *p nos va a traer "manzana" y con este mismo operador puedo cambiar el valor de la variable */
	fmt.Printf("\n Tipo %T, Valor: %v, Desrefereneciacion: %s", p, p, *p)
}
