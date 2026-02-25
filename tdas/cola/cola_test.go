package cola_test

import (
	TDAcola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	/*- Se fija que EstaVacia() de una cola vacia devuelva True.
	- Se fija que VerPrimero() de una cola vacia arroje un panic.
	- Se fija que Desencolar() de una cola vacia arroje un panic.*/
	t.Log("Hacemos pruebas con una cola vacia.")

	cola := TDAcola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	t.Log("La cola vacia se comporta como es esperado!")
}

func TestUnicoElementoEntero(t *testing.T) {
	/*- Prueba encolando un entero a una cola vacia.
	- Se fija que una vez encolado el entero la cola no este vacia.
	- Se fija que VerPrimero() devuelva el uno, y Desencolar() efectivamente desencola y devuelve el 1.
	- Comprueba que la cola desencolada se comporte como una cola vacia.*/
	t.Log("Hacemos pruebas con un unico elemento entero.")
	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

	t.Log("La cola con un unico elemento entero se comporta como es esperado!")
}

func TestColaUnicoElementoString(t *testing.T) {
	/*- Prueba encolando un string a una cola vacia.
	- Se fija que una vez encolado el string la cola no este vacia.
	- Se fija que VerPrimero() devuelva "hola", y Desencolar() efectivamente desencola y devuelve el "hola".
	- Comprueba que la cola desencolada se comporte como una cola vacia.*/
	t.Log("Hacemos pruebas con un unico elemento string.")

	cola := TDAcola.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "hola", cola.VerPrimero())
	require.EqualValues(t, "hola", cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

	t.Log("La cola con unico elemento string se comporta como es esperado!")
}

func TestColaUnicoElementoFloat(t *testing.T) {
	/*- Prueba encolando un float a una cola vacia.
	- Se fija que una vez encolado el float la cola no este vacia.
	- Se fija que VerPrimero() devuelva 2.43, y Desencolar() efectivamente desencola y devuelve 2.43.
	- Comprueba que la cola desencolada se comporte como una cola vacia.*/
	t.Log("Hacemos pruebas con un unico elemento float.")

	cola := TDAcola.CrearColaEnlazada[float64]()
	cola.Encolar(2.43)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 2.43, cola.VerPrimero())
	require.EqualValues(t, 2.43, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

	t.Log("La cola con unico elemento float se comporta como es esperado!")
}

func TestColaConElementosDiferentes(t *testing.T) {
	/* - Encola varios enteros, desencolando una vez en el medio, comprobando que se respete el FIFO.
	- Una vez encolados, desencola todos los elementos y se fija si esta vacia.*/
	t.Log("Ahora, probamos con una cola encolando y desencolando varios elementos enteros.")

	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.VerPrimero())
	cola.Encolar(4)
	cola.Encolar(5)
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
	require.EqualValues(t, 4, cola.Desencolar())
	require.EqualValues(t, 5, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

	t.Log("Funciona!")
}

func TestVolumen(t *testing.T) {
	/* - Realiza una prueba con 10000 enteros, encolandolos, luego desencolandolos todos y
	se fija que una vez desencolados todos, la cola este vacia*/
	t.Log("Probamos encolando y desencolando 10000 elementos enteros.")

	tam := 10000
	cola := TDAcola.CrearColaEnlazada[int]()
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())

	t.Log("Exito!")
}

func TestDesencolarHastaVacia(t *testing.T) {
	/* - Encola 10 enteros y los desencola.
	- Se fija que una vez desencolados los elementos, se comporte como una cola vacia.*/
	t.Log("Encolamos 10 elementos enteros y luego los desencolamos.")

	tam := 10
	cola := TDAcola.CrearColaEnlazada[int]()
	for i := 0; i < tam; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < tam; i++ {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	t.Log("Se comporta como se esperaba!")
}
