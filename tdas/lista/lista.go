package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el dato al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista. Si la lista se encontraba vacía, entra el pánico con el
	// mensaje 'La lista esta vacia'.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al inicio de la lista (el primero). Si la lista se encontraba vacía, entra en
	// pánico con el mensaje 'La lista esta vacia'.
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista (el ultimo). Si la lista se encontraba vacía, entra em
	// pánico con el mensaje 'La lista esta vacia'.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	// Iterar aplica la funcion pasada por parametro a todos los elementos de la lista, hasta que no hayan más
	// elementos, o la función en cuestión devualva false.
	Iterar(visitar func(T) bool)

	// Iterador crea un IteradorExterno para la lista usada que permite recorrer todos los elementos de la lista.

	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual en la iteración.
	VerActual() T

	// HaySiguiente devuelve verdadero si el elemento por iterar no es nil, falso en caso contrario.
	HaySiguiente() bool

	// Siguiente pasa al siguiente elemento a iterar.
	Siguiente()

	// Insertar inserta el elemento pasado por parámetro detrás del actual.
	Insertar(T)

	// Borra el elemento actual en la iteración.
	Borrar() T
}
