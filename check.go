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

	var bestSize int

	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		bestSize += int(field.Type.Size())
	}

	if bestSize%t.Align() != 0 {
		size := bestSize / t.Align()
		size++
		bestSize = size * t.Align()
	}

	return currentSize, uint(bestSize), currentSize == uint(bestSize)
}
