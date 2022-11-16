package cola_prioridad

const (
	_TAMAÑO_HEAP_MIN = 5
	_AUMENTAR_CAP    = 2.0
	_REDUCIR_CAP     = 0.5
	_CAP_MULTIPLO    = 4
	_INICIO_ARREGLO  = 0
	_COMPARACION     = 0
)

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func crearArreglo[T comparable](tamaño int) []T { return make([]T, tamaño) }

func CrearHeap[T comparable](func_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{crearArreglo[T](_TAMAÑO_HEAP_MIN), _INICIO_ARREGLO, func_cmp}
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		heap := CrearHeap(funcion_cmp)
		return heap
	}
	nuevo := crearArreglo[T](len(arreglo) * _AUMENTAR_CAP)
	copy(nuevo, arreglo)
	heap := &heap[T]{nuevo, len(arreglo), funcion_cmp}
	heapify(heap)
	return heap
}

func heapify[T comparable](heap *heap[T]) {
	arreglo_sin_hojas := heap.cantidad / 2
	for j := arreglo_sin_hojas; j >= 0; j-- {
		downHeap(heap.datos, j, heap.cantidad, heap.cmp)
	}
}

func (h heap[T]) EstaVacia() bool { return h.cantidad == _INICIO_ARREGLO }

func (h *heap[T]) Encolar(dato T) {
	if h.cantidad == cap(h.datos) {
		redimensionar(h, _AUMENTAR_CAP)
	}
	h.datos[h.cantidad] = dato
	h.cantidad++
	upHeap(h.datos, h.cantidad-1, h.cmp)
}

func (h heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := h.datos[0]
	h.datos[0] = h.datos[h.cantidad-1]
	h.cantidad--
	downHeap(h.datos, 0, h.cantidad, h.cmp)
	if h.cantidad > _TAMAÑO_HEAP_MIN && h.cantidad*_CAP_MULTIPLO <= cap(h.datos) {
		redimensionar(h, _REDUCIR_CAP)
	}
	return dato
}

func (h heap[T]) Cantidad() int { return h.cantidad }

// funciones auxiliares //
func upHeap[T comparable](datos []T, i int, cmp func(T, T) int) {
	if i == 0 {
		return
	}
	i_padre := (i - 1) / 2
	if i_padre >= 0 {
		if cmp(datos[i], datos[i_padre]) > _COMPARACION {
			swap(datos, i, i_padre)
			upHeap(datos, i_padre, cmp)
		}
	}
}

func downHeap[T comparable](datos []T, i int, cantidad int, cmp func(T, T) int) {
	i_hijo_izq := 2*i + 1
	i_hijo_der := 2*i + 2
	// veo que hijo es mayor
	hijo_mayor := mayor(datos, i_hijo_izq, i_hijo_der, cantidad, cmp)

	// si el hijo mayor es mayor que el padre, hago swap
	if hijo_mayor < cantidad && cmp(datos[hijo_mayor], datos[i]) > _COMPARACION {
		swap(datos, i, hijo_mayor)
		downHeap(datos, hijo_mayor, cantidad, cmp)
	}
}

func mayor[T comparable](datos []T, i_hijo_izq int, i_hijo_der int, cantidad int, cmp func(T, T) int) int {
	if i_hijo_izq >= cantidad && i_hijo_der >= cantidad {
		return cantidad
	}
	if i_hijo_izq >= cantidad {
		return i_hijo_der
	}
	if i_hijo_der >= cantidad {
		return i_hijo_izq
	}
	if cmp(datos[i_hijo_izq], datos[i_hijo_der]) > _COMPARACION {
		return i_hijo_izq
	}
	return i_hijo_der
}

func swap[T comparable](datos []T, i, j int) { datos[i], datos[j] = datos[j], datos[i] }

func redimensionar[T comparable](h *heap[T], factor float64) {
	nuevos_datos := crearArreglo[T](int(float64(cap(h.datos)) * factor))
	copy(nuevos_datos, h.datos)
	h.datos = nuevos_datos
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heap := CrearHeapArr(elementos, funcion_cmp)
	for i := len(elementos) - 1; i >= 0; i-- {
		elementos[i] = heap.Desencolar()
	}
}
