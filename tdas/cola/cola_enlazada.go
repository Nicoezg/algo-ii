package cola

/* Una cola es una estructura de datos lineal que se basa en el
principio FIFO (first in, first out). */

type colaEnlazada[T any] struct {
	prim *nodo[T]
	ult  *nodo[T]
}

type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	/* Crea una colaEnlazada vacia, sin referenicas y devuelve una cola.*/
	return new(colaEnlazada[T])
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	nodo := crearNodo[T](elem)
	if c.EstaVacia() {
		c.prim = nodo
	} else {
		c.ult.prox = nodo
	}
	c.ult = nodo

}

func (c *colaEnlazada[T]) Desencolar() T {
	validarNoVacia(c)
	dato := c.prim.dato
	c.prim = c.prim.prox
	if c.prim == nil {
		c.ult = nil
	}
	return dato

}

func (c *colaEnlazada[T]) VerPrimero() T {
	validarNoVacia(c)
	return c.prim.dato

}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.prim == nil
}

func crearNodo[T any](elem T) *nodo[T] {
	return &nodo[T]{elem, nil}

}

func validarNoVacia[T any](c *colaEnlazada[T]) bool {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return true
}
