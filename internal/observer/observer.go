package observer

import "sync"

type Publisher[T any] interface {
	Publish(value T)
}

type Subscriber[T any] interface {
	Subscribe() <-chan T
}

type Observer[T any] interface {
	Publisher[T]
	Subscriber[T]
}

type observer[T any] struct {
	mu   sync.Mutex
	subs []chan T
}

func New[T any]() Observer[T] {
	return &observer[T]{
		subs: []chan T{},
	}
}

func (n *observer[T]) Publish(value T) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, ch := range n.subs {
		ch <- value
	}
}

func (n *observer[T]) Subscribe() <-chan T {
	n.mu.Lock()
	defer n.mu.Unlock()

	ch := make(chan T)
	n.subs = append(n.subs, ch)
	return ch
}
