package jslice

import "errors"

type Iterator[T any, V any] func(elem T, idx int, list []T) V

func At[T any](list []T, idx int) (T, error) {
	var zero T
	if idx < 0 {
		idx = idx + len(list)
	}

	if idx < 0 || idx > len(list)-1 {
		return zero, errors.New("index out of range")
	}

	return list[idx], nil
}

func Concat[T any](l1, l2 []T) []T {
	result := make([]T, 0, len(l1)+len(l2))
	for _, e := range l1 {
		result = append(result, e)
	}

	for _, e := range l2 {
		result = append(result, e)
	}

	return result
}

func Pop[T any](list *[]T) (T, bool) {
	var zero T
	length := len(*list)

	if length == 0 {
		return zero, false
	}

	tail := (*list)[length-1]
	*list = (*list)[:len(*list)-1]

	return tail, true
}

func Push[T any](list *[]T, elems ...T) {
	*list = append(*list, elems...)
}

func Unshift[T any](list *[]T, elems ...T) {
	if len(elems) == 0 {
		return
	}

	*list = Concat(elems, *list)
}

func Shift[T any](list *[]T) (T, bool) {
	var zero T
	if len(*list) == 0 {
		return zero, false
	}

	first := (*list)[0]
	*list = (*list)[1:]

	return first, true
}

func Find[T any](list []T, iterator Iterator[T, bool]) (T, bool) {
	var zero T
	if idx := FindIndex(list, iterator); idx != -1 {
		return list[idx], true
	}

	return zero, false
}

func FindIndex[T any](list []T, iterator Iterator[T, bool]) int {
	for idx, elem := range list {
		found := iterator(elem, idx, list)
		if found {
			return idx
		}
	}

	return -1
}

func FindLast[T any](list []T, iterator Iterator[T, bool]) (T, bool) {
	var zero T
	if idx := FindLastIndex(list, iterator); idx != -1 {
		return list[idx], true
	}

	return zero, false
}

func FindLastIndex[T any](list []T, iterator Iterator[T, bool]) int {
	for i := len(list) - 1; i > -1; i-- {
		if found := iterator(list[i], i, list); found {
			return i
		}
	}

	return -1
}

func Some[T any](list []T, iterator Iterator[T, bool]) bool {
	for idx, elem := range list {
		found := iterator(elem, idx, list)
		if found {
			return true
		}
	}

	return false
}

func Every[T any](list []T, iterator Iterator[T, bool]) bool {
	for idx, elem := range list {
		found := !iterator(elem, idx, list)
		if found {
			return false
		}
	}

	return true
}

func Includes[T comparable](list []T, target T) bool {
	for _, elem := range list {
		if elem == target {
			return true
		}
	}

	return false
}

func Filter[T any](list []T, iterator Iterator[T, bool]) []T {
	result := make([]T, 0, len(list))
	ForEach(list, func(elem T, idx int, l []T) {
		if iterator(elem, idx, l) {
			result = append(result, elem)
		}
	})

	return result
}

func ForEach[T any](list []T, iterator func(elem T, idx int, list []T)) {
	for idx, elem := range list {
		iterator(elem, idx, list)
	}
}

func Map[T any, V any](list []T, iterator Iterator[T, V]) []V {
	result := make([]V, 0, len(list))

	ForEach(list, func(elem T, idx int, list []T) {
		converted := iterator(elem, idx, list)
		result = append(result, converted)
	})

	return result
}

func MapFilter[T any, V any](list []T, iterator func(elem T, idx int, list []T) (V, bool)) []V {
	result := make([]V, 0, len(list))

	ForEach(list, func(elem T, idx int, l []T) {
		if converted, ok := iterator(elem, idx, list); ok {
			result = append(result, converted)
		}
	})

	return result
}

func Reduce[T any, V any](list []T, accumulator func(acc V, prev T, idx int, l []T) V, initial V) V {
	acc := initial
	ForEach(list, func(elem T, idx int, l []T) {
		acc = accumulator(acc, elem, idx, l)
	})

	return acc
}
