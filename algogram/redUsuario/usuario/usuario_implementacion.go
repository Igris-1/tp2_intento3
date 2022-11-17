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

func compareInts(af1, af2, idA, idB int) int {
	if af1 < af2 {
		return 1
	}
	if af1 > af2 {
		return -1
	}
	if idA < idB {
		return 1
	}
	if idA > idB {
		return -1
	}
	return 0
}

func Cmp(u usuario) func(a, b *post.Post) int {
	return func(a, b *post.Post) int {
		registro := u.Registro()
		postA := *a
		postB := *b

		usuarioA := registro.Obtener(postA.Publicador())
		usuarioB := registro.Obtener(postB.Publicador())

		afinidadA := modulo(u.PosicionUsuario() - usuarioA.PosicionUsuario())
		afinidadB := modulo(u.PosicionUsuario() - usuarioB.PosicionUsuario())

		idA := postA.PostID()
		idB := postB.PostID()

		return compareInts(afinidadA, afinidadB, idA, idB)
	}
}

type usuario struct {
	nombre   string
	posicion int
	cmp      func(u usuario) func(a, b *post.Post) int
	feed     heap.ColaPrioridad[*post.Post]
	reg      *hash.Diccionario[string, User]
}

// crear un usuario
func CrearUsuario(nombre string, posicion int, registro *hash.Diccionario[string, User], cmp func(u usuario) func(a *post.Post, b *post.Post) int) User {
	var usuario usuario
	usuario.nombre = nombre
	usuario.posicion = posicion
	usuario.cmp = cmp
	usuario.reg = registro
	usuario.feed = heap.CrearHeap(cmp(usuario))
	return usuario
}

// ver proximo post
func (u usuario) VerProximoPost() *post.Post { return u.feed.Desencolar() }

// Feed devuelve el feed del usuario
func (u usuario) Feed() heap.ColaPrioridad[*post.Post] { return u.feed }

// devolver el nombre
func (u usuario) NombreUsuario() string { return u.nombre }

// devolver el id
func (u usuario) PosicionUsuario() int { return u.posicion }

// devuelve el registro de usuarios
func (u usuario) Registro() hash.Diccionario[string, User] { return *u.reg }
