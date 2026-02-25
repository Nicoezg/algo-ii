package impresion

import (
	algogram "algogram/diseno/usuarios"
	"fmt"
	"os"
	"tdas/diccionario"
)

func ImprimirError(err error) {
	fmt.Fprintln(os.Stdout, err.Error())
}

func Bienvenida(usuario string) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("Hola %s", usuario))
}

func Despedida() {
	fmt.Fprintln(os.Stdout, "Adios")
}

func PostPublicado() {
	fmt.Fprintln(os.Stdout, "Post publicado")
}

func ImprimirPost(publicacion algogram.Publicacion) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("Post ID %d", publicacion.ObtenerId()))
	fmt.Fprintln(os.Stdout, fmt.Sprintf("%s dijo: %s", publicacion.ObtenerUsuario().ObtenerNombre(), publicacion.ObtenerContenido()))
	fmt.Fprintln(os.Stdout, fmt.Sprintf("Likes: %d", publicacion.ObtenerLikes()))
}

func ImprimirLikes(likeados diccionario.DiccionarioOrdenado[string, bool]) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("El post tiene %d likes:", likeados.Cantidad()))

	likeados.Iterar(func(clave string, valor bool) bool {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("	%s", clave))
		return true
	})
}

func ImprimirLike() {
	fmt.Fprintln(os.Stdout, "Post likeado")
}
