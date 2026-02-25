package diccionario

import (
	"tdas/diccionario/fnv"
	"tdas/lista"
)

const NO_PERTENECE = "La clave no pertenece al diccionario"
const TERMINO_ITERAR = "El iterador termino de iterar"
const CAPACIDAD_INICIAL = 73
const CRITERIO_REDIMENSION_AGRANDAR = 2
const CRITERIO_REDIMENSION_ACHICAR = 4
const FACTOR_REDIMENSION_AGRANDAR = 2
const FACTOR_REDIMENSION_ACHICAR = 2

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla     []lista.Lista[parClaveValor[K, V]]
	capacidad int
	cantidad  int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := creartabla[K, V](CAPACIDAD_INICIAL)
	return &hashAbierto[K, V]{tabla: tabla, capacidad: CAPACIDAD_INICIAL}
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	iterador := h.buscarClave(clave)
	campo := parClaveValor[K, V]{clave, dato}

	if iterador.HaySiguiente() {
		iterador.Borrar()
		h.cantidad--
	}

	iterador.Insertar(campo)
	h.cantidad++

	if h.esRedimensionable() {
		h.redimension(h.capacidad * FACTOR_REDIMENSION_AGRANDAR)
	}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	return h.buscarClave(clave).HaySiguiente()
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	iter := h.buscarClave(clave)

	if !iter.HaySiguiente() {
		h.verificarNoVacio()
	}

	par := iter.Borrar()
	h.cantidad--

	if h.esRedimensionable() {
		h.redimension(h.capacidad / FACTOR_REDIMENSION_ACHICAR)
	}
	return par.dato
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	iter := h.buscarClave(clave)

	if !iter.HaySiguiente() {
		h.verificarNoVacio()
	}
	return iter.VerActual().dato
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	seguir := true
	for i := 0; i < h.capacidad && seguir == true; i++ {

		h.tabla[i].Iterar(func(par parClaveValor[K, V]) bool {
			seguir = visitar(par.clave, par.dato)
			return seguir
		})
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	for i, lista := range h.tabla {
		iter := lista.Iterador()

		if iter.HaySiguiente() {
			return &iterHash[K, V]{act: iter, posicion: i, tabla: h.tabla}
		}
	}
	return &iterHash[K, V]{act: h.tabla[h.capacidad-1].Iterador(), posicion: h.capacidad, tabla: h.tabla}
}

type iterHash[K comparable, V any] struct {
	act      lista.IteradorLista[parClaveValor[K, V]]
	tabla    []lista.Lista[parClaveValor[K, V]]
	posicion int
}

func (i *iterHash[K, V]) HaySiguiente() bool {
	return i.posicion < len(i.tabla)
}

func (i *iterHash[K, V]) Siguiente() {
	i.validarNoVacio()
	i.act.Siguiente()

	for !i.act.HaySiguiente() && i.HaySiguiente() {
		i.posicion++

		if i.HaySiguiente() {
			i.act = i.tabla[i.posicion].Iterador()
		}
	}
}

func (i *iterHash[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic(TERMINO_ITERAR)
	}
	return i.act.VerActual().clave, i.act.VerActual().dato
}

func (h *hashAbierto[K, V]) esRedimensionable() bool {
	if (h.cantidad < h.capacidad/CRITERIO_REDIMENSION_ACHICAR && h.capacidad >= CAPACIDAD_INICIAL*CRITERIO_REDIMENSION_ACHICAR) || h.cantidad/h.capacidad > CRITERIO_REDIMENSION_AGRANDAR {
		return true
	}
	return false
}

func (h *hashAbierto[K, V]) redimension(nueva_capacidad int) {
	nueva_tabla := creartabla[K, V](nueva_capacidad)
	for _, lista := range h.tabla {
		iter := lista.Iterador()

		for iter.HaySiguiente() {
			pos := fnv.Hashing(iter.VerActual().clave, nueva_capacidad)
			nueva_tabla[pos].InsertarUltimo(iter.VerActual())
			iter.Siguiente()
		}
	}
	h.capacidad = nueva_capacidad
	h.tabla = nueva_tabla
}

func creartabla[K comparable, V any](capacidad int) []lista.Lista[parClaveValor[K, V]] {
	//Inicializa las listas enlazadas en cada posición de la tabla del hash
	tabla := make([]lista.Lista[parClaveValor[K, V]], capacidad)

	for i := range tabla {
		tabla[i] = lista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}

func (h *hashAbierto[K, V]) buscarClave(clave K) lista.IteradorLista[parClaveValor[K, V]] {
	//Crea un iterador y lo deja parado en la posición cuyas clave coincide con la pasada por parámetro
	pos := int(fnv.Hashing(clave, h.capacidad))
	iter := h.tabla[pos].Iterador()
	for iter.HaySiguiente() {

		if iter.VerActual().clave == clave {
			return iter
		}

		iter.Siguiente()
	}
	return iter
}

func (h hashAbierto[K, V]) verificarNoVacio() {
	panic(NO_PERTENECE)
}

func (i iterHash[K, V]) validarNoVacio() {
	if !i.HaySiguiente() {
		panic(TERMINO_ITERAR)
	}
}
