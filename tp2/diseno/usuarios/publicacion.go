package algogram

import "tdas/diccionario"

type Publicacion struct {
	id        int
	likeado   diccionario.DiccionarioOrdenado[string, bool]
	usuario   *Usuario
	contenido string
}

func (publicacion *Publicacion) ObtenerId() int {
	return publicacion.id
}

func (publicacion *Publicacion) ObtenerLikes() int {
	return publicacion.ObtenerLikeados().Cantidad()
}

func (publicacion *Publicacion) ObtenerLikeados() diccionario.DiccionarioOrdenado[string, bool] {
	likeado := publicacion.likeado
	return likeado
}

func (publicacion *Publicacion) ObtenerContenido() string {
	return publicacion.contenido
}

func (publicacion *Publicacion) ObtenerUsuario() *Usuario {
	return publicacion.usuario
}
