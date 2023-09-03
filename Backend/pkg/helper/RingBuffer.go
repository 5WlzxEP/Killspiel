package helper

import "slices"

const bufferSize = 10

type Buffer[T comparable] struct {
	buf   [bufferSize]T
	index int8
}

func (b *Buffer[T]) Add(val T) {
	b.buf[b.index] = val
	b.index++
	b.index %= bufferSize
}

func (b *Buffer[T]) Contains(value T) bool {
	return slices.Index(b.buf[:], value) != -1
}
