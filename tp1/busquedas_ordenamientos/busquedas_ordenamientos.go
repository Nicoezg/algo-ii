package dni

import (
	"rerepolez/diseno_alumnos/errores"
	"rerepolez/diseno_alumnos/votos"
	"rerepolez/impresion"
)

func Ordenar(votantes []votos.Votante) []votos.Votante {
	if len(votantes) < 2 {
		return votantes
	}
	izq := Ordenar(votantes[:len(votantes)/2])
	der := Ordenar(votantes[len(votantes)/2:])
	return merge(izq, der)
}

func merge(a []votos.Votante, b []votos.Votante) []votos.Votante {
	res := []votos.Votante{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].LeerDNI() < b[j].LeerDNI() {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return res
}

func Busqueda_DNI(dni int, padron []votos.Votante, ini, fin int) (int, error) {
	if ini > fin {
		return -1, errores.DNIFueraPadron{}
	}
	medio := (ini + fin) / 2
	if padron[medio].LeerDNI() == dni {
		return medio, nil
	}
	if padron[medio].LeerDNI() < dni {
		return Busqueda_DNI(dni, padron, medio+1, fin)
	}
	return Busqueda_DNI(dni, padron, ini, medio-1)
}

func Checkear_dni(dni int, err error) bool {
	if dni < 0 || err != nil {
		impresion.Error(errores.DNIError{})
		return false
	}
	return true
}
