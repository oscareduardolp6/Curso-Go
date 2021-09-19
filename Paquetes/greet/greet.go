package greet

// Las variables a nivel de paquete no pueden usar el shortcut
// Si las cosa incian con mayusculas se exportan, si no, no
var emoji = "Amigo"

func English() string {
	return "Hi" + emoji
}

func Italian() string {
	return "Ciao" + emoji
}
