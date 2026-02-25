package archivos

import (
	"bufio"
	"os"
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"strconv"
	"strings"
)

const CANT_POSTULANTES = 3
const CANT_ARGUMENTOS = 3
const SEPERADOR_PARTIDOS = "," //Es usada para separar el nombre del partido y los partidos del archivo que contiene los partidos (formato csv)

func Lectura_partidos(arch string) ([]votos.Partido, error) {
	blanco := votos.CrearVotosEnBlanco()
	partidos := []votos.Partido{blanco}
	archivo, err := os.Open(arch)
	if err != nil {
		return partidos, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		partido_arr := strings.Split(scanner.Text(), SEPERADOR_PARTIDOS)
		if len(partido_arr) != CANT_POSTULANTES+1 {
			return partidos, errores.ErrorLeerArchivo{}
		}
		candidatos := [CANT_POSTULANTES]string{partido_arr[1], partido_arr[2], partido_arr[3]}
		partidos = append(partidos, votos.CrearPartido(partido_arr[0], candidatos))
	}
	return partidos, nil

}

func Lectura_padron(arch string) ([]votos.Votante, error) {
	padron := []votos.Votante{}
	archivo, err := os.Open(arch)
	if err != nil {
		return padron, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return padron, errores.ErrorLeerArchivo{}
		}
		padron = append(padron, votos.CrearVotante(num))
	}
	return padron, nil
}

func Error_lectura(err1, err2 error) error {
	if err1 != nil {
		return err1
	}
	return err2
}
