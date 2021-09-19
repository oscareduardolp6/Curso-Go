package main

import "fmt"

func main() {
	/* Los mapas son los diccionarios, tablas de hash, objetos, etc. */
	animals := make(map[string]string) // map(porque eso vamos a crear) [tipo de dato de la llave] tipo de dato del valor
	animals["cat"] = "joe"
	animals["dog"] = "lucho"

	/* Mapas literales */
	fruits := map[string]string{
		"apple":  "emojiManzana",
		"banana": "emojiBanana", // En Go siempre hay que poner la coma al final
	}

	fmt.Println(fruits)

	delete(fruits, "banana") // Esto elimina el elemento del mapa

	/* Obtener los elementos */
	manzana := fruits["apple"]

	/* Si intentamos recuperar el valor de una llave que no existe, nos va a regresar un string vacío */

	/* Al recuperar el valor de un mapa, nos va a devolver 2 cosas, la primera va
	a ser el valor tal cuál y la segunda un bool de si la llave existe o no */
	_, ok := fruits["gorilla"] // El identificado blanco guardaría el valor, que cómo no exite sería un string vacio
	fmt.Println(ok)

	fmt.Println(manzana)
}
