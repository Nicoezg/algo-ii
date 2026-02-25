package practica

import (
	"tdas/cola"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/lista"
	"tdas/pila"
)

func Ej1Parcial(arreglo []int, k int) []int {
	heap := cola_prioridad.CrearHeap[int](func(a, b int) int {
		return b - a
	})
	max_actual := 0
	for i, numero := range arreglo {
		if i >= k {

			if heap.VerMax() < numero {
				max_actual -= heap.Desencolar()
				max_actual += numero
				heap.Encolar(numero)
			}
			arreglo[i] = max_actual

		} else if i < k-1 {
			arreglo[i] = -1
			heap.Encolar(numero)
			max_actual += numero

		} else {
			heap.Encolar(numero)
			max_actual += numero
			arreglo[i] = max_actual

		}
	}

	return arreglo
}

// [1,5,3,4,2,8], k = 3

func dictMedio[K comparable, T any](dict diccionario.Diccionario[K, lista.Lista[T]]) diccionario.Diccionario[K, T] {
	res := diccionario.CrearHash[K, T]()
	iter := dict.Iterador()

	for iter.HaySiguiente() {
		clave, lista := iter.VerActual()
		iterador_lista := lista.Iterador()
		contador := 0
		for iterador_lista.HaySiguiente() {
			if contador == lista.Cantidad()/2 {
				año := iterador_lista.VerActual()
				res.Guardar(clave, año)
				break
			}
			contador += 1
			iterador_lista.Siguiente()

		}
		iter.Siguiente()
	}

	return res

}

func (abb *abb[K, V]) primerosMNiveles(M int, inicio, final K) lista.Lista[V] {
	res := lista.CrearListaEnlazada[V]()
	contador := 0
	return abb._primerosMNiveles(M, inicio, final, lista, abb.cmp, contador)
}

func (nodo *NodoAbb[K, V]) _primerosMNiveles(M int, inicio, final K, lista lista.Lista[V], cmp funccmp, contador int) lista.Lista[V] {
	if nodo == nil || contador >= M {
		return lista
	}

	if cmp(nodo.clave, inicio) >= 0 && cmp(nodo.clave, final) <= 0 {
		lista.InsertarUltimo(nodo.clave)
	}
	if cmp(nodo.clave, inicio) >= 1 && cmp(nodo.clave, final) >= 0 {
		nodo.izq(M, inicio, final, lista, cmp, contador+1)
	}
	if cmp(nodo.clave, inicio) <= 0 && cmp(nodo.clave, final) <= -1 {
		nodo.der(M, inicio, final, lista, cmp, contador+1)
	}
	return lista

}

type multiconj[K comparable] struct {
	dicc diccionario.Diccionario[K, int]
}

func CrearMultiConj[K comparable]() *multiconj[K] {
	dicc := diccionario.CrearHash[K, int]()
	return &multiconj[K]{dicc: dicc}

}

func (multi *multiconj[K]) Guardar(elemento K) {
	if multi.Pertenece(elemento) {
		multi.dicc.Guardar(elemento, multi.dicc.Obtener(elemento)+1)
		return
	}
	multi.dicc.Guardar(elemento, 1)

}

func (multi *multiconj[K]) Pertenece(elemento K) bool {
	if multi.dicc.Pertenece(elemento) {
		return multi.dicc.Obtener(elemento) != 0
	}
	return false
}

func (multi *multiconj[K]) Borrar(elemento K) bool {
	if !multi.Pertenece(elemento) {
		panic("El elemento no pertenece")
	}
	multi.dicc.Guardar(elemento, multi.dicc.Obtener(elemento)-1)
	return true
}

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int {
	aux := diccionario.CrearHash[K, V]()
	for _, elemento := range dicc.tabla {
		if elemento.estado != "borrado" && elemento.estado != "vacio" {
			aux.Guardar(elemento.dato, elemento.clave)
		}
	}
	return aux.Cantidad()
}

//Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y un numero K
// y determinar en O(n) si existe un par de elementos en el arreglo que sumen exactamente K.

func sumaExactamenteK(arreglo []int, largo, k int) bool {
	dicc := diccionario.CrearHash[int, int]()
	for i, elem := range arreglo {
		if dicc.Pertenece(k - elem) {
			return true
		}
		dicc.Guardar(elem, i)
	}
	return false

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

}
