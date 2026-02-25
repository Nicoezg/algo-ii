package algogram

import (
	"algogram/diseno/errores"
	"strings"
	"tdas/cola_prioridad"
	"tdas/diccionario"
)

type Usuario struct {
	nombre string
	feed   cola_prioridad.ColaPrioridad[*Publicacion]
	orden  int
}

func CrearUsuario(nombre string, feed *cola_prioridad.ColaPrioridad[*Publicacion], contador int) *Usuario {
	return &Usuario{nombre: nombre, feed: *feed, orden: contador}
}

func (usuario *Usuario) Publicar(contenido string, cantidad int) *Publicacion {
	abb := diccionario.CrearABB[string, bool](strings.Compare)
	return &Publicacion{id: cantidad, likeado: abb, usuario: usuario, contenido: contenido}
}

func (usuario *Usuario) VerProximoPost() (error, *Publicacion) {
	feed := usuario.feed
	if feed.Cantidad() == 0 {
		return errores.ErrorPost{}, &Publicacion{}
	}
	return nil, feed.Desencolar()
}

func (usuario *Usuario) Likear(publicacion *Publicacion) {
	likeados := publicacion.ObtenerLikeados()
	likeados.Guardar(usuario.nombre, true)
}

func (usuario Usuario) ObtenerNombre() string {
	return usuario.nombre
}

func (usuario Usuario) ObtenerOrden() int {
	return usuario.orden
}

func (usuario Usuario) ActualizarFeed(publicacion *Publicacion) {
	feed := usuario.feed
	feed.Encolar(publicacion)
}
