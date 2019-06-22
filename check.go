package struct_size

import (
	"reflect"
)

func CheckSize(item interface{}) (uint, uint, bool) {
	if item == nil {
		currentSize := uint(ptr) * 2

		return currentSize, currentSize, true
	}

	t := reflect.TypeOf(item)

	currentSize := uint(t.Size())

	if t.Kind() != reflect.Struct {
		return currentSize, currentSize, true
	}

	var bestSize uintptr

	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		bestSize += field.Type.Size()
	}

	if bestSize > ptr && bestSize%ptr != 0 {
		size := bestSize / ptr
		size++
		bestSize = size * ptr
	}

	return currentSize, uint(bestSize), currentSize == uint(bestSize)
}
