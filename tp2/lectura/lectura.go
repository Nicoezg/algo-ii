package lectura

import (
	algogram "algogram/diseno/usuarios"
	"bufio"
	"math"
	"os"
	"tdas/cola_prioridad"
	"tdas/diccionario"
)

func LecturaUsuarios(arch string) (diccionario.Diccionario[string, *algogram.Usuario], error) {
	archivo, err := os.Open(arch)
	diccionario := diccionario.CrearHash[string, *algogram.Usuario]()

	if err != nil {
		return diccionario, err
	}

	defer archivo.Close()
	scanner := bufio.NewScanner(archivo)
	contador := 0

	for scanner.Scan() {
		contador++
		name := scanner.Text()
		f := crearFuncionComparacion(contador)

		feed_usuario := cola_prioridad.CrearHeap[*algogram.Publicacion](f)
		usuario := algogram.CrearUsuario(name, &feed_usuario, contador)
		diccionario.Guardar(name, usuario)
	}

	return diccionario, nil
}

func crearFuncionComparacion(valor int) func(a, b *algogram.Publicacion) int {
	f := func(a, b *algogram.Publicacion) int {

		dist := int(math.Abs(float64(valor-b.ObtenerUsuario().ObtenerOrden())) - math.Abs(float64(valor-a.ObtenerUsuario().ObtenerOrden())))
		if dist == 0 {
			return b.ObtenerId() - a.ObtenerId()
		}
		return dist

	}
	return f

}
