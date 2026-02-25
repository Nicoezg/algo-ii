package pila

/* Una pila es una estructura de datos que permite almacenar
y recuperar datos, siendo el modo de acceso a sus elementos de tipo LIFO (last in, first out).*/

const _CAP_INICIAL = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	/* Crea una pilaDinamica vacia, sin elementos y devuelve una pila.*/
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAP_INICIAL, _CAP_INICIAL)
	pila.cantidad = 0
	return pila
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	/* Apila el elemento pasado por parametro, redimensionando si es necesario.*/
	p.datos[p.cantidad] = elem
	p.cantidad++
	p.datos = redimensionar(p.datos, p.cantidad)
}

func (p *pilaDinamica[T]) Desapilar() T {
	/* Desapila el elemento que se encuentra en el tope, redimensionando si es necesario.
	Si no hay elementos en la pila, arroja un panic.*/
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	p.cantidad--
	elem := p.datos[p.cantidad]
	p.datos = redimensionar(p.datos, p.cantidad)
	return elem

}

func (p *pilaDinamica[T]) VerTope() T {
	/* Devuelve el tope de la pila. Si la pila esta vacia, arroja un panic.*/
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]

}

func (p *pilaDinamica[T]) EstaVacia() bool {
	/* Devuelve true si la pila esta vacía y false en caso contrario.*/
	return p.cantidad == 0
}

func redimensionar[T any](arreglo []T, cant int) []T {

	cap := cap(arreglo)
	if cant == cap {
		s := make([]T, cap*2, cap*2)
		copy(s, arreglo)
		return s
	}
	if cant*4 <= cap {
		s := make([]T, cap/2, cap/2)
		copy(s, arreglo)
		return s
	}
	return arreglo
}
