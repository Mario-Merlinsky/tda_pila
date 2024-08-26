package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

// se crea una pila y se verifica que se comporta como recien creada
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

}

// se apilan elementos y se desapilan verificando que los topes son correctos
func TestApilaDesapila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Desapilar()
	require.False(t, pila.EstaVacia())
	require.Equal(t, 1, pila.VerTope())
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Desapilar()
	pila.Desapilar()
	require.Equal(t, 3, pila.VerTope())

}

// se apilan muchos elementos y luego se desapilan, verificando que los topes sean correctos y que cuando quede vacia, se comporte como recien creada
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10001; i++ {
		pila.Apilar(i)
	}
	require.Equal(t, 10000, pila.VerTope())
	for i := 0; i < 10000; i++ {
		pila.Desapilar()
	}

	require.Equal(t, 0, pila.VerTope())
	pila.Desapilar()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// se apilan elementos y luego se desapilan hasta vaciarla, verificando que se comporte como recien creada
func TestDesapilarHastaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}
	for i := 0; i < 10; i++ {
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// se apilan elementos enteros en una pila, y strings en otra, verificando que los topes sean correctos
func TestApilarDistintos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i > -51; i-- {
		pila.Apilar(i)
	}
	require.Equal(t, -50, pila.VerTope())

	pila_2 := TDAPila.CrearPilaDinamica[string]()
	pila_2.Apilar("Hola")
	pila_2.Apilar("Como")
	pila_2.Apilar("Estas?")
	require.Equal(t, "Estas?", pila_2.VerTope())
	pila_2.Desapilar()
	require.Equal(t, "Como", pila_2.VerTope())
}
