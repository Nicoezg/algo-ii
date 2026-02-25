package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	/*- Se fija que EstaVacia() de una pila vacia devuelva True
	- Se fija que VerTope() de una pila vacia arroje un panic
	- Se fija que Desapilar() de una pila vacia arroje un panic*/
	t.Log("Hacemos pruebas con una pila vacia.")

	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	t.Log("La pila vacia se comporta como es esperado!")
}

func TestPilaUnicoElementoEntero(t *testing.T) {
	/*- Prueba apilando un entero a una pila vacia.
	- Se fija que una vez apilado el entero la pila no este vacia.
	- Se fija que al VerTope() devuelva el uno, y Desapilar() efectivamente desapila y devuelve el 1.
	- Comprueba que la pila desapilada se comporte como una pila vacia.*/
	t.Log("Hacemos pruebas con un unico elemento entero.")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 1, pila.VerTope())
	require.EqualValues(t, 1, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

	t.Log("La pila con un unico elemento entero se comporta como es esperado!")
}

func TestPilaUnicoElementoString(t *testing.T) {
	/*- Prueba apilando un string a una pila vacia.
	- Se fija que una vez apilado el string la pila no este vacia.
	- Se fija que al VerTope() devuelva "hola", y Desapilar() efectivamente desapila y devuelve el "hola".
	- Comprueba que la pila desapilada se comporte como una pila vacia.*/
	t.Log("Hacemos pruebas con un unico elemento string.")

	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "hola", pila.VerTope())
	require.EqualValues(t, "hola", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

	t.Log("La pila con unico elemento string se comporta como es esperado!")
}

func TestPilaUnicoElementoFloat(t *testing.T) {
	/*- Prueba apilando un float a una pila vacia.
	- Se fija que una vez apilado el float la pila no este vacia.
	- Se fija que al VerTope() devuelva 2.43, y Desapilar() efectivamente desapila y devuelve 2.43.
	- Comprueba que la pila desapilada se comporte como una pila vacia.*/
	t.Log("Hacemos pruebas con un unico elemento float.")

	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(2.43)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 2.43, pila.VerTope())
	require.EqualValues(t, 2.43, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

	t.Log("La pila con unico elemento float se comporta como es esperado!")
}

func TestPilaConElementosDiferentes(t *testing.T) {
	/* - Apila varios enteros, desapilando una vez en el medio, comprobando que el tope sea el correcto
	- Una vez apilados, desapila todos los elementos y se fija si esta vacia.*/
	t.Log("Ahora, probamos con una pila apliando y desapilando varios elementos enteros.")

	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(4)
	pila.Apilar(5)
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

	t.Log("Funciona!")
}

func TestVolumen(t *testing.T) {
	/* - Realiza una prueba con 10000 enteros, apilandolos, luego desapilandolos todos y
	se fija que una vez desapilados todos, la pila este vacia*/
	t.Log("Probamos apilando y desapilando 10000 elementos enteros.")

	tam := 10000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for i := tam - 1; i > -1; i-- {
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())

	t.Log("Exito!")
}

func TestDesapilarHastaVacia(t *testing.T) {
	/* - Apila 10 enteros y los desapila.
	- Se fija que una vez desapilados los elementos, se comporte como una pila vacia.
	- Apila un elemento a la pila vacia para comprobar que su funcionamiento es correcto.*/
	t.Log("Apilamos 10 elementos enteros y luego los desapilamos.")

	tam := 10
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < tam; i++ {
		pila.Apilar(i)
	}
	for i := 9; i > -1; i-- {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.Desapilar())

	t.Log("Se comporta como se esperaba!")
}
