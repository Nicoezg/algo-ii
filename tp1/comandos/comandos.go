package comandos

import (
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"rerepolez/fila_votantes"
	"rerepolez/impresion"
	"rerepolez/padron_electoral"
	"strconv"
	"tdas/cola"
)

const PRESIDENTE = "Presidente"
const GOBERNADOR = "Gobernador"
const INTENDENTE = "Intendente"

func Ingresar(entrada []string, padron []votos.Votante, cola_votantes cola.Cola[votos.Votante]) {
	if len(entrada) == 1 {
		return
	}
	DNI, err := strconv.Atoi(entrada[1])
	if padron_electoral.Checkear_dni(DNI, err) {
		posicion, err := padron_electoral.Busqueda_DNI(DNI, padron, 0, len(padron)-1)
		if err != nil {
			impresion.Error(err)
			return
		}
		cola_votantes.Encolar(padron[posicion])
		impresion.Ok()
	}
}

func Votar(entrada []string, cola_votantes cola.Cola[votos.Votante], partidos []votos.Partido) {
	if len(entrada) == 1 {
		return
	}
	if cola_votantes.EstaVacia() {
		impresion.Error(errores.FilaVacia{})
		return
	}
	tipo, err := fila_votantes.ObtenerTipoVoto(entrada[1])
	if err != nil {
		impresion.Error(err)
		return
	}
	alternativa, err := strconv.Atoi(entrada[2])
	if err != nil || alternativa >= len(partidos) || alternativa < 0 {
		impresion.Error(errores.ErrorAlternativaInvalida{})
		return
	}
	err = cola_votantes.VerPrimero().Votar(tipo, alternativa)
	if err != nil {
		impresion.Error(err)
		fila_votantes.Actualizar_fraudulentos(err, cola_votantes)
		return
	}
	impresion.Ok()
}

func Fin_votar(cola_votantes cola.Cola[votos.Votante], partidos []votos.Partido, impugnados *int) {
	/// Fin_votar recuenta el voto y lo guarda. En caso de que no haya votante, lanza el error correspondiente.
	/// En caso de que el votante sea fraudulento, lo descarta de la fila, lanza el error correspondiente y pasa el siguiente votante.
	if cola_votantes.EstaVacia() {
		impresion.Error(errores.FilaVacia{})
		return
	}
	voto, err := cola_votantes.VerPrimero().FinVoto()
	if err != nil {
		impresion.Error(err)
		fila_votantes.Actualizar(cola_votantes)
		return
	}
	if !voto.Impugnado {
		for i := 0; i < len(voto.VotoPorTipo); i++ {
			partidos[voto.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
		}
	} else {
		*impugnados++
	}
	impresion.Ok()
	fila_votantes.Actualizar(cola_votantes)
}

func Deshacer(cola_votantes cola.Cola[votos.Votante]) {
	/// Deshacer permite al votante deshacer el último voto. En caso de que no haya un votante, lanza el error correspondiente
	/// En caso de que el votante sea fraudulento, lanza el error correspondiente, elimina al votante de la fila y pasa el proximo votante.
	/// Si no hay votos para deshacer, lo indica mediante un error.
	if cola_votantes.EstaVacia() {
		impresion.Error(errores.FilaVacia{})
		return
	}
	err := cola_votantes.VerPrimero().Deshacer()
	if err != nil {
		impresion.Error(err)
		fila_votantes.Actualizar_fraudulentos(err, cola_votantes)
		return
	}
	impresion.Ok()
}
