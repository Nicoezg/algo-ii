package impresion

import (
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"tdas/cola"
)

const PRESIDENTE = "Presidente"
const GOBERNADOR = "Gobernador"
const INTENDENTE = "Intendente"
const OK = "OK"

var ELECCIONES = []string{PRESIDENTE, GOBERNADOR, INTENDENTE}

func Resultado_elecciones(cola_votantes cola.Cola[votos.Votante], partidos []votos.Partido, impugnados int) {
	if !cola_votantes.EstaVacia() {
		Error(errores.ErrorCiudadanosSinVotar{})
	}
	for i, cargo := range ELECCIONES {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%s:", cargo))
		for _, partido := range partidos {
			fmt.Fprintln(os.Stdout, partido.ObtenerResultado(votos.TipoVoto(i)))
		}
		fmt.Fprintln(os.Stdout, "")
	}
	fmt.Fprintln(os.Stdout, fmt.Sprintf("Votos Impugnados: %d %s", impugnados, votos.VerificarSingularOPlural(impugnados)))
}
func Error(err error) {
	fmt.Fprintln(os.Stdout, err.Error())
}

func Ok() {
	fmt.Fprintln(os.Stdout, OK)
}
