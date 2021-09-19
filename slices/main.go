package main

import "fmt"

func main() {
	//Slice son apuntadores a arrays
	set := [4]string{"Perro", "Gato", "pajaro", "hamster"}
	//Lo que se pasa dentro de los parentesis es el index inicial y el final, pero el final es excluyente (no se incluye)

	// Al ser apuntadores a arrays, si se hace un cambio en un array, se va a cambiar también en el array al que apunta
	animales := set[0:3] // animales := set[:3]  Son iguales
	fly := set[2:4]      // fly := set[2:] Son iguales
	fmt.Println(animales, fly)

	/* 	El len de un slice es el número de elementos que tiene el slice
	La cap de un slice es el número de elementos del array del slice DESDE el indice que abarca el slice */

	/* Para agregar elementos a un slice, append va a utilizar la capacidad para guardar el núevo elemento,
	hay que tenerlo en cuenta a la hora de tomar solo partes parciales de un arreglo  .
	Cuando ya nos pasamos del tamaño original del array, Go crea un nuevo array, por lo tanto se pierde la referencia
	al array origina, se duplican los valores, pero la referencia cambia */
	fly = append(fly, "Nueva Fruta")

	/* Crear Slice desde 0 */
	// fruit := []string{"Fresa", "Mango"} // Es igual que crear un array, pero sin el tamaño
	/* Esto también crea un slice */
	fruits := make([]string, 0, 0) //Tipo, tamaño y capacidad
	fmt.Println(fruits)

	//Los slices son inicializados en nil ... a menos que lo hagamos literal (cuando lo inicializamos con {}) ahí solo es un slice vacío

}
