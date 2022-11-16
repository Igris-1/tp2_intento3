package redUsuario

import (
	post "algogram/redUsuario/post"
	heap "algogram/tdas/heap"
	hash "algogram/tdas/hash"
)

type User interface {
	// ver proximo post
	VerProximoPost() *post.Post

	// UsarioID devuelve el id del usuario
	PosicionUsuario() int

	// Feed devuelve el feed del usuario
	Feed() heap.ColaPrioridad[*post.Post]

	// devolver el nombre
	NombreUsuario() string

	// devuelve el registro de usuarios
	Registro() *hash.Diccionario[string, User]
}
