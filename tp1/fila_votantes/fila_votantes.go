package fila_votantes

import (
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"tdas/cola"
)

const PRESIDENTE = "Presidente"
const GOBERNADOR = "Gobernador"
const INTENDENTE = "Intendente"

func Actualizar(fila cola.Cola[votos.Votante]) votos.Votante {
	if fila.EstaVacia() {
		return nil
	}
	return fila.Desencolar()
}

func ObtenerTipoVoto(puesto string) (votos.TipoVoto, error) {
	/// obtenerTipoVoto devuelve el puesto correspondiente al que se esta votando
	switch puesto {
	case PRESIDENTE:
		return votos.PRESIDENTE, nil
	case GOBERNADOR:
		return votos.GOBERNADOR, nil
	case INTENDENTE:
		return votos.INTENDENTE, nil
	default:
		return 0, errores.ErrorTipoVoto{}
	}
}

func Actualizar_fraudulentos(err error, cola_votantes cola.Cola[votos.Votante]) {
	fraudulento := errores.ErrorVotanteFraudulento{Dni: cola_votantes.VerPrimero().LeerDNI()}
	if err == fraudulento {
		Actualizar(cola_votantes)
	}
}
