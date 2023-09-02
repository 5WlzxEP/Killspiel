package User

import "slices"

const bufferSize = 10

var Buffer = buffer{}

type buffer struct {
	buf   [bufferSize]key
	index int8
}

type key struct {
	Key string `json:"key,omitempty"`
	id  int
}

func (b *buffer) add(val key) {
	b.buf[b.index] = val
	b.index++
	b.index %= bufferSize
}

func (b *buffer) contains(value key) bool {
	return slices.Index(b.buf[:], value) != -1
}
