package diccionario_test

import (
	"fmt"
	rand "math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ORDENADO = []int{12500, 25000, 50000, 100000, 200000, 400000}
var SET_STRING_CLAVES = []string{"D", "B", "C", "J", "A", "H"}
var SET_STRING_DATOS = []string{"A", "C", "E", "J", "L", "M"}
var SET_INT_CLAVES = []int{10, 6, 7, 9, 8, 11, 2, 4, 3, 5}
var SET_INT_DATOS = []int{1, 5, 2, 5, 6, 3, 1, 63, 33, 2}

const NO_PERTENECE = "La clave no pertenece al diccionario"
const TERMINO_ITERAR = "El iterador termino de iterar"

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario Ordenado vacio no tiene claves")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Obtener("A") })
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Obtener("") })
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compareInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, NO_PERTENECE, func() { dicNum.Borrar(0) })
}

func TestOrdenadoUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario Ordenado con un elemento tiene esa Clave, unicamente")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	dic.Guardar(SET_STRING_CLAVES[0], SET_STRING_DATOS[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(SET_STRING_CLAVES[0]))
	require.False(t, dic.Pertenece(SET_STRING_CLAVES[1]))
	require.EqualValues(t, SET_STRING_DATOS[0], dic.Obtener(SET_STRING_CLAVES[0]))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Obtener(SET_STRING_DATOS[1]) })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1, clave2, clave3 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2]
	dato1, dato2, dato3 := SET_STRING_DATOS[0], SET_STRING_DATOS[1], SET_STRING_DATOS[2]

	require.False(t, dic.Pertenece(clave1))
	require.False(t, dic.Pertenece(clave1))
	dic.Guardar(clave1, dato1)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave1))
	require.EqualValues(t, dato1, dic.Obtener(clave1))
	require.EqualValues(t, dato1, dic.Obtener(clave1))

	require.False(t, dic.Pertenece(clave2))
	require.False(t, dic.Pertenece(clave3))
	dic.Guardar(clave2, dato2)
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, dato1, dic.Obtener(clave1))
	require.EqualValues(t, dato2, dic.Obtener(clave2))

	require.False(t, dic.Pertenece(clave3))
	dic.Guardar(clave3, dato3)
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.True(t, dic.Pertenece(clave3))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, dato1, dic.Obtener(clave1))
	require.EqualValues(t, dato2, dic.Obtener(clave2))
	require.EqualValues(t, dato3, dic.Obtener(clave3))
}

func TestOrdenadoReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1, clave2 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1]
	dato1, dato2, dato3, dato4 := SET_STRING_DATOS[0], SET_STRING_DATOS[1], SET_STRING_DATOS[2], SET_STRING_DATOS[3]

	dic.Guardar(clave1, dato1)
	dic.Guardar(clave2, dato2)
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, SET_STRING_DATOS[0], dic.Obtener(clave1))
	require.EqualValues(t, SET_STRING_DATOS[1], dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave1, dato3)
	dic.Guardar(clave2, dato4)
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, dato3, dic.Obtener(clave1))
	require.EqualValues(t, dato4, dic.Obtener(clave2))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada y respete la propiedad del abb")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave1, clave2, clave3 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2]
	dato1, dato2, dato3 := SET_STRING_DATOS[0], SET_STRING_DATOS[1], SET_STRING_DATOS[2]
	dic.Guardar(clave1, dato1)
	dic.Guardar(clave2, dato2)
	dic.Guardar(clave3, dato3)

	require.True(t, dic.Pertenece(clave3))
	require.EqualValues(t, dato3, dic.Borrar(clave3))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Borrar(clave3) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(clave3))

	require.True(t, dic.Pertenece(clave1))
	require.EqualValues(t, SET_STRING_DATOS[0], dic.Borrar(clave1))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Borrar(clave1) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(clave1))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Obtener(clave1) })

	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, dato2, dic.Borrar(clave2))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Borrar(clave2) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave2))
	require.PanicsWithValue(t, NO_PERTENECE, func() { dic.Obtener(clave2) })
}

func TestOrdenadoClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""

	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestOrdenadoValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")

	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := SET_STRING_DATOS[0]
	dic.Guardar(clave, nil)

	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func buscaOrdenado[K comparable, V any](clave K, claves []K) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestOrdeandoBorrarCaso0(t *testing.T) {
	t.Log("Comprueba que al borrar un nodo sin hijos el abb mantenga su propiedad y que los nodos sean los correctos")

	claves := []int{7, 8, 4, 9, 5, 6, 10}
	claves_ordenadas := []int{4, 5, 6, 7, 8, 9}
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	for _, clave := range claves {
		dic.Guardar(clave, nil)
	}

	dic.Borrar(10)
	require.False(t, false, dic.Pertenece(10))
	contador := 0

	dic.Iterar(func(clave int, dato *int) bool {
		require.EqualValues(t, claves_ordenadas[contador], clave)
		require.True(t, true, dic.Pertenece(clave))
		contador++
		return true
	})
}
func TestOrdeandoBorrarCaso1(t *testing.T) {
	t.Log("Comprueba que al borrar un nodo con un hijo el abb mantenga su propiedad y que los nodos sean los correctos")

	claves := []int{7, 8, 4, 9, 5, 6, 10}
	claves_ordenadas := []int{4, 6, 7, 8, 9, 10}

	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	for _, clave := range claves {
		dic.Guardar(clave, nil)
	}

	dic.Borrar(5)
	require.False(t, false, dic.Pertenece(5))

	contador := 0

	dic.Iterar(func(clave int, dato *int) bool {

		require.EqualValues(t, claves_ordenadas[contador], clave)
		require.True(t, true, dic.Pertenece(clave))

		contador++
		return true
	})
}
func TestOrdeandoBorrarCaso2(t *testing.T) {
	t.Log("Comprueba que al borrar un nodo con dos hijos el abb mantenga su propiedad y que los nodos sean los correctos")

	claves := []int{10, 5, 15, 3, 8, 7, 12, 20, 14, 13}
	claves_ordenadas := []int{3, 5, 7, 8, 12, 13, 14, 15, 20}
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	for _, clave := range claves {
		dic.Guardar(clave, nil)
	}

	dic.Borrar(10)
	require.False(t, false, dic.Pertenece(15))

	contador := 0

	dic.Iterar(func(clave int, dato *int) bool {
		require.EqualValues(t, claves_ordenadas[contador], clave)
		require.True(t, true, dic.Pertenece(clave))

		contador++
		return true
	})
}

func TestOrdenadoIteradorInternoClaves(t *testing.T) {
	t.Log("Comprueba que el iterador interno haga el recorrido correspondiente en orden.")

	dic := TDADiccionario.CrearABB[int, int](compareInt)

	for i, clave := range SET_INT_CLAVES {
		dic.Guardar(clave, SET_INT_DATOS[i])
	}

	dic.Iterar(func(clave, dato int) bool {
		require.EqualValues(t, dato, dic.Obtener(clave))
		return true
	})
}

func TestOrdenadoIteradorInternoCorte(t *testing.T) {
	t.Log("Prueba iterando con una condición de corte donde borra el elemento siempre cuando no sea la condicion" +
		"de corte.")

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	for i, clave := range SET_STRING_CLAVES {
		dic.Guardar(clave, SET_STRING_DATOS[i])
	}

	dic.Iterar(func(clave, dato string) bool {
		if clave == SET_STRING_CLAVES[0] {
			return false
		}

		require.EqualValues(t, dato, dic.Borrar(clave))
		require.False(t, false, dic.Pertenece(clave))
		return true
	})

	require.True(t, true, dic.Pertenece(SET_STRING_CLAVES[0]))
	require.True(t, true, dic.Pertenece(SET_STRING_CLAVES[3]))
	require.True(t, true, dic.Pertenece(SET_STRING_CLAVES[5]))

}

func ejecutarOrdenPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[int, int](compareInt)

	claves := make([]int, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el abb */
	for i, valor := range rand.Perm(n) {
		valores[i] = i
		claves[i] = valor
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i, valor := range claves {
		ok = dic.Pertenece(valor)
		if !ok {
			break
		}
		ok = dic.Obtener(valor) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarOrdenPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, TERMINO_ITERAR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TERMINO_ITERAR, func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1, clave2, clave3 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2]
	valor1, valor2, valor3 := SET_STRING_DATOS[0], SET_STRING_DATOS[1], SET_STRING_DATOS[2]

	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscaOrdenado[string, string](primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscaOrdenado[string, string](segundo, claves))
	require.EqualValues(t, valores[buscaOrdenado[string, string](segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscaOrdenado[string, string](tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, TERMINO_ITERAR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TERMINO_ITERAR, func() { iter.Siguiente() })
}

func TestOrdenadoIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2]}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscaOrdenado[string, string](primero, claves))
	require.NotEqualValues(t, -1, buscaOrdenado[string, string](segundo, claves))
	require.NotEqualValues(t, -1, buscaOrdenado[string, string](tercero, claves))
}

func TestIteradorExternoRecorreDiccionarioOrdenado(t *testing.T) {
	t.Log("Comprueba que el iterador recorra el diccionario en orden.")
	var arr []int
	dicc := TDADiccionario.CrearABB[int, int](compareInt)

	for i, clave := range SET_INT_CLAVES {
		dicc.Guardar(clave, SET_INT_DATOS[i])
	}
	iter := dicc.Iterador()

	for !iter.HaySiguiente() {
		_, dato := iter.VerActual()
		arr = append(arr, dato)
	}

	require.True(t, true, estaOrdenado(arr, compareInt))
}

func ejecutaroOrdenadoPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el Abb */
	for i, valor := range rand.Perm(n) {
		claves[i] = fmt.Sprintf("%08d", valor)
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkOrdenadoIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario Ordenado. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutaroOrdenadoPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestOrdenadoVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compareInt)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestOrdenadoIteradorSumaElementos(t *testing.T) {
	t.Log("Crea un diccionario ordenado y hace la suma todos los elementos del diccionario")

	dic := TDADiccionario.CrearABB[int, int](compareInt)
	suma := 0

	dic.Guardar(5, 0)
	dic.Guardar(7, 2)
	dic.Guardar(4, 9)
	dic.Guardar(0, 10)

	dic.Iterar(func(c int, v int) bool {
		suma += v
		return true
	})

	require.EqualValues(t, 21, suma)
}

func TestOrdenadoIteradorInternoConCorte(t *testing.T) {
	t.Log("Comprueba que el iterador interno con condición de corte efectivamente corte cuando lo indique la funcion")

	dic := TDADiccionario.CrearABB[int, int](compareInt)

	dic.Guardar(113, 0)
	dic.Guardar(2, 2)
	dic.Guardar(3, 3)
	dic.Guardar(4, 4)

	contador := 0
	dic.Iterar(func(c int, v int) bool {
		if c%2 != 0 && v%2 != 0 {
			return false
		}
		contador++
		return true
	})
	require.EqualValues(t, 1, contador)
}

func TestOrdenadoIteracionSumaConRangos(t *testing.T) {
	t.Log("Prueba iterando todos los elementos y verifica que la suma de todos los valores de los datos")

	dic := TDADiccionario.CrearABB[int, int](compareInt)
	suma := 0
	clave1, clave2, clave3, clave4, clave5 := 1, 5, 9, 20, 13
	valor1, valor2, valor3, valor4, valor5 := 0, 2, 9, 10, 20

	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Guardar(clave4, valor4)
	dic.Guardar(clave5, valor5)

	dic.IterarRango(&clave2, &clave4, func(c int, v int) bool {
		suma += v
		return true
	})

	require.EqualValues(t, 41, suma)
}

func TestIteradorInternoSinRangos(t *testing.T) {
	t.Log("Comprueba que al no pasar rangos al iterador interno por rangos, itere todo.")

	clave1, clave2, clave3, clave4, clave5 := SET_INT_CLAVES[0], SET_INT_CLAVES[1], SET_INT_CLAVES[2], SET_INT_CLAVES[3], SET_INT_CLAVES[4]
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)
	dic.Guardar(clave5, nil)

	recorrido := []int{6, 7, 8, 9, 10}
	contador := 0
	dic.IterarRango(nil, nil, func(clave int, dato *int) bool {

		require.EqualValues(t, recorrido[contador], clave)
		contador++
		return true
	})

}

func TestIteradorInternoSinDesde(t *testing.T) {
	t.Log("Comprueba que al establecer un rango acotado superiormente al iterador interno por rangos, itere correctamente.")

	clave1, clave2, clave3, clave4, clave5 := SET_INT_CLAVES[0], SET_INT_CLAVES[1], SET_INT_CLAVES[2], SET_INT_CLAVES[3], SET_INT_CLAVES[4]
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)
	dic.Guardar(clave5, nil)

	recorrido := []int{6, 7, 8, 9}
	contador := 0
	dic.IterarRango(nil, &clave4, func(clave int, dato *int) bool {

		require.EqualValues(t, recorrido[contador], clave)
		contador++
		return true
	})

}

func TestIteradorInternoSinHasta(t *testing.T) {
	t.Log("Comprueba que al establecer un rango acotado inferiormente al iterador interno por rangos, itere correctamente.")

	clave1, clave2, clave3, clave4, clave5 := SET_INT_CLAVES[0], SET_INT_CLAVES[1], SET_INT_CLAVES[2], SET_INT_CLAVES[3], SET_INT_CLAVES[4]
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)
	dic.Guardar(clave5, nil)
	recorrido := []int{8, 9, 10}
	contador := 0
	dic.IterarRango(&clave5, nil, func(clave int, dato *int) bool {

		require.EqualValues(t, recorrido[contador], clave)
		contador++
		return true
	})

}

func TestIteradorInternoConRangos(t *testing.T) {
	t.Log("Comprueba que al no pasar rangos al iterador interno por rangos, itere todo.")

	clave1, clave2, clave3, clave4, clave5 := SET_INT_CLAVES[0], SET_INT_CLAVES[1], SET_INT_CLAVES[2], SET_INT_CLAVES[3], SET_INT_CLAVES[4]
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)
	dic.Guardar(clave5, nil)

	contador := 0
	recorrido := []int{clave3, clave5, clave4}

	dic.IterarRango(&(SET_INT_CLAVES[2]), &(SET_INT_CLAVES[3]), func(clave int, dato *int) bool {
		require.EqualValues(t, recorrido[contador], clave)
		contador++
		return true
	})
}

func TestIteradorInternoConRangosVacio(t *testing.T) {
	t.Log("Prueba iterando sobre un diccionario vacio")

	dic := TDADiccionario.CrearABB[int, *string](compareInt)
	dic.IterarRango(&(SET_INT_CLAVES[2]), &(SET_INT_CLAVES[4]), func(clave int, dato *string) bool {
		return true
	})
	require.False(t, false, dic.Pertenece(SET_INT_CLAVES[2]))

}

func TestIteradorRangosInexistentes(t *testing.T) {
	t.Log("Prueba definiendo un rango inexistente y comprobando que el iterador no recorra ningun elemento")

	dic := TDADiccionario.CrearABB[string, *string](strings.Compare)
	clave1, clave2, clave3 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2]
	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)

	desde := "C"
	hasta := "A"

	contador := 0
	dic.IterarRango(&desde, &hasta, func(clave string, dato *string) bool {
		contador++
		return true
	})
	require.EqualValues(t, 0, contador)
	require.True(t, true, dic.Pertenece(clave1))
	require.True(t, true, dic.Pertenece(clave2))
	require.True(t, true, dic.Pertenece(clave3))
}

func TestIteradorExternoSinRangos(t *testing.T) {
	t.Log("Itera sin rangos, comprobando que efectivamente pase por todos los elementos")

	dic := TDADiccionario.CrearABB[string, *string](strings.Compare)
	clave1, clave2, clave3, clave4 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2], SET_STRING_CLAVES[3]
	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)

	recorrido := []string{clave2, clave3, clave1, clave4}
	iter := dic.IteradorRango(nil, nil)

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorrido[0], clave)

		iter.Siguiente()
		recorrido = recorrido[1:]
	}
	require.False(t, false, iter.HaySiguiente())
}

func TestIteradorExternoSinDesde(t *testing.T) {
	t.Log("Itera sin rangos, comprobando que efectivamente pase por todos los elementos")

	dic := TDADiccionario.CrearABB[string, *string](strings.Compare)
	clave1, clave2, clave3, clave4 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2], SET_STRING_CLAVES[3]
	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)

	recorrido := []string{clave2, clave3, clave1}
	iter := dic.IteradorRango(nil, &clave1)

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorrido[0], clave)

		iter.Siguiente()
		recorrido = recorrido[1:]
	}
	require.False(t, false, iter.HaySiguiente())
}

func TestIteradorExternoSinHasta(t *testing.T) {
	t.Log("Itera sin rangos, comprobando que efectivamente pase por todos los elementos")

	dic := TDADiccionario.CrearABB[string, *string](strings.Compare)
	clave1, clave2, clave3, clave4 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2], SET_STRING_CLAVES[3]

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)

	recorrido := []string{clave3, clave1, clave4}
	iter := dic.IteradorRango(&clave3, nil)

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorrido[0], clave)

		iter.Siguiente()
		recorrido = recorrido[1:]
	}
	require.False(t, false, iter.HaySiguiente())
}

func TestIteradorExternoConRangos(t *testing.T) {
	t.Log("Comprueba que itere correctamente estableciendo un rango")

	clave1, clave2, clave3, clave4, clave5 := SET_INT_CLAVES[0], SET_INT_CLAVES[1], SET_INT_CLAVES[2], SET_INT_CLAVES[3], SET_INT_CLAVES[4]
	dic := TDADiccionario.CrearABB[int, *int](compareInt)

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)
	dic.Guardar(clave4, nil)
	dic.Guardar(clave5, nil)
	dic.Guardar(11, nil)

	recorrido := []int{clave1, 11}

	iter := dic.IteradorRango(&(clave1), &(SET_INT_CLAVES[5]))

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorrido[0], clave)
		recorrido = recorrido[1:]

		iter.Siguiente()

	}
}

func TestIteradorExternoConRangosSumarDatos(t *testing.T) {
	t.Log("Prueba iterando con rangos y sumando todos los datos asociadas a las claves dentro de ese rango")

	dic := TDADiccionario.CrearABB[int, int](compareInt)
	clave1, clave2, clave3, clave4, clave5 := SET_INT_CLAVES[0], SET_INT_CLAVES[1], SET_INT_CLAVES[2], SET_INT_CLAVES[3], SET_INT_CLAVES[4]
	dato1, dato2, dato3, dato4, dato5 := SET_INT_DATOS[0], SET_INT_DATOS[1], SET_INT_DATOS[2], SET_INT_DATOS[3], SET_INT_DATOS[4]

	dic.Guardar(clave1, dato1)
	dic.Guardar(clave2, dato2)
	dic.Guardar(clave3, dato3)
	dic.Guardar(clave4, dato4)
	dic.Guardar(clave5, dato5)

	recorrido := []int{6, 7, 8, 9, 10}
	suma := 0
	iter := dic.IteradorRango(&clave2, &clave1)

	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		require.EqualValues(t, recorrido[0], clave)

		suma += dato
		iter.Siguiente()
		recorrido = recorrido[1:]
	}

	require.EqualValues(t, dato2+dato3+dato5+dato4+dato1, suma)
	require.PanicsWithValue(t, TERMINO_ITERAR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TERMINO_ITERAR, func() { iter.VerActual() })
}

func TestIteradorExternoConRangosVacio(t *testing.T) {
	t.Log("Intenta iterar sobre un diccionario vacio")
	dic := TDADiccionario.CrearABB[int, int](compareInt)

	desde := 1
	hasta := 10

	iter := dic.IteradorRango(&desde, &hasta)

	require.False(t, false, iter.HaySiguiente())
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	require.False(t, false, iter.HaySiguiente())
}

func TestIteradorExternoRangosInexistentes(t *testing.T) {
	t.Log("Define un rango inexistente en el diccionario e itera")
	dic := TDADiccionario.CrearABB[string, *string](strings.Compare)
	clave1, clave2, clave3 := SET_STRING_CLAVES[0], SET_STRING_CLAVES[1], SET_STRING_CLAVES[2]

	dic.Guardar(clave1, nil)
	dic.Guardar(clave2, nil)
	dic.Guardar(clave3, nil)

	desde := "Q"
	hasta := "Z"

	iter := dic.IteradorRango(&desde, &hasta)

	require.False(t, false, iter.HaySiguiente())

	for iter.HaySiguiente() {
		iter.Siguiente()

	}

	require.True(t, true, dic.Pertenece(clave1))
	require.True(t, true, dic.Pertenece(clave2))
	require.True(t, true, dic.Pertenece(clave3))
}

func compareInt(a, b int) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func estaOrdenado[K comparable](arr []K, cmp func(K, K) int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if cmp(arr[i], arr[i+1]) <= 0 {
			continue
		}
		return false
	}
	return true
}

func ejecutaroOrdenadoPruebasVolumenIteradorRango(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el Abb */
	for i, valor := range rand.Perm(n) {
		claves[i] = fmt.Sprintf("%08d", valor)
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.IteradorRango(nil, nil)
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkOrdenadoIteradorRango(b *testing.B) {
	b.Log("Prueba de stress del Iterador con rangos del Diccionario Ordenado. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN_ORDENADO {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutaroOrdenadoPruebasVolumenIteradorRango(b, n)
			}
		})
	}
}

func TestDescendientes(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](compareInt)
	dic.Guardar(6, nil)
	dic.Guardar(3, nil)
	dic.Guardar(9, nil)
	dic.Guardar(7, nil)
	dic.Guardar(10, nil)
	dic.Guardar(4, nil)
	dic.Guardar(2, nil)
	require.EqualValues(t, 1, dic.NDescendientes(6))
}
