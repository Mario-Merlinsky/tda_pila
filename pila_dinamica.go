package pila

/* Definición del struct pila proporcionado por la cátedra. */
const TAMANO_INICAL_PILA = 50
const ERROR_PILA_VACIA = "La pila esta vacia"
const CONST_REDIMENSION = 2

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, 0, TAMANO_INICAL_PILA)
	pila.cantidad = 0
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() == true {
		panic(ERROR_PILA_VACIA)
	} else {
		return pila.datos[pila.cantidad-1]
	}
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	pila.datos = redimensionar(pila)
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() == true {
		panic(ERROR_PILA_VACIA)
	} else {
		pila.cantidad--
		pila.datos = redimensionar(pila)
		return pila.datos[pila.cantidad]
	}

}

func redimensionar[T any](pila *pilaDinamica[T]) []T {
	if pila.cantidad == cap(pila.datos) {
		aux := make([]T, CONST_REDIMENSION*cap(pila.datos))
		copy(aux, pila.datos)
		return aux
	} else if pila.cantidad*4 <= cap(pila.datos) && cap(pila.datos) >= TAMANO_INICAL_PILA {
		aux := make([]T, cap(pila.datos)/CONST_REDIMENSION)
		copy(aux, pila.datos)
		return aux
	} else {
		return pila.datos
	}
}
