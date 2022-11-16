package redUsuario

import (
	post "algogram/redUsuario/post"
	user "algogram/redUsuario/usuario"
	abb "algogram/tdas/abb"
	hash "algogram/tdas/hash"
)

type Red interface {
	// LoggIn permite al usuario loggearse en la red
	LoggIn(nombre string) error

	// LoggOut permite al usuario desloggearse de la red
	LoggOut() error

	// Loggeado permite ver si hay un usuario loggeado
	Loggeado() (user.User, error)

	// MostrarLikes devuelve el arbol de likes de un determinado post
	MostrarLikes(id int) (abb.DiccionarioOrdenado[string, int], error)

	// PublicarPost permite publicar un post en la red
	PublicarPost(post post.Post) error

	// Likear permite al usuario loggeado likear un post
	Likear(id int) error

	// GuardarFeed guarda el post en el heap de cada usuario
	GuardarFeed(post post.Post)

	// devolver hash de registrados
	Registrados() hash.Diccionario[string, user.User]

	// devolver la cantidad de post
	CantidadPost() int
}
