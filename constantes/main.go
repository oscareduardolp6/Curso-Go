package main

func main() {
	const pi = 3.14
	/* En Go igual que en java los ' son para chars y las "" para Strings  */

	/*CAsting */
	// var a uint8 = 100
	// var b uint16 = 23000
	// _ := uint16(a) + b

	/*Identificador blanco, variables que no hemos declarado, pero que queremos mantener */
	_ = 234
	var _ string = "test"

	/* Go inicializa las variables "", 0, 0.0 etc. */

	/* En Go se ++ y -- son declaraciones, tienen que ir solas */

}
