package votos

import (
	"rerepolez/diseno_alumnos/errores"
	"tdas/pila"
)

type votanteImplementacion struct {
	dni        int
	habilitado bool
	pila_Votos pila.Pila[Voto]
	voto       Voto
}

func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni: dni, habilitado: true, pila_Votos: pila.CrearPilaDinamica[Voto]()}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	votante.pila_Votos.Apilar(votante.voto)
	if !votante.habilitado {
		return errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	if alternativa == 0 {
		votante.voto.Impugnado = true
	}
	votante.voto.VotoPorTipo[tipo] = alternativa
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if !votante.habilitado {
		return errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	if !votante.pila_Votos.EstaVacia() {
		votante.voto = votante.pila_Votos.Desapilar()
		return nil
	}
	return errores.ErrorNoHayVotosAnteriores{}
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if !votante.habilitado {
		return votante.voto, errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}
	votante.habilitado = false
	return votante.voto, nil
}
