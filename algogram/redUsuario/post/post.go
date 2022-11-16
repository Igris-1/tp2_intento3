package redUsuario

import (
	abb "algogram/tdas/abb"
)

type Post interface {
	// PostID devuelve el id del post
	PostID() int

	// devuelve el nombre del posteador
	Publicador() string

	// devuelve el mensaje del post
	Contenido() string

	// PostLikes devuelve el arbol de likes del post
	PostLikes() abb.DiccionarioOrdenado[string, int]
}
