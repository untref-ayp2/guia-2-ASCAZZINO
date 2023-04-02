package queue

import (
	"errors"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	elementos []any //se encapusula y lo hhace privado

}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	(*q).elementos = append((*q).elementos, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	//al valor que esta en la direccion de memoria Q(slice), contruime otro slice nuevo que guarde en q* que lleva en la possicion 1 hasta la n - 1 y reemplazo por un nuevo valor
	//El [1:] es mirar el valor desde ahi en adelante
	head := (*q)[0]
	*q = (*q)[1:] //Solo lo mira y no lo bora
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*q)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
// SI O SI se necesita para acceder a la pila y pregunte si esta vacia en caso de que la PILA sea privada
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type QueueS struct {
	pila1 stack.Stack
	pila2 stack.Stack
}

func (q *QueueS) Enqueue(v any) {
	if q.pila2.IsEmpty() || q.pila2.IsEmpty() {
		q.pila1.Push(v)
	} else { //si pila dos no esta vacio
		for !q.pila2.IsEmpty() {
			aux, _ := q.pila2.pop()
			q.pila1.Push(aux)
		}
		q.pila1.Push(v)
	}
}

func (q *QueueS) Dequeue() (any, error) {
	if len((q.pila2)) == 0 {
		for len(q.pila1) > 0 {
			// variable aux para extraer el primer elemento del array
			aux := q.pila1[len(q.pila1)]
			q.pila1 = q.pila1[:len(q.pila1)]

			// Agregar el elemento a la pila pila2
			q.pila2 = append(q.pila2, aux)
		}
	}
}

func (q *QueueS) IsEmpty() bool {
	return q.pila1.IsEmpty() && q.pila2.IsEmpty()
}

func (q *QueueS) Front() (any, error) {
	//TODO
}
