package comandos

import (
	errores "algogram/errores"
	post "algogram/redUsuario/post"
	red "algogram/redUsuario/red"
	"fmt"
	"strconv"
)

// LoggIn
func LoggIn(red red.Red, nombre string) {
	err := red.LoggIn(nombre)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola", nombre)
	}
}

// LoggOut
func LoggOut(red red.Red) {
	err := red.LoggOut()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adios")
	}
}

// Publicar
func Publicar(red red.Red, contenido string) {
	loggeado, err := red.Loggeado()
	if err != nil {
		fmt.Println(err)
	} else {
		cantPost := red.CantidadPost()
		post := post.CrearPost(cantPost, loggeado.NombreUsuario(), contenido)
		red.PublicarPost(post)
		fmt.Println("Post publicado")
	}
}

// VerSiguienteFeed
func VerSiguienteFeed(red red.Red) {
	usuario, err := red.Loggeado()
	if err != nil || usuario.Feed().EstaVacia() {
		fmt.Println(errores.SinPostsOsinLoggeado{})
	} else {
		post := *usuario.Feed().Desencolar()
		fmt.Println("Post ID", post.PostID())
		fmt.Println(post.Publicador(), "dijo:", post.Contenido())
		fmt.Println("Likes:", post.PostLikes().Cantidad())
	}
}

// likear un post
func Likear(red red.Red, id string) {
	//transformar a int
	idInt, _ := strconv.Atoi(id)
	err := red.Likear(idInt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Post likeado")
	}
}

// mostrar likes
func MostrarLikes(red red.Red, id string) {
	// transformar a int
	idInt, _ := strconv.Atoi(id)
	arbol, err := red.MostrarLikes(idInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El post tiene", arbol.Cantidad(), "likes:")
	arbol.Iterar(func(key string, value int) bool {
		fmt.Println("	" + key)
		return true
	})
}
