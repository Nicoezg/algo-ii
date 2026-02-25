package main

import (
	"algogram/comandos"
	"algogram/diseno/errores"
	algogram "algogram/diseno/usuarios"
	"algogram/impresion"
	"algogram/lectura"
	"bufio"
	"os"
	"strings"
	"tdas/diccionario"
)

const CANT_ARGUMENTOS = 2
const SEPARADOR_ENTRADA = " " //Usado para separar la entrada por consola
const LOGIN = "login"
const LOGOUT = "logout"
const PUBLICAR = "publicar"
const VER_SIGUIENTE_FEED = "ver_siguiente_feed"
const LIKEAR = "likear_post"
const MOSTRAR_LIKES = "mostrar_likes"

func main() {
	if len(os.Args) != CANT_ARGUMENTOS {
		impresion.ImprimirError(errores.ErrorParametros{})
		return
	}
	params := os.Args[1:]
	usuarios, error_lectura := lectura.LecturaUsuarios(params[0])

	if error_lectura != nil {
		impresion.ImprimirError(error_lectura)
		return
	}

	var usuario_act *algogram.Usuario = nil
	publicaciones := diccionario.CrearHash[int, *algogram.Publicacion]()
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		entrada := strings.Split(s.Text(), SEPARADOR_ENTRADA)
		switch entrada[0] {

		case LOGIN:
			entrada = entrada[1:]
			usuario_act = comandos.Ingresar(usuario_act, strings.Join(entrada, " "), usuarios)

		case LOGOUT:
			usuario_act = comandos.Salir(usuario_act)

		case PUBLICAR:
			entrada = entrada[1:]
			comandos.Publicar(usuario_act, strings.Join(entrada, " "), usuarios, &publicaciones)

		case VER_SIGUIENTE_FEED:
			comandos.VerSiguiente(usuario_act)

		case LIKEAR:
			comandos.Likear(entrada[1], usuario_act, publicaciones)

		case MOSTRAR_LIKES:
			comandos.MostrarLikes(entrada[1], publicaciones)
		}
	}
}
