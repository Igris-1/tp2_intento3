package redUsuario

import (
	post "algogram/redUsuario/post"
	hash "algogram/tdas/hash"
	heap "algogram/tdas/heap"
)

// funcion auxiliar para devolver el modulo de un numero
func modulo(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// funcion de comparacion de post respecto a afinidades
func (u usuario) comparePosts(a, b *post.Post) int {
	registro := u.Registro()
	


	afinidadA := modulo((u.PosicionUsuario() - a.publicador.UsuarioID()))
	afinidadB := modulo((u.PosicionUsuario() - b.publicador.UsuarioID()))
	if afinidadA < afinidadB {
		return 1
	}
	if afinidadA > afinidadB {
		return 1
	}
	if a.id < b.id {
		return 1
	}
	if a.id > b.id {
		return -1
	}
	return 0
}

type usuario struct {
	nombre   string
	posicion int
	feed     heap.ColaPrioridad[*post.Post]
	reg      *hash.Diccionario[string, User]
}

// crear un usuario
func CrearUsuario(nombre string, posicion int, registro *hash.Diccionario[string, User]) User {
	var usuario usuario
	usuario.nombre = nombre
	usuario.posicion = posicion
	usuario.feed = heap.CrearHeap(usuario.comparePosts)
	usuario.reg = registro
	return usuario
}

// ver proximo post
func (u usuario) VerProximoPost() *post.Post {
	return u.feed.Desencolar()
}

// Feed devuelve el feed del usuario
func (u usuario) Feed() heap.ColaPrioridad[*post.Post] {
	return u.feed
}

// devolver el nombre
func (u usuario) NombreUsuario() string {
	return u.nombre
}

// devolver el id
func (u usuario) PosicionUsuario() int {
	return u.posicion
}

// devuelve el registro de usuarios
func (u usuario) Registro() *hash.Diccionario[string, User] {
	return u.reg
}
