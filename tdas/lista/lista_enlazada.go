package lista

const _ERROR_LISTA = "La lista esta vacia"

const _ERROR_ITERADOR = "El iterador termino de iterar"

type listaEnlazada[T any] struct { //Estructura de lista de forma enlazada, tiene como estructuras un puntero al primer nodo,otro al ultimo y su largo
	prim  *nodo[T]
	ult   *nodo[T]
	largo int
}

type nodo[T any] struct { //La estructura nodo posee un dato y un puntero a su proximo
	dato T
	prox *nodo[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	//Crea una lista enlazada vacía, sin referencias y devuelve una lista
	return new(listaEnlazada[T])
}

func (l listaEnlazada[T]) EstaVacia() bool { //Verifica que la lista enlazada este vacia
	return l.largo == 0
}
func (l *listaEnlazada[T]) InsertarPrimero(elem T) { //Inserta un elemento al principio de la LE
	nuevo := crearNodo(elem)
	if l.EstaVacia() {
		l.ult = nuevo
	} else {
		nuevo.prox = l.prim
	}
	l.prim = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) { //Inserta un elemento al final de la LE
	nuevo := crearNodo(elem)
	if l.EstaVacia() {
		l.prim = nuevo
	} else {
		l.ult.prox = nuevo
	}
	l.ult = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T { //Borra el primer elemento de la LE y lo devuelve
	l.verificarVacia()
	dato := l.prim.dato
	if l.prim == l.ult {
		l.ult = nil
	}
	l.prim = l.prim.prox
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T { //Devuelve el primer elemento de la LE. Si no hay ocurre un panic
	l.verificarVacia()
	return l.prim.dato
}

func (l *listaEnlazada[T]) VerUltimo() T { //Devuelve el ultimo elemento de la LE. Si no hay ocurre un panic
	l.verificarVacia()
	return l.ult.dato
}

func (l listaEnlazada[T]) Largo() int { //Devuelve el largo de la LE
	return l.largo
}

func (l listaEnlazada[T]) Iterar(visitar func(T) bool) { //Iterador interno que recibe una funcion para controlar la iteracion
	act := l.prim
	for act != nil && visitar(act.dato) {
		act = act.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] { //Devuelve un iterador externo
	return &iteradorListaEnlazada[T]{act: l.prim, lista: l}
}

type iteradorListaEnlazada[T any] struct { //Estructura de iterador de LE. Posee una referencia al nodo actual,al anterior y a la lista enlazada para asi poder modificar su largo,prim y ult
	ant   *nodo[T]
	act   *nodo[T]
	lista *listaEnlazada[T]
}

func (iter *iteradorListaEnlazada[T]) VerActual() T { //Devuelve el dato de actual. Si no hay actual da un panic
	iter.verificarVacio()
	return iter.act.dato
}
func (iter iteradorListaEnlazada[T]) HaySiguiente() bool { //Devuelve un booleano si hay un actual
	return iter.act != nil
}

func (iter *iteradorListaEnlazada[T]) Borrar() T { //Borra el elemento actual de la iteracion, despues el actual es el proximo al que se borro
	iter.verificarVacio()
	dato := iter.act.dato
	if !iter.apuntaAlPrimero() {
		iter.ant.prox = iter.act.prox
	}
	if iter.apuntaAlUltimo() {
		iter.lista.ult = iter.ant
	}
	if iter.apuntaAlPrimero() {
		iter.lista.prim = iter.act.prox
	}
	iter.act = iter.act.prox
	iter.lista.largo--
	return dato
}

func (iter *iteradorListaEnlazada[T]) Siguiente() { //Avanza el iterador sobre la lista
	iter.verificarVacio()
	iter.ant = iter.act
	iter.act = iter.act.prox
}

func (iter *iteradorListaEnlazada[T]) Insertar(item T) { //Inserta un nodo entre ant y act. Despues de insertarlo el act es el nuevo nodo
	nodo := crearNodo[T](item)
	aux := iter.act
	if !iter.HaySiguiente() && iter.apuntaAlPrimero() {
		iter.lista.prim = nodo
		iter.lista.ult = nodo
	} else if iter.apuntaAlPrimero() {
		iter.lista.prim = nodo
	} else if !iter.HaySiguiente() {
		iter.ant.prox = nodo
		iter.lista.ult = nodo
	} else {
		iter.ant.prox = nodo
	}
	nodo.prox = aux
	iter.act = nodo
	iter.lista.largo++
}

func crearNodo[T any](datos T) *nodo[T] { //Crea nodo
	return &nodo[T]{dato: datos}
}

func (iter *iteradorListaEnlazada[T]) verificarVacio() {
	if !iter.HaySiguiente() {
		panic(_ERROR_ITERADOR)
	}
}

func (l *listaEnlazada[T]) verificarVacia() {
	if l.EstaVacia() {
		panic(_ERROR_LISTA)
	}
}

func (iter *iteradorListaEnlazada[T]) apuntaAlUltimo() bool {
	return iter.act == iter.lista.ult
}

func (iter *iteradorListaEnlazada[T]) apuntaAlPrimero() bool {
	return iter.act == iter.lista.prim
}
