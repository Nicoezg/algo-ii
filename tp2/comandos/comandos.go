package comandos

import (
	"algogram/diseno/errores"
	algogram "algogram/diseno/usuarios"
	"algogram/impresion"
	"strconv"
	"tdas/diccionario"
)

func Ingresar(usuario_act *algogram.Usuario, nombre string, cuentas diccionario.Diccionario[string, *algogram.Usuario]) *algogram.Usuario {
	if usuario_act != nil {
		impresion.ImprimirError(errores.ErrorUsuarioLoggeado{})
		return usuario_act
	}

	if !cuentas.Pertenece(nombre) {
		impresion.ImprimirError(errores.ErrorUsuarioInexistente{})
		return nil
	}

	usuario_act = cuentas.Obtener(nombre)
	impresion.Bienvenida(nombre)
	return usuario_act

}

func Salir(usuario_act *algogram.Usuario) *algogram.Usuario {
	if !verificarUsuarioLoggeado(usuario_act) {
		impresion.ImprimirError(errores.ErrorUsuarioNoLoggeado{})
		return nil
	}

	impresion.Despedida()
	return nil
}

func Publicar(usuario_act *algogram.Usuario, contenido string, usuarios diccionario.Diccionario[string, *algogram.Usuario], publicaciones *diccionario.Diccionario[int, *algogram.Publicacion]) {
	if !verificarUsuarioLoggeado(usuario_act) {
		impresion.ImprimirError(errores.ErrorUsuarioNoLoggeado{})
		return
	}

	publicacion := usuario_act.Publicar(contenido, (*publicaciones).Cantidad())
	iter := usuarios.Iterador()
	for iter.HaySiguiente() {
		_, usuario := iter.VerActual()

		if usuario.ObtenerNombre() == publicacion.ObtenerUsuario().ObtenerNombre() {
			iter.Siguiente()
			continue
		}
		usuario.ActualizarFeed(publicacion)
		iter.Siguiente()
	}
	(*publicaciones).Guardar(publicacion.ObtenerId(), publicacion)
	impresion.PostPublicado()

}

func VerSiguiente(usuario_act *algogram.Usuario) {
	if !verificarUsuarioLoggeado(usuario_act) {
		impresion.ImprimirError(errores.ErrorPost{})
		return
	}
	err, publicacion := usuario_act.VerProximoPost()

	if err != nil {
		impresion.ImprimirError(err)
		return
	}

	impresion.ImprimirPost(*publicacion)
}

func Likear(id string, usuario *algogram.Usuario, diccionario diccionario.Diccionario[int, *algogram.Publicacion]) {
	if usuario == nil {
		impresion.ImprimirError(errores.ErrorPostInexistenteSinLikes{})
		return
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		impresion.ImprimirError(errores.ErrorParametros{})
		return
	}

	if !diccionario.Pertenece(id_int) {
		impresion.ImprimirError(errores.ErrorPostInexistenteSinLikes{})
		return
	}

	pub := diccionario.Obtener(id_int)
	usuario.Likear(pub)
	impresion.ImprimirLike()
}

func MostrarLikes(id string, pub_total diccionario.Diccionario[int, *algogram.Publicacion]) {
	id_int, err := strconv.Atoi(id)
	if err != nil {
		impresion.ImprimirError(errores.ErrorParametros{})
		return
	}

	if !pub_total.Pertenece(id_int) || pub_total.Cantidad() == 0 {
		impresion.ImprimirError(errores.ErrorMostrarLikes{})
		return
	}

	pub := pub_total.Obtener(id_int)

	if pub.ObtenerLikes() == 0 {
		impresion.ImprimirError(errores.ErrorMostrarLikes{})
		return
	}
	likeados := pub.ObtenerLikeados()
	impresion.ImprimirLikes(likeados)
}

func ValidarEntrada(entrada []string) bool {
	if len(entrada) == 1 {
		impresion.ImprimirError(errores.ErrorEntrada{})
		return false
	}

	return true
}

func verificarUsuarioLoggeado(usuario_act *algogram.Usuario) bool {
	return usuario_act != nil
}
