package result

type Result[T any] struct {
	Value T
	Err   error
}

func Make[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Err: err}
}

func Wrap[T, U any](f func(...T) (U, error)) func(...T) Result[U] {
	return func(v ...T) Result[U] {
		return Make(f(v...))
	}
}

func Wrap1[T, U any](f func(T) (U, error)) func(T) Result[U] {
	return func(v T) Result[U] {
		return Make(f(v))
	}
}

func Wrap2[T1, T2, U any](f func(T1, T2) (U, error)) func(T1, T2) Result[U] {
	return func(v1 T1, v2 T2) Result[U] {
		return Make(f(v1, v2))
	}
}

func Wrap3[T1, T2, T3, U any](f func(T1, T2, T3) (U, error)) func(T1, T2, T3) Result[U] {
	return func(v1 T1, v2 T2, v3 T3) Result[U] {
		return Make(f(v1, v2, v3))
	}
}

func Wrap4[T1, T2, T3, T4, U any](f func(T1, T2, T3, T4) (U, error)) func(T1, T2, T3, T4) Result[U] {
	return func(v1 T1, v2 T2, v3 T3, v4 T4) Result[U] {
		return Make(f(v1, v2, v3, v4))
	}
}

func Wrap5[T1, T2, T3, T4, T5, U any](f func(T1, T2, T3, T4, T5) (U, error)) func(T1, T2, T3, T4, T5) Result[U] {
	return func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Result[U] {
		return Make(f(v1, v2, v3, v4, v5))
	}
}

func Wrap6[T1, T2, T3, T4, T5, T6, U any](f func(T1, T2, T3, T4, T5, T6) (U, error)) func(T1, T2, T3, T4, T5, T6) Result[U] {
	return func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Result[U] {
		return Make(f(v1, v2, v3, v4, v5, v6))
	}
}
