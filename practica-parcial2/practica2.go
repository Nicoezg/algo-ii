package practica

//  Implementar en lenguaje Go una función recursiva con la firma func esHeap(arr []int).
// Esta función debe devolver true o false de acuerdo a si el arreglo que recibe como parámetro
// cumple la propiedad de heap (de mínimos).

// Hacer el seguimiento de la función para el arreglo [ 1, 7, 2, 8, 7, 6, 3, 3, 9, 10 ].

func esHeap(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	if arr[(len(arr)-2)/2] > arr[len(arr)-1] {
		return false
	}
	return esHeap(arr[:len(arr)-1])
}

func (h heap[T]) Max3() []T {
	var res []T
	for i := 0; i < 3 && i < h.cantidad; i++ {
		res = append(res, h.datos[i])
	}
	if len(res) == 3 && h.cmp(res[1], res[2]) < 0 { //si res[1] menor que res[2]
		res[1], res[2] = res[2], res[1]
	}
	return res
}

// (★★★) Escribir una función en Go que, dado un arreglo de  n cadenas y un entero positivo  k, devuelva una lista con las
//k cadenas más largas. Se espera que el orden del algoritmo sea O(n+klogn). Justificar el orden.

func CadenasLargas(arreglo []string, k int) Lista[K] {
	lista := CrearListaEnlazada[string]()
	heap := CrearHeapArr[string](compararCadenas) // Usa heapify O(n)
	contador := 0
	for contador != k() { // O(K)
		lista.InsertarUltimo(heap.Desencolar()) // O(log n)
		contador++
	}
	return lista
}

// Complejidad final : O(n + k log n)

func CadenasLargas(arr []string, k int) []string {
	res := make([]string, k)
	if k > len(arr) {
		res = make([]string, len(arr))
	}
	heap := CrearHeapArr[string](arr, func(a, b string) int {
		if len(a) > len(b) {
			return +1
		}
		if len(a) < len(b) {
			return -1
		}
		return 0
	})

	for i := 0; i < k && i < len(arr); i++ {
		res[i] = heap.Desencolar()
	}
	return res
}

func compararCadenas(a string, b string) int { // a = 1 y b = 0
	return len(a) - len(b)
}

func compararCadenasMin(a []string, b []string) int { // a = 1 y b = 0
	return b[0] - a[0]
}

// (★★★★) Se tienen k arreglos de enteros previamente ordenados y se quiere obtener un arreglo
// ordenado que contenga a todos los elementos de los k arreglos. Sabiendo que cada arreglo tiene tamaño h, definimos como
// n a la sumatoria de la cantidad de elementos de todos los arreglos, es decir, n=k×h.

//Escribir en Go una función func KMerge(arr [][]int) que reciba los k arreglos y devuelva uno nuevo con los
//n elementos ordenados entre sí. La función debe ser de orden O(nlogk). Justificar el orden del algoritmo propuesto.

func KMerge(arr [][]int) []int {
	res := make([]int, len(arr[0]))
	heap := CrearHeapArr(arr, compararCadenasMin)

	n := len(arr[0]) * len(arr)

	for i := 0; i < n; i++ {
		arr1 := heap.Desencolar()
		res[i] = arr1[0]
		arr1 = arr1[1:]
		if len(arr1) == 0 {
			continue
		}
		heap.Encolar(arr1)
	}
	return res

}

// (★★★) Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números
// dentro del arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo).
// La solución debe ser mejor que O(nlogn) si k << n. Indicar y justificar la complejidad de la función implementada.

func suma_max(arr []int, k int) {
	heap := CrearHeap[int](func(a, b) int { return b - a })
	suma_max := 0
	contador := 0
	for i, elemento := range arr {
		if contador < kelementoheap.Encolar(numero)
			suma_maxelementomero
			contador++
			continue
		}
		if elemento > heap.VerMax() {
			suma_max -= heap.VerMax()
elementop.Desenelemento)
			heap.Encolar(numero)
			suma_max += numero
		}
	}
	return
}

func Combinar(arr []int, k int) []int {
	heap := CrearHeapArr[int](arr, func(a, b int) int { return b - a }) // O(n) , heapify

	for heap.Cantidad() >= 2 && heap.VerMax() <= k { // O()
		mas_chico := heap.Desencolar()
		mas_chico2 := heap.Desencolar()
		if mas_chico2 > k {
			return nil
		}
		heap.Encolar(mas_chico + mas_chico2*2)
	}

	if heap.VerMax() <= k {
		return nil
	}

	res := make([]int, heap.Cantidad())
	for i, _ := range res {
		res[i] = heap.Desencolar()
	}
	return res
}

// 1. Implementar en Go una primitiva de ABB (DiccionarioOrdenado) que reciba un número entero N
// y devuelva la cantidad de nodos del árbol que tienen exactamente N descendientes. Indicar y justificar
// la complejidad de la primitiva implementada.
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

// 2. Sabiendo que se tiene un diccionario en que las claves y valores son cadenas, implementar una función
// DiccionarioInvertirUnico que devuelva un nuevo Diccionario cuyas claves sean los valores del
// original y sus valores asociados sean las claves que originalmente tenían asociados dichos valores. En
// el caso algún valor esté repetido en el Diccionario original, se debe lanzar un panic con el mensaje
// hay valores repetidos. Indicar y justificar la complejidad de la función.

func DiccionarioInvertirUnico(dicc Diccionario[string, string]) Diccionario[string, string] {
	res := CrearHash[string, string]()
	iter := dicc.Iterador()
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		if res.Pertenece(valor) {
			panic("Hay valores repetidos")
		}
		res.Guardar(valor, clave)
		iter.Siguiente()
	}
	return res
}

// Se tienen k arreglos de enteros previamente ordenados y se quiere obtener un arreglo ordenado que contenga a todos los
// elementos de los k arreglos. Sabiendo que cada arreglo tiene tamaño h, definimos como n a la sumatoria de la cantidad
// de elementos de todos los arreglos, es decir, n = k*h.
// Escribir en Go la función func KMerge(arreglos [][]int, k int, h int) []int que reciba los k arreglos y devuelva
// uno nuevo con los n elementos ordenados entre sí.
// La función debe ser de orden O(n log k). Justificar el orden del algoritmo propuesto.

func KMerge(arreglos [][]int, k int, h int) []int {
	n := k * h
	res := make([]int, n)
	heap := CrearHeapArr(arr, func(a, b []int) int { return b[0] - a[0] })

	for i, _ := range res {
		arr1 := heap.Desencolar()
		res[i] = arr1[0]
		arr1[1:]
		if len(arr1) != 0 {
			heap.Encolar(arr1)
		}
	}

	return res
}

//Implementar una función func Xor[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] que
//devuelva un nuevo Diccionario que contenga los elementos que están en el primer diccionario pero no en el segundo, y
//los elementos que estén en el segundo diccionario pero no en el primero.

func Xor[K comparable, V any](d1, d2 Diccionario[K, V]) Diccionario[K, V] {
	res := CrearHash[K, V]()
	iter1 := d1.Iterador()
	iter2 := d2.Iterador()

	for iter1.HaySiguiente() {
		clave, valor := iter1.VerActual()
		if !d2.Pertenece(clave) {
			res.Guardar(clave, valor)
		}
		iter1.Siguiente()
	}

	for iter2.HaySiguiente() {
		clave, valor := iter2.VerActual()
		if !d1.Pertenece(clave) {
			res.Guardar(clave, valor)
		}
		iter2.Siguiente()
	}

	return res
}

// 3. Implementar la primitiva func (ab *ArbolBinario) EsAbb(cmp funcCmp) bool que devuelva true o false de acuerdo
// a si el árbol binario cumple la condición de ser ABB. Considerar que el árbol binario tiene la siguiente estructura:
type funcCmp[K comparable] func(K, K) int

type ArbolBinario[K comparable, V any] struct {
	izq   *ArbolBinario
	der   *ArbolBinario
	clave K
	dato  V
}

// Indicar el tipo de recorrido usado y la complejidad de la primitiva.

func (ab *ArbolBinario) EsAbb(cmp funcCmp) bool {
	res := []K{}
	ab.esAbb(cmp, &res)
	for i := 0; i < len(arr)-1; i++ {
		if cmp(res[i], res[i+1]) >= 1 {
			return false
		}
	}

	return true
}

func (ab *ArbolBinario) esAbb(cmp funcCmp, condicion *[]K) {
	if ab == nil {
		return
	}
	ab.izq.esabb(cmp, condicion)
	res = append(res, ab.clave)
	ab.der.esabb(cmp, condicion)
	return
}

// Dada una lista, hacer una función que devuelva una nueva lista con todos los elementos de la lista original 
// que no estén repetidos, y en el mismo orden

func noRepetidos(lista Lista[T]) Lista[T] {
	dicc := CrearHash[T, int]()
	iter := Lista.Iterador()
	res := CrearListaEnlazada[T]()

	for iter.HaySiguiente(){
		elemento := iter.VerActual()
		if !dicc.Pertenece(elemento){
			dicc.Guardar(elemento, 1)
		}else{
			dicc.Guardar(elemento, dicc.Obtener(elemento) + 1)
		}
		iter.Siguiente()
	}

	iter := Lista.Iterador()

	for iter.HaySiguiente(){
		clave := iter.VerActual()
		if dicc.Obtener(clave) == 1{
			lista.InsertarUltimo(clave)
		}

	}

	return res

}

// Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números dentro 
//del arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo). 
//La solución debe ser mejor que O(n log n) si k << n. Indicar y justificar la complejidad de la función implementada.

func maximizarSuma(arr []int, k int) []int{
	heap := CrearHeapArr[int](arr, func(a, b int)int{ return a - b})
	res := make([]int, k)
	for i := 0; i < k; i++{
		res[i] = append(res, heap.Desencolar())
	}
	return res
}

// O(k log n), desencolo k veces en un heap de n elementos

// AB es AVL?

func (arbol *ab[K]) EsAVL(cmp) bool {
	if arbol == nil{
		return true
	}
	altura_izq := arbol.izq.Altura() + 1
	altura_der := arbol.der.Altura() + 1
	if math.Abs(altura_izq - altura_der) > 1{
		return false
	}

func (arbol *ab[K]) EsAVL(funcion_cmp func(K, K) int) (bool,int) {
	if arbol == nil {	return true, 0	}
	esAVLizq, alturaIzq:= arbol.izq.EsAVL(cmp)
	esAVLder, alturaDer:= arbol.der.EsAVL(cmp)
	altura := math.Max(alturaIzq, alturaDer) + 1
	if !esAVLizq || !esAVLder || math. Abs(alturaIzq - alturaDer) > 1 {
		return false, altura
	}
	if ab.izq != nil && cmp(ab.dato, ab.izq.dato) < 0 {
		return false, altura
	}
	if ab.der != nil && cmp(ab.dato, ab.der.dato) > 0 {
		return false, altura
	}
	return true, altura
	}


func (heap *Heap[K]) upHeap(pos_hijo int){
	pos_padre := (pos_hijo - 1) / 2
	if pos_hijo == 0 || heap.cmp(heap.arr[pos_padre], heap.arr[pos_hijo]) >= 0{
		return
	}
	heap.arr[pos_padre], heap.arr[pos_hijo] = heap.arr[pos_hijo], heap.arr[pos_padre]
	return upHeap(pos_padre)
}

func (heap *Heap[K]) downHeap(pos_padre, ult int){
	hijo_izq := pos_padre * 2 + 1
	hijo_der := pos_padre * 2 + 2
	if hijo_izq >= ult && hijo_der >= ult{
		return 
	}
	hijo_max := hijo_izq
	if hijo_der <= ult{
		if heap.cmp(heap.arr[hijo_izq], heap.arr[hijo_der]) <= -1{
			hijo_max = hijo_der
		}
	}
	if heap.cmp(heap.arr[hijo_max], heap.arr[pos_padre]) >= 1{
		heap.arr[hijo_max] , heap.arr[pos_padre] = heap.arr[pos_padre], heap.arr[hijo_max]
		downHeap(hijo_max, ult)
	}
	return
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

func HeapSort[T any](arr []T, cmp funcCmp[T]) arr[]T {
	heap := CrearHeapArr[T](arr, cmp)
	for i, valor := range arr{
		arr[i] = heap.Desencolar()
	}
	return arr
}
	