package diccionario

import (
	"tdas/cola"
	"tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	par       *parClaveValor[K, V]
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func CrearABB[K comparable, V any](funcion_cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return *buscarNodo(&abb.raiz, clave, abb.cmp) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo_deseado := buscarNodo(&abb.raiz, clave, abb.cmp)

	abb.validarClave(*nodo_deseado)
	return (*nodo_deseado).par.dato
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	pos_arbol := buscarNodo(&abb.raiz, clave, abb.cmp)

	if *pos_arbol == nil {
		nodo_nuevo := crearNodo(clave, dato)

		abb.cantidad++
		*pos_arbol = nodo_nuevo
		return
	}

	(*pos_arbol).par.dato = dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := buscarNodo(&abb.raiz, clave, abb.cmp)

	abb.validarClave(*nodo)
	abb.cantidad--
	dato := (*nodo).par.dato

	if (*nodo).izquierdo == nil {
		(*nodo) = (*nodo).derecho
		return dato
	}

	if (*nodo).derecho == nil {
		(*nodo) = (*nodo).izquierdo
		return dato
	}

	reemplazo := buscarPosibleReemplazo(&(*nodo).derecho)
	(*nodo).par.dato = (*reemplazo).par.dato
	(*nodo).par.clave = (*reemplazo).par.clave
	(*reemplazo) = (*reemplazo).derecho

	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.raiz.iterarRango(nil, nil, visitar, abb.cmp)
}

// Implementar en Go una primitiva de ABB (DiccionarioOrdenado) que funcione como un iterador
// interno que haga un recorrido por niveles inverso. Es decir, que visite los elementos del nivel
// más inferior hasta la raiz. Para el ABB cuyo preorder es 5, 2, 1, 3, 4, 7, 9 (comparación numérica habitual),
// el recorrido debe ser: 4, 9, 3, 1, 7, 2, 5. En el 4to nivel está sólo el 4. En el 3er nivel están el 1, 3 y 9
// pero hay que leerlos de derecha a izquierda. Luego en el 2do el 2 y 7 con la misma lógica, y finalmente la raíz 5 al final.
//Indicar y justificar la complejidad de la primitiva implementada.

func (abb *abb[K, V]) IteradorInternoNivelesInverso(visitar func(K, V) bool) {
	pila := pila.CrearPilaDinamica[*parClaveValor[K, V]]()
	abb.raiz._IteradorInternoInverso(pila, visitar)
}

func (nodo *nodoAbb[K, V]) _IteradorInternoInverso(pila pila.Pila[*parClaveValor[K, V]], visitar func(K, V) bool) bool {
	q := cola.CrearColaEnlazada[*nodoAbb[K, V]]()
	q.Encolar(nodo)
	for !q.EstaVacia() {
		nodo := q.Desencolar()
		pila.Apilar(nodo.par)

		if nodo.izquierdo != nil {
			q.Encolar(nodo.izquierdo)
		}
		if nodo.derecho != nil {
			q.Encolar(nodo.derecho)
		}
	}

	for !pila.EstaVacia() {
		par := pila.Desapilar()
		visitar(par.clave, par.dato)
	}
	return true

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	act := abb.raiz
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()

	iter := &iteradorAbb[K, V]{pila: pila, desde: desde, hasta: hasta, cmp: abb.cmp}
	iter.actualizarPila(act)
	return iter
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

type iteradorAbb[K comparable, T any] struct {
	pila  pila.Pila[*nodoAbb[K, T]]
	desde *K
	hasta *K
	cmp   funcCmp[K]
}

func (iter *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorAbb[K, T]) Siguiente() {
	iter.verificarNoVacio()

	act := iter.pila.Desapilar()
	act = act.derecho
	iter.actualizarPila(act)
}

func (iter *iteradorAbb[K, T]) VerActual() (K, T) {
	iter.verificarNoVacio()
	return iter.pila.VerTope().par.clave, iter.pila.VerTope().par.dato
}

func (abb *abb[K, V]) NDescendientes(n int) int {
	contador := 0
	abb.raiz._nDescendientes(n, &contador)
	return contador
}

func (nodo *nodoAbb[K, V]) _nDescendientes(n int, contador *int) int {
	if nodo == nil {
		return -1
	}
	descendientes_izq := nodo.izquierdo._nDescendientes(n, contador) + 1
	descendientes_der := nodo.derecho._nDescendientes(n, contador) + 1
	if descendientes_izq+descendientes_der == n {
		*contador++
	}
	return descendientes_izq + descendientes_der

}

func buscarPosibleReemplazo[K comparable, V any](nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).izquierdo != nil {
		return buscarPosibleReemplazo(&(*nodo).izquierdo)
	}
	return nodo
}

func (nodo *nodoAbb[K, V]) iterarRango(desde, hasta *K, visitar func(K, V) bool, cmp funcCmp[K]) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || cmp(nodo.par.clave, *desde) >= 0 {
		if !nodo.izquierdo.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}

	if (desde == nil || cmp(nodo.par.clave, *desde) >= 0) && (hasta == nil || cmp(nodo.par.clave, *hasta) <= 0) {
		if !visitar(nodo.par.clave, nodo.par.dato) {
			return false
		}
	}

	if hasta == nil || cmp(nodo.par.clave, *hasta) <= 0 {
		if !nodo.derecho.iterarRango(desde, hasta, visitar, cmp) {
			return false
		}
	}
	return true
}

func (abb *abb[K, V]) validarClave(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic(NO_PERTENECE)
	}
}

func (iter iteradorAbb[K, V]) validarNoVacio() {
	if !iter.HaySiguiente() {
		panic(TERMINO_ITERAR)
	}
}

func (iter *iteradorAbb[K, V]) verificarNoVacio() {
	if !iter.HaySiguiente() {
		panic(TERMINO_ITERAR)
	}
}

func buscarNodo[K comparable, V any](nodo **nodoAbb[K, V], clave K, cmp funcCmp[K]) **nodoAbb[K, V] {
	if *nodo == nil {
		return nodo
	}
	diferencia := cmp((*nodo).par.clave, clave)

	if diferencia == 0 {
		return nodo
	}

	if diferencia > 0 {
		return buscarNodo(&(*nodo).izquierdo, clave, cmp)
	}

	return buscarNodo(&(*nodo).derecho, clave, cmp)

}

func (iter *iteradorAbb[K, T]) actualizarPila(nodo *nodoAbb[K, T]) {
	if nodo == nil {
		return
	}
	if iter.hasta != nil && iter.cmp(nodo.par.clave, *iter.hasta) > 0 {
		iter.actualizarPila(nodo.izquierdo)
		return

	}
	if iter.desde != nil && iter.cmp(nodo.par.clave, *iter.desde) < 0 {
		iter.actualizarPila(nodo.derecho)
		return
	}
	iter.pila.Apilar(nodo)
	iter.actualizarPila(nodo.izquierdo)
	return
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &(nodoAbb[K, V]{par: &parClaveValor[K, V]{clave: clave, dato: dato}})
}
