package lista_test

import (
	TDAlista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// Sets modificables.
var _SET_UNICOELEM = []string{"hola"} // Debe contener un único elemento
var _SET_VARIOSELEMS = []bool{true}   // Debe contener al menos un elemento

// Este set no puede ser modificable
var _SET_ENTEROS = []int{10, 25, 40, 99}

const _A = 10
const _B = 25
const _C = 40
const _D = 99
const _ERROR_MSG = "La lista esta vacia"
const _TAM = 10000
const _ERROR_ITERADOR = "El iterador termino de iterar"

func TestListaVacia(t *testing.T) {
	/*- Se fija que EstaVacia() de una lista vacia devuelva True
	- Se fija que VerPrimero() y VerUltimo() de una lista vacia arroje un panic
	- Se fija que BorrarPrimero() de una lista vacia arroje un panic*/

	lista := TDAlista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestUnicoElementoGenérico(t *testing.T) {
	/*- Prueba insertando al principio y al final un entero a dos listas vacias y comprueba que se comproten de la misma manera.
	- Se fija que una vez instertado el entero la lista no este vacia.
	- Se fija que VerPrimero() y VerUltimo() devuelvan el mismo valor, y BorrarPrimero() efectivamente borra y devuelve el 1.
	- Comprueba que la lista vacia se comporte como tal.*/
	lista := crearEjemploLista(_SET_UNICOELEM)
	lista2 := crearEjemploLista(_SET_UNICOELEM)
	require.False(t, lista.EstaVacia())
	require.False(t, lista2.EstaVacia())
	require.EqualValues(t, lista.VerUltimo(), lista.VerPrimero())
	require.EqualValues(t, lista2.VerUltimo(), lista2.VerPrimero())
	require.EqualValues(t, lista.BorrarPrimero(), lista2.BorrarPrimero())
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista2.BorrarPrimero() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista2.VerPrimero() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista2.VerUltimo() })
}

func TestVolumen(t *testing.T) {
	/* - Realiza una prueba con 10000 enteros, enlistandolos, luego desenlistandolos todos y
	se fija que una vez desenlistados todos, la lista este vacia*/

	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 0; i < _TAM; i++ {
		lista.InsertarUltimo(i)
	}
	for i := 0; i < _TAM; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())

}
func TestListaOperacionesMultiples(t *testing.T) {
	/* Realiza multiples operaciones, fijandose que actue como una lista enlazada.*/

	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(_A)
	lista.InsertarPrimero(_B)
	lista.InsertarUltimo(_D)
	require.EqualValues(t, _B, lista.VerPrimero())
	require.EqualValues(t, _D, lista.VerUltimo())
	lista.BorrarPrimero()
	require.EqualValues(t, _A, lista.VerPrimero())
	lista.InsertarPrimero(_C)
	lista.InsertarUltimo(_A)
	lista.InsertarPrimero(_B)
	lista.InsertarUltimo(_D)
	require.EqualValues(t, _B, lista.BorrarPrimero())
	require.EqualValues(t, _C, lista.BorrarPrimero())
	require.EqualValues(t, _A, lista.BorrarPrimero())
	require.EqualValues(t, _D, lista.BorrarPrimero())
	require.EqualValues(t, _A, lista.BorrarPrimero())
	require.EqualValues(t, _D, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerPrimero() })
}

func TestIteradorExternoIteraListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.Borrar() })
	require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.VerActual() })
}

func TestIteradorExternoIterarLista(t *testing.T) {
	/* Itera con el iterador externo hasta el final*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.VerActual() })
}

func TestIteradorExternoInsertarPrincipio(t *testing.T) {
	/* Inserta un elemento en la posición en la que se crea el iterador.*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	iter.Insertar(_SET_VARIOSELEMS[0])
	require.EqualValues(t, _SET_VARIOSELEMS[0], lista.VerPrimero())
}

func TestIteradorExternoInsertarUltimo(t *testing.T) {
	/* Inserta un elemento en la posición final.*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	require.EqualValues(t, false, iter.HaySiguiente())
	iter.Insertar(_SET_VARIOSELEMS[0])
	require.EqualValues(t, _SET_VARIOSELEMS[0], lista.VerUltimo())
}

func TestIteradorExternoInsertarMedio(t *testing.T) {
	/* Inserta un elemento en el medio de la lista.*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	for i := 0; i < len(_SET_VARIOSELEMS)/2; i++ {
		iter.Siguiente()
	}
	iter.Insertar(_SET_VARIOSELEMS[0])
	require.EqualValues(t, _SET_VARIOSELEMS[0], iter.VerActual())
}

func TestIteradorExternoEliminarPrimero(t *testing.T) {
	/*Remueve el primer elemento de la lista luego de haber creado el iterador y comprueba
	que haya cambiado el primero.*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	iter.Borrar()
	if len(_SET_VARIOSELEMS) == 1 {
		require.True(t, lista.EstaVacia())
	} else {
		require.EqualValues(t, _SET_VARIOSELEMS[1], lista.VerPrimero())
		require.EqualValues(t, _SET_VARIOSELEMS[1], iter.VerActual())
	}
}

func TestIteradorExternoEliminarUltimo(t *testing.T) {
	/*Remueve el ultimo elemento de la lista luego de haber creado el iterador y comprueba
	que haya cambiado el ultimo.*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	for i := 0; i < len(_SET_VARIOSELEMS)-1; i++ {
		iter.Siguiente()
	}
	iter.Borrar()
	if len(_SET_VARIOSELEMS) == 1 {
		require.True(t, lista.EstaVacia())
	} else {
		require.EqualValues(t, _SET_VARIOSELEMS[len(_SET_VARIOSELEMS)-2], lista.VerUltimo())
		require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.VerActual() })
	}
}

func TestIteradorExternoEliminarMedio(t *testing.T) {
	/*Remueve el elemento del medio de la lista luego de haber creado el iterador y comprueba
	que haya cambiado el medio.*/
	lista := crearEjemploLista(_SET_VARIOSELEMS)
	iter := lista.Iterador()
	for i := 0; i < len(_SET_VARIOSELEMS)/2; i++ {
		iter.Siguiente()
	}
	elem := iter.Borrar()
	if len(_SET_VARIOSELEMS) == 1 {
		require.True(t, lista.EstaVacia())
	} else {
		require.EqualValues(t, _SET_VARIOSELEMS[len(_SET_VARIOSELEMS)/2], elem)
	}
}
func TestIteradorExternoVolumen(t *testing.T) {
	/* Inserta muchos elementos a la lista y comprueba que hayan sido instertados correctamente*/
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := 0; i < _TAM; i++ {
		iter.Insertar(i)
		require.EqualValues(t, i, iter.VerActual())
	}
	for i := _TAM - 1; i >= 0; i-- {
		require.EqualValues(t, i, iter.Borrar())
	}
	require.True(t, lista.EstaVacia())
}

func TestIteradorInternoVacio(t *testing.T) {
	/*Comprueba que al hacer operaciones invalidas arroje panic*/
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, _ERROR_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, _ERROR_MSG, func() { lista.VerUltimo() })
	require.True(t, lista.EstaVacia())
}

func TestIteradorInternoNumerosPares(t *testing.T) {
	/*Comprueba que el iterador interno funcione correctamente al pasarle una funcion
	que suma los numeros pares*/
	lista := crearEjemploLista(_SET_ENTEROS)
	suma := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
		}
		return true
	})
	require.EqualValues(t, _SET_ENTEROS[0]+_SET_ENTEROS[2], suma)
}

func TestIteradorInternoSinCondicionDeCorte(t *testing.T) {
	/*Comprueba que el iterador interno recorra toda la lista al no haber condicion
	de corte*/
	lista := crearEjemploLista(_SET_ENTEROS)
	contador := 0
	lista.Iterar(func(v int) bool {
		contador++
		return true
	})
	require.EqualValues(t, lista.Largo(), contador)
}

func TestIteradorInternoCondicionDeCorte(t *testing.T) {
	/*Comprueba que el iterador interno recorra hasta la posicion 2 al haber condicion
	de corte*/
	lista := crearEjemploLista(_SET_ENTEROS)
	contador := 0
	lista.Iterar(func(v int) bool {
		contador++
		return contador < 2
	})
	require.EqualValues(t, 2, contador)
}

func crearEjemploLista[T any](arr []T) TDAlista.Lista[T] {
	// Crea un ejemplo de lista genérico
	lista := TDAlista.CrearListaEnlazada[T]()
	for _, valor := range arr {
		lista.InsertarUltimo(valor)
	}
	return lista
}
