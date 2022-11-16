package redUsuario

import (
	errores "algogram/errores"
	post "algogram/redUsuario/post"
	usuario "algogram/redUsuario/usuario"
	abb "algogram/tdas/abb"
	hash "algogram/tdas/hash"
	pila "algogram/tdas/pila"
)

type red struct {
	registrados hash.Diccionario[string, usuario.User]
	posteados   hash.Diccionario[int, *post.Post]
	loggeado    pila.Pila[usuario.User]
}

// crear una red
func CrearRed() *red {
	return &red{hash.CrearHash[string, usuario.User](), hash.CrearHash[int, *post.Post](), pila.CrearPilaDinamica[usuario.User]()}
}

// LoggIn permite al usuario loggearse en la red
func (r *red) LoggIn(nombre string) error {
	if !r.loggeado.EstaVacia() {
		return errores.UsuarioYaLoggeado{}
	}
	if !r.registrados.Pertenece(nombre) {
		return errores.UsuarioNoExiste{}
	}
	r.loggeado.Apilar(r.registrados.Obtener(nombre))
	return nil
}

// LoggOut permite al usuario desloggearse de la red
func (r *red) LoggOut() error {
	if r.loggeado.EstaVacia() {
		return errores.UsuarioNoLoggeado{}
	}
	r.loggeado.Desapilar()
	return nil
}

// Loggeado permite ver si hay un usuario loggeado
func (r *red) Loggeado() (usuario.User, error) {
	if r.loggeado.EstaVacia() {
		return nil, errores.UsuarioNoLoggeado{}
	}
	return r.loggeado.VerTope(), nil
}

// PublicarPost permite publicar un post en la red
func (r *red) PublicarPost(post post.Post) error {
	if r.loggeado.EstaVacia() {
		return errores.UsuarioNoLoggeado{}
	}
	r.posteados.Guardar(post.PostID(), &post)

	//guardar el post en el feed de cada usuario
	r.GuardarFeed(post)
	return nil
}

// guardar el post en el feed de cada usuario
func (r *red) GuardarFeed(post post.Post) {
	iterRegistrados := r.registrados.Iterador()
	for iterRegistrados.HaySiguiente() {
		_, usuario := iterRegistrados.VerActual()
		if usuario.NombreUsuario() != post.Publicador() {
			usuario.Feed().Encolar(&post)
		}
		iterRegistrados.Siguiente()
	}
}

// Likear permite al usuario loggeado likear un post
func (r *red) Likear(id int) error {
	if r.loggeado.EstaVacia() || !r.posteados.Pertenece(id) {
		return errores.PostNoExiste{}
	}
	usuario := r.loggeado.VerTope()
	var post post.Post
	post = *r.posteados.Obtener(id)
	arbolDeLikes := post.PostLikes()

	if arbolDeLikes.Pertenece(usuario.NombreUsuario()) {
		return nil
	}
	// la cantidad de likes del post es la cantidad de usuarios en el arbol
	arbolDeLikes.Guardar(usuario.NombreUsuario(), 1)
	return nil
}

// MostrarLikes devuelve el arbol de likes de un post
func (r *red) MostrarLikes(id int) (abb.DiccionarioOrdenado[string, int], error) {
	if !r.posteados.Pertenece(id) {
		return nil, errores.PostNoExiste{}
	}
	var post post.Post
	post = *r.posteados.Obtener(id)
	arbolDeLikes := post.PostLikes()
	return arbolDeLikes, nil
}

// devolver registrados
func (r *red) Registrados() hash.Diccionario[string, usuario.User] { return r.registrados }

// devolver posteados
func (r red) CantidadPost() int { return r.posteados.Cantidad() }
