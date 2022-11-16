package main

import (
	comandos "algogram/comandos"
	errores "algogram/errores"
	funciones "algogram/funciones"
	"bufio"
	"os"
)

func main() {
	ruta := os.Args[1]

	// agregar usuarios a la red
	red, error := funciones.AgregarUsuarios(ruta)
	if error != nil {
		panic(errores.ErrorLecturaArchivo{})
	}

	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		comando, resto := funciones.LeerComando(entrada.Text())

		switch comando {
		case "login":
			comandos.LoggIn(red, resto)
		case "logout":
			comandos.LoggOut(red)
		case "publicar":
			comandos.Publicar(red, resto)
		case "ver_siguiente_feed":
			comandos.VerSiguienteFeed(red)
		case "likear_post":
			comandos.Likear(red, resto)
		case "mostrar_likes":
			comandos.MostrarLikes(red, resto)
		}
	}
}
