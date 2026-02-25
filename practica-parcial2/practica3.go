// 1. Implementar en Go una primitiva de ABB (DiccionarioOrdenado) que funcione como un iterador interno que haga un
// recorrido por niveles inverso. Es decir, que visite los elementos del nivel más inferior hasta la raiz.
//Ejemplo, la siguiente
// invocación:
// abb.iterNivelesInverso(func (clave int, _ int) {
// fmt.Printf(clave)
// return true
// })
// para el árbol del dorso deberia mostrar por pantalla los nodos en este orden: 4 9 3 1 7 2 5. En el 4to nivel está sólo
// el 4. En el 3er nivel están el 1, 3 y 9 pero hay que leerlos de derecha a izquierda. Luego en el 2do el 2 y 7 con la
//misma lógica, y finalmente la raíz 5 al final. Indicar y justificar la complejidad de la primitiva implementada.

func (abb *abb[K,V])iterNivelesInverso(visitar func (clave int, dato int){ 
	fmt.Prinf(clave)
	return true
}){
	abb.raiz._iterNivelesInverso(visitar)
}

func (nodo *nodoAbb[K,V]) _iterNivelesInverso(visitar func ( clave int, dato int) bool){
	cola := CrearColaEnlazada[nodoAbb[K,V]]()
	pila := CrearPilaDinamica[nodoAbb[K,V]]()
	cola.Encolar(nodo)
	for !cola.EstaVacia(){
		nodo := cola.Desencolar()
		pila.Apilar(nodo)
		if nodo.izq != nil{
			cola.Encolar(nodo.izq)
		}
		if nodo.der != nil{
			cola.Encolar(nodo.der)
		}
	}

	for !pila.EstaVacia(){
		nodo := pila.Desencolar()
		visitar(nodo.clave, nodo.dato)
	}
}

// Para una app de fútbol nos pidieron crear un registro del historial de partidos de la selección Argentina, de modo
// de que se puedan ver por cuántos partidos tenemos de hijo ventaja a las otras selecciones. Se pide crear una función
// func paternidad(resuls []Resultado) []Resumen que recibe un slice de Resultados, con la info de cada partido y
// devuelva un slice de Resumenes con el resumen de victorias por contrincante (ver al dorso la definición de los structs).
// Indicar y justificar la complejidad de la función implementada. Ejemplo:
// entrada: [["Francia", "v"], ["Croacia", "v"], ["Arabia", "d"], ["Uruguay", "e"],
// ["Croacia", "d"], ["Francia", "v"]]
// salida: [["Francia", 2], ["Croacia", 0], ["Arabia", -1], ["Uruguay", 0]]
type Resultado struct {
	pais string
	resultado string // "v" si ganamos, "d" perdimos y "e" empatamos
	}

type Resumen struct {
	pais string
	ventaja int // de partidos de ventaja sobre el contrincante
	}

func paternidad(resultados []Resultado) []Resumen{
	dicc := CrearHash[string, int]()

	for i, resultado := range resultados{ // O (N) siendo n los resultados, ya que guardar y pertenece en el dicc es de tiempo cosntante
		if !dicc.Pertenece(resultado.pais){
			if resultado.resultado == "v"{
				dicc.Guardar(resultado.pais, 1)
			}else if resultado.resultado == "d"{
				dicc.Guardar(resultado.pais, -1)
			}else{
				dicc.Guardar(resultado.pais, 0)
			}
			continue
		}
		if resultado.resultado == "v"{
			dicc.Guardar(resultado.pais, dicc.Obtener(resultado.pais) + 1)
			continue	
		}
		if resultado.resultado == "d"{
			dicc.Guardar(resultado.pais, dicc.Obtener(resultado.pais) - 1)
			continue
		}
	}

	res := make([]Resumen, dicc.Cantidad())
	contador := 0
	iter := dicc.Iterador()
	for iter.HaySiguiente(){ // O(P) siendo p los distintos paises en el diccionario, en el peor caso P = N
		seleccion, ventaja := iter.VerActual()
		nuevo := Resumen{pais : seleccion, ventaja : ventaja}
		res[contador] = nuevo
		contador++
		iter.Siguiente()
	}
	return res

}

// 3. Implementar en Go una primitiva Invertir() para el Heap, que haga que el heap se comporte con la función de
// comparación contraria a la que venía utilizando hasta ese momento. El heap debe quedar en estado correcto para que las
// operaciones siguientes sean válidas considerando la función de comparación, ahora invertida. Todas las operaciones deben
// seguir teniendo la complejidad esperada de un Heap. No se puede modificar la estructura del heap para implementar
// esta primitiva. Es decir, La implementación actual que tienen del heap debería poder trabajar con esta primitiva,
// sin tener que modificar ninguna primitiva ni función auxilia ni agregar campos a la estructura. Indicar y justificar la
//complejidad de la función implementada.

// O(n), cambiar la funcion de comparacion es O(1) y hacer el heapify es O(n). Downheap es O(log n)
func (heap *heap[K]) Invertir() {
	cmp := heap.cmp // func(a, b int) int{ return a - b}
	heap.cmp = cmp *(-1)
	cmp.arr = heapify(cmp.arr, heap.cmp)

}

func heapify(arr []K, cmp funcCmp[K]) arr []K{
	res := make([]K, len(arr))
	copy(res, arr)
	for i := len(arr) - 1 ; i >= 0; i--{
		downHeap(i, &arr)
	}
	return res
}

func downHeap(padre int, arr *[]K, cmp funcCmp[K]) {
	hijo_izq := padre * 2 + 1
	hijo_der := padre * 2 + 2
	if hijo_izq >= len(arr) && hijo_der >= len(arr){
		return 
	}
	max_hijo := hijo_izq
	if hijo_der < len(arr) && cmp(arr[hijo_der], arr[hijo_izq]) >= 1{
		max_hijo = hijo_der
	}

	if cmp(arr[padre], arr[hijo_max]) >= 0{
		return
	}
	arr[padre], arr[max_hijo] = arr[max_hijo], arr[padre]
	downHeap(max_hijo, arr, cmp)
}