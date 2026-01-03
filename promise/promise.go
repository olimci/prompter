package promise

import "context"

type result[T any] struct {
	value T
	err   error
}

func New[T any]() (Promise[T], Resolver[T]) {
	ch := make(chan result[T])
	return Promise[T]{ch}, Resolver[T](ch)
}

type Promise[T any] struct {
	ch <-chan result[T]
}

func (p Promise[T]) Await(ctx context.Context) (T, error) {
	select {
	case res := <-p.ch:
		return res.value, res.err
	case <-ctx.Done():
		return *new(T), ctx.Err()
	}
}

type Resolver[T any] chan<- result[T]

func (r Resolver[T]) resolve(res result[T]) {
	select {
	case r <- res:
	default:
	}
}

func (r Resolver[T]) Success(value T) {
	r.resolve(result[T]{value: value})
}

func (r Resolver[T]) Failure(err error) {
	r.resolve(result[T]{err: err})
}
