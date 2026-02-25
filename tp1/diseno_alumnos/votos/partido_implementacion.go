package votos

import (
	"fmt"
)

const CANT_POSTULANTES = 3

type partidoImplementacion struct {
	nombre     string
	candidatos [CANT_POSTULANTES]string
	votos      [CANT_POSTULANTES]int
}

type partidoEnBlanco struct {
	votos [CANT_POSTULANTES]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := &partidoImplementacion{nombre, candidatos, [CANT_VOTACION]int{}}
	return partido
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	votos := partido.votos[tipo]
	candidato := partido.candidatos[tipo]
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, candidato, votos, VerificarSingularOPlural(votos))
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	votos := blanco.votos[tipo]
	return fmt.Sprintf("Votos en Blanco: %d %s", votos, VerificarSingularOPlural(votos))
}

func VerificarSingularOPlural(votos int) string {
	if votos == 1 {
		return "voto"
	}
	return "votos"
}
