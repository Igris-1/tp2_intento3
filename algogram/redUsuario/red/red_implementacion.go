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

// CrearRed: crea una red con un registro de usuarios, un registro de posts y un usuario loggeado
func CrearRed() *red {
	return &red{hash.CrearHash[string, usuario.User](), hash.CrearHash[int, *post.Post](), pila.CrearPilaDinamica[usuario.User]()}
}

// LoggIn: loggea al usuario en la red, caso contrario devuelve error
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

// LoggOut: desloggea al usuario de la red, caso contrario devuelve error
func (r *red) LoggOut() error {
	if r.loggeado.EstaVacia() {
		return errores.UsuarioNoLoggeado{}
	}
	r.loggeado.Desapilar()
	return nil
}

// Loggeado: devuelve el usuario loggeado, caso contrario devuelve error
func (r *red) Loggeado() (usuario.User, error) {
	if r.loggeado.EstaVacia() {
		return nil, errores.UsuarioNoLoggeado{}
	}
	return r.loggeado.VerTope(), nil
}

// PublicarPost: publicar un post en la red lo guarda en el registro de posts y en el feed de cada usuario
func (r *red) PublicarPost(post post.Post) error {
	r.posteados.Guardar(post.PostID(), &post)
	r.guardarFeed(post)
	return nil
}

// guardarFeed: funcion auxiliar de PublicarPost() que guarda el post en el feed de cada usuario
func (r *red) guardarFeed(post post.Post) {
	iterRegistrados := r.registrados.Iterador()
	for iterRegistrados.HaySiguiente() {
		_, usuario := iterRegistrados.VerActual()
		if usuario.NombreUsuario() != post.Publicador() {
			feedDelUser := usuario.Feed()
			feedDelUser.Encolar(&post)
		}
		iterRegistrados.Siguiente()
	}
}

// Likear: permite al usuario loggeado likear un post, caso contrario devuelve error
func (r *red) Likear(id int) error {
	if r.loggeado.EstaVacia() || !r.posteados.Pertenece(id) {
		return errores.PostNoExiste{}
	}
	usuario := r.loggeado.VerTope()
	post := *r.posteados.Obtener(id)
	arbolDeLikes := post.PostLikes()

	if arbolDeLikes.Pertenece(usuario.NombreUsuario()) {
		return nil
	}
	arbolDeLikes.Guardar(usuario.NombreUsuario(), 1)
	return nil
}

// MostrarLikes: devuelve el diccionario de personas que le dieron like al post, caso contrario devuelve error
func (r *red) MostrarLikes(id int) (abb.DiccionarioOrdenado[string, int], error) {
	if !r.posteados.Pertenece(id) {
		return nil, errores.PostInexistenteOSinLikes{}
	}
	post := *r.posteados.Obtener(id)
	arbolDeLikes := post.PostLikes()
	if arbolDeLikes.Cantidad() == 0 {
		return nil, errores.PostInexistenteOSinLikes{}
	}
	return arbolDeLikes, nil
}

// Registrados: devuelve el diccionario de usuarios registrados en la red
func (r *red) Registrados() hash.Diccionario[string, usuario.User] { return r.registrados }

// CantidadPost: devuelve la cantidad de posts publicados en la red
func (r red) CantidadPost() int { return r.posteados.Cantidad() }
