package redUsuario

import (
	abb "algogram/tdas/abb"
)

type post struct {
	id         int
	publicador string
	contenido  string
	personas   abb.DiccionarioOrdenado[string, int]
}

func compareString(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

// crear un post
func CrearPost(id int, publicador string, contenido string) Post {
	return &post{id, publicador, contenido, abb.CrearABB[string, int](compareString)}
}

// PostID devuelve el id del post
func (p post) PostID() int { return p.id }

// PostLikes devuelve el arbol de likes
func (p post) PostLikes() abb.DiccionarioOrdenado[string, int] { return p.personas }

// devolver el nombre del post
func (p post) Publicador() string { return p.publicador }

// devolver el contenido del post
func (p post) Contenido() string { return p.contenido }