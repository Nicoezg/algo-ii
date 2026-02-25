package cola_prioridad_test

import (
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var SET_ENTEROS = []int{1, 5, 9, 4, 2}
var SET_STRINGS = []string{"A", "D", "B", "H", "W"}
var SET_COMPLEJOS = []complejo{{5, 2.5}, {5, 3.5}, {0, 0.5}, {0, 0.6}, {0, 0.5}}

const COLA_VACIA = "La cola esta vacia"
const TAM = 10000

func TestColaPrioridadVacia(t *testing.T) {
	t.Log("Comprueba que una cola de prioridad vacia no tenga elementos")

	cola := TDAColaPrioridad.CrearHeap[string](strings.Compare)
	require.EqualValues(t, 0, cola.Cantidad())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, COLA_VACIA, func() { cola.VerMax() })
	require.PanicsWithValue(t, COLA_VACIA, func() { cola.Desencolar() })
}

func TestUnicoElemento(t *testing.T) {
	t.Log("Prueba insertando un elemento a la cola de prioridad")

	//heap := TDAColaPrioridad.CrearHeapArr[int](SET_ENTEROS, compareInt)
	cola := TDAColaPrioridad.CrearHeap[int](compareInt)
	cola.Encolar(SET_ENTEROS[0])
	require.EqualValues(t, 1, cola.Cantidad())
	require.EqualValues(t, SET_ENTEROS[0], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[0], cola.Desencolar())

	require.True(t, true, cola.EstaVacia())
	require.EqualValues(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, COLA_VACIA, func() { cola.VerMax() })
	require.PanicsWithValue(t, COLA_VACIA, func() { cola.Desencolar() })
}

func TestColaPrioridadEnteros(t *testing.T) {
	t.Log("Corrobora que la cola de prioridad respete su propiedad")

	cola := TDAColaPrioridad.CrearHeap[int](compareInt)

	cola.Encolar(SET_ENTEROS[2])
	cola.Encolar(SET_ENTEROS[3])
	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_ENTEROS[2], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[2], cola.Desencolar())

	cola.Encolar(SET_ENTEROS[1])

	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_ENTEROS[1], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[1], cola.Desencolar())

	require.EqualValues(t, SET_ENTEROS[3], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[3], cola.Desencolar())

	require.EqualValues(t, 0, cola.Cantidad())

	cola.Encolar(SET_ENTEROS[0])
	cola.Encolar(SET_ENTEROS[4])
	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_ENTEROS[4], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[4], cola.Desencolar())

	require.EqualValues(t, SET_ENTEROS[0], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[0], cola.Desencolar())

	require.True(t, true, cola.EstaVacia())
	require.EqualValues(t, 0, cola.Cantidad())

}

func TestColaPrioridadStrings(t *testing.T) {
	t.Log("Corrobora que la cola de prioridad respete su propiedad con strings tambien")

	cola := TDAColaPrioridad.CrearHeap[string](strings.Compare)

	cola.Encolar(SET_STRINGS[4])
	cola.Encolar(SET_STRINGS[0])
	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_STRINGS[4], cola.VerMax())
	require.EqualValues(t, SET_STRINGS[4], cola.Desencolar())
	cola.Encolar(SET_STRINGS[1])

	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_STRINGS[1], cola.VerMax())
	require.EqualValues(t, SET_STRINGS[1], cola.Desencolar())

	require.EqualValues(t, SET_STRINGS[0], cola.VerMax())
	require.EqualValues(t, SET_STRINGS[0], cola.Desencolar())

	require.EqualValues(t, 0, cola.Cantidad())

	cola.Encolar(SET_STRINGS[2])
	cola.Encolar(SET_STRINGS[3])

	require.EqualValues(t, SET_STRINGS[3], cola.VerMax())
	require.EqualValues(t, SET_STRINGS[3], cola.Desencolar())

	require.EqualValues(t, SET_STRINGS[2], cola.VerMax())
	require.EqualValues(t, SET_STRINGS[2], cola.Desencolar())

	require.True(t, true, cola.EstaVacia())
	require.EqualValues(t, 0, cola.Cantidad())

}

type complejo struct {
	prioridad1 int
	prioridad2 float64
}

func TestEstructurasComplejas(t *testing.T) {
	t.Log("Corrobora que la cola de prioridad respete su propiedad con una estructura mas compleja")

	cola := TDAColaPrioridad.CrearHeap[complejo](compareComplejo)

	cola.Encolar(SET_COMPLEJOS[0])
	cola.Encolar(SET_COMPLEJOS[1])
	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_COMPLEJOS[1], cola.VerMax())
	require.EqualValues(t, SET_COMPLEJOS[1], cola.Desencolar())
	cola.Encolar(SET_COMPLEJOS[3])

	require.EqualValues(t, 2, cola.Cantidad())

	require.EqualValues(t, SET_COMPLEJOS[0], cola.VerMax())
	require.EqualValues(t, SET_COMPLEJOS[0], cola.Desencolar())

	require.EqualValues(t, SET_COMPLEJOS[3], cola.VerMax())
	require.EqualValues(t, SET_COMPLEJOS[3], cola.Desencolar())

	require.EqualValues(t, 0, cola.Cantidad())

	cola.Encolar(SET_COMPLEJOS[2])
	cola.Encolar(SET_COMPLEJOS[4])

	require.EqualValues(t, SET_COMPLEJOS[2], cola.VerMax())
	require.EqualValues(t, SET_COMPLEJOS[2], cola.Desencolar())

	require.EqualValues(t, SET_COMPLEJOS[4], cola.VerMax())
	require.EqualValues(t, SET_COMPLEJOS[4], cola.Desencolar())

	require.True(t, true, cola.EstaVacia())
	require.EqualValues(t, 0, cola.Cantidad())

}

func TestCrearHeapSobreArreglo(t *testing.T) {
	t.Log("Comprueba que al crear un heap sobre un arreglo funcione correctamente")

	cola := TDAColaPrioridad.CrearHeapArr[int](SET_ENTEROS, compareInt)
	require.EqualValues(t, len(SET_ENTEROS), cola.Cantidad())

	require.EqualValues(t, SET_ENTEROS[2], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[2], cola.Desencolar())

	require.EqualValues(t, SET_ENTEROS[1], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[1], cola.Desencolar())

	require.EqualValues(t, SET_ENTEROS[3], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[3], cola.Desencolar())

	require.EqualValues(t, SET_ENTEROS[4], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[4], cola.Desencolar())

	require.EqualValues(t, SET_ENTEROS[0], cola.VerMax())
	require.EqualValues(t, SET_ENTEROS[0], cola.Desencolar())

	require.True(t, true, cola.EstaVacia())

}

func TestCrearHeapSobreArregloVacio(t *testing.T) {
	t.Log("Comprueba que al crear un heap sobre un arreglo vacio se comporte como tal")

	cola := TDAColaPrioridad.CrearHeapArr[int]([]int{}, compareInt)

	require.True(t, true, cola.EstaVacia())
	require.EqualValues(t, 0, cola.Cantidad())
}

func TestHeapSortEnteros(t *testing.T) {
	t.Log("Verifica que HeapSort ordene correctamente los enteros")

	TDAColaPrioridad.HeapSort(SET_ENTEROS, compareInt)

	res := []int{1, 2, 4, 5, 9}

	for i, dato := range SET_ENTEROS {
		require.EqualValues(t, res[i], dato)
	}
}

func TestHeapSortString(t *testing.T) {
	t.Log("Comprueba que HeapSort ordene correctamente un arreglo de strings")

	TDAColaPrioridad.HeapSort(SET_STRINGS, strings.Compare)

	res := []string{"A", "B", "D", "H", "W"}

	for i, dato := range SET_STRINGS {
		require.EqualValues(t, res[i], dato)
	}

}

func TestHeapSortComplejos(t *testing.T) {
	t.Log("Comprueba que HeapSort ordene correctamente un arreglo con tipos de datos mas complejos")

	TDAColaPrioridad.HeapSort(SET_COMPLEJOS, compareComplejo)

	res := []complejo{{0, 0.5}, {0, 0.5}, {0, 0.6}, {5, 2.5}, {5, 3.5}}

	for i, dato := range SET_COMPLEJOS {
		require.EqualValues(t, res[i], dato)
	}

}

func TestVolumen(t *testing.T) {
	t.Log("Realiza un test de volumen encolando muchos enteros")

	cola := TDAColaPrioridad.CrearHeap[int](compareInt)
	for i := 0; i < TAM; i++ {
		cola.Encolar(i)
	}
	for i := TAM - 1; i >= 0; i-- {
		require.EqualValues(t, i, cola.Desencolar())

	}
	require.True(t, cola.EstaVacia())

}

func TestVolumenHeapsort(t *testing.T) {
	t.Log("Realiza un test de volumen de heapsort")

	arr := make([]int, 15000)
	for i := 0; i < 15000; i++ {
		arr[i] = i
	}
	TDAColaPrioridad.HeapSort(arr, compareInt)

	for i, valor := range arr {
		require.EqualValues(t, i, valor)
	}
}

func compareInt(a, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}

func compareComplejo(a, b complejo) int {
	if a.prioridad1 > b.prioridad1 {
		return 1
	}

	if a.prioridad1 == b.prioridad1 {

		if a.prioridad2 > b.prioridad2 {
			return 1
		}

		if a.prioridad2 == b.prioridad2 {
			return 0
		}

		return -1

	}
	return -1
}
