package cola_prioridad

const ERROR_MSG = "La cola esta vacia"
const CRITERIO_REDIMENSION_AGRANDAR = 2
const CRITERIO_REDIMENSION_ACHICAR = 4
const FACTOR_REDIMENSION = 2
const CAPACIDAD_INICIAL = 15

type heap[T any] struct {
	arreglo   []T
	cantidad  int
	capacidad int
	cmp       funcCmp[T]
}

type funcCmp[T any] func(T, T) int

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, CAPACIDAD_INICIAL)
	return &heap[T]{cmp: funcion_cmp, arreglo: arr, capacidad: CAPACIDAD_INICIAL}
}

func CrearHeapArr[T any](arr []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := heapifyArr[T](arr, funcion_cmp)
	return &heap
}

func (h *heap[T]) VerMax() T {
	h.verificarNoVacio()
	return h.arreglo[0]
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func (h *heap[T]) Encolar(item T) {
	h.arreglo[h.cantidad] = item
	h.upHeap(h.cantidad)
	h.cantidad++
	if h.esRedimensionable() {
		h.redimensionar(h.capacidad * CRITERIO_REDIMENSION_AGRANDAR)
	}
}

func (h *heap[T]) Desencolar() T {
	h.verificarNoVacio()
	elemento := h.VerMax()
	h.cantidad--
	h.swappearPadreHijo(0, h.cantidad)

	h.downHeap(0, h.cantidad-1)

	if h.esRedimensionable() {
		h.redimensionar(h.capacidad / FACTOR_REDIMENSION)
	}

	return elemento
}

func HeapSort[T any](arr []T, cmp func(T, T) int) []T {
	heap := heapifyArr(arr, cmp)

	i := heap.cantidad - 1

	for !heap.EstaVacia() {
		arr[i] = heap.Desencolar()
		i--
	}

	return arr
}

func heapifyArr[T any](arr []T, cmp funcCmp[T]) heap[T] {
	capacidad := ajustarCapacidad(arr)
	nuevo := make([]T, capacidad*2)
	copy(nuevo, arr)
	h := heap[T]{cmp: cmp, arreglo: nuevo, cantidad: len(arr), capacidad: capacidad * 2}

	for i := h.cantidad - 1; i >= 0; i-- {
		h.arreglo = h.downHeap(i, h.cantidad-1)
	}

	return h
}

func (h *heap[T]) obtenerHijoMayor(pos_padre int) int {
	pos_hijo_der := obtenerHijoDer(pos_padre)
	pos_hijo_izq := obtenerHijoIzq(pos_padre)

	if pos_hijo_der >= len(h.arreglo) {
		return pos_hijo_izq
	}

	if h.cmp(h.arreglo[pos_hijo_izq], h.arreglo[pos_hijo_der]) >= 0 {
		return pos_hijo_izq
	}

	return pos_hijo_der
}

func (h *heap[T]) esRedimensionable() bool {
	if h.cantidad < h.capacidad/CRITERIO_REDIMENSION_ACHICAR && h.cantidad >= CAPACIDAD_INICIAL*CRITERIO_REDIMENSION_AGRANDAR {
		return true
	}

	if h.cantidad*CRITERIO_REDIMENSION_AGRANDAR > h.capacidad {
		return true
	}
	return false
}

func (h *heap[T]) redimensionar(nueva_capacidad int) { //redimensiona dependiendo la situacion. Recibe por parametro la situacion, si se llama de desapilar es true sino es false
	dst := make([]T, nueva_capacidad)
	copy(dst, h.arreglo)

	h.arreglo = dst
	h.capacidad = nueva_capacidad
}

func (h *heap[T]) upHeap(pos_hijo int) []T {
	pos_padre := obtenerPosicionPadre(pos_hijo)

	if !h.esNecesarioAplicarUpheap(pos_hijo, pos_padre) {
		return h.arreglo
	}

	h.swappearPadreHijo(pos_hijo, pos_padre)
	return h.upHeap(pos_padre)
}

func (h *heap[T]) downHeap(pos_padre int, ultimo_relativo int) []T {
	pos_hijo_izq := obtenerHijoIzq(pos_padre)
	pos_hijo_der := obtenerHijoDer(pos_padre)

	if !verificarPosicionesObtenidasValidas(pos_hijo_izq, pos_hijo_der, ultimo_relativo) {
		return h.arreglo
	}

	pos_hijo_mayor := pos_hijo_izq

	if verificarPosicionHijoDer(pos_hijo_der, ultimo_relativo) {
		pos_hijo_mayor = h.obtenerHijoMayor(pos_padre)
	}

	if !h.validarPadreMayorAHijo(pos_hijo_mayor, pos_padre) {
		h.swappearPadreHijo(pos_padre, pos_hijo_mayor)
		return h.downHeap(pos_hijo_mayor, ultimo_relativo)
	}
	return h.arreglo

}

func (h heap[T]) verificarNoVacio() {
	if h.EstaVacia() {
		panic(ERROR_MSG)
	}
}

func (h heap[T]) swappearPadreHijo(pos1, pos2 int) {
	h.arreglo[pos1], h.arreglo[pos2] = h.arreglo[pos2], h.arreglo[pos1]
}

func (h *heap[T]) validarPadreMayorAHijo(pos_hijo_mayor, pos_padre int) bool {
	if h.cmp(h.arreglo[pos_padre], h.arreglo[pos_hijo_mayor]) < 0 {
		return false
	}

	return true
}

func (h heap[T]) esNecesarioAplicarUpheap(pos_hijo, pos_padre int) bool {
		if pos_hijo == 0 || h.cmp(h.arreglo[pos_hijo], h.arreglo[pos_padre]) <= 0 {
			return false
		}

	return true
}

func ajustarCapacidad[T any](arr []T) int {
	if len(arr) < CAPACIDAD_INICIAL {
		return CAPACIDAD_INICIAL
	}
	return cap(arr)
}

func obtenerPosicionPadre(pos_hijo int) int {
	return (pos_hijo - 1) / 2
}

func verificarPosicionHijoDer(pos_hijo_der, ultimo_relativo int) bool {
	if pos_hijo_der <= ultimo_relativo {
		return true
	}

	return false
}

func verificarPosicionesObtenidasValidas(pos_hijo_izq, pos_hijo_der, ultimo_relativo int) bool {
	if pos_hijo_der > ultimo_relativo && pos_hijo_izq > ultimo_relativo {
		return false
	}

	return true
}

func obtenerHijoIzq(pos_padre int) int {
	return 2*pos_padre + 1
}

func obtenerHijoDer(pos_padre int) int {
	return 2*pos_padre + 2
}
