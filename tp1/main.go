package main

import (
	"bufio"
	"os"
	"rerepolez/archivos"
	"rerepolez/comandos"
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"rerepolez/impresion"
	"rerepolez/padron_electoral"
	"strings"
	"tdas/cola"
)

const CANT_POSTULANTES = 3
const CANT_ARGUMENTOS = 3
const INGRESAR = "ingresar"
const VOTAR = "votar"
const FIN_VOTO = "fin-votar"
const DESHACER = "deshacer"
const SEPARADOR_ENTRADA = " " //Es usada para separar la entrada recibida por consola
func main() {
	if len(os.Args) != CANT_ARGUMENTOS {
		impresion.Error(errores.ErrorParametros{})
		return
	}
	params := os.Args[1:]
	partidos, error_partidos := archivos.Lectura_partidos(params[0])
	padron, error_lectura := archivos.Lectura_padron(params[1])
	if error_lectura != nil || error_partidos != nil {
		impresion.Error(archivos.Error_lectura(error_lectura, error_partidos))
		return
	}
	padron = padron_electoral.Ordenar(padron)
	impugnados := 0
	s := bufio.NewScanner(os.Stdin)
	cola_votantes := cola.CrearColaEnlazada[votos.Votante]()
	for s.Scan() {
		entrada := strings.Split(s.Text(), SEPARADOR_ENTRADA)
		switch entrada[0] {
		case INGRESAR:
			comandos.Ingresar(entrada, padron, cola_votantes)
		case VOTAR:
			comandos.Votar(entrada, cola_votantes, partidos)
		case FIN_VOTO:
			comandos.Fin_votar(cola_votantes, partidos, &impugnados)
		case DESHACER:
			comandos.Deshacer(cola_votantes)
		}
	}
	impresion.Resultado_elecciones(cola_votantes, partidos, impugnados)
}
