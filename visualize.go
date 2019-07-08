package struct_size

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	skip    = "   "
	used    = "[x]"
	notUsed = "[ ]"
)

var ptr = reflect.TypeOf(&struct{}{}).Size()
var nilSize = fmt.Sprintf(`sizeof(nil)=%d
`, ptr*2)

func Visualize(items ...interface{}) (result []byte) {
	for _, item := range items {
		if item == nil {
			result = append(result, nilSize...)
			continue
		}

		t := reflect.TypeOf(item)

		result = append(result, "sizeof("...)
		result = append(result, t.String()...)
		result = append(result, ")="...)
		result = append(result, strconv.Itoa(int(t.Size()))...)
		if t.Kind() == reflect.Struct {
			result = append(result, " with alignment="...)
			result = append(result, strconv.Itoa(t.Align())...)
		}
		result = append(result, '\n')

		if t.Kind() != reflect.Struct {
			continue
		}

		var nameLength, typeLength int
		numFields := t.NumField()
		for i := 0; i < numFields; i++ {
			field := t.Field(i)

			lengthName := len(field.Name)
			if lengthName > nameLength {
				nameLength = lengthName
			}

			lengthType := len(field.Type.String())
			if lengthType > typeLength {
				typeLength = lengthType
			}
		}

		row := uintptr(0)
		newRowLength := nameLength + typeLength + 4

		for i := 0; i < numFields; i++ {
			field := t.Field(i)

			if row < field.Offset {
				result = result[:len(result)-1]
				for ; row < field.Offset; row++ {
					result = append(result, notUsed...)
				}
				result = append(result, '\n')
			}

			if field.Anonymous {
				result = append(result, ' ', ' ', ' ', '~')
			} else {
				result = append(result, ' ', ' ', ' ', ' ')
			}

			lengthName := len(field.Name)
			result = append(result, field.Name...)
			for ; lengthName <= nameLength; lengthName++ {
				result = append(result, ' ')
			}

			result = append(result, field.Type.String()...)
			lengthType := len(field.Type.String())
			for ; lengthType <= typeLength; lengthType++ {
				result = append(result, ' ')
			}

			if int(field.Offset)%t.Align() != 0 {
				i := int(row)
				for ; i%t.Align() != 0; i-- {
					result = append(result, skip...)
				}
			}

			length := field.Offset + field.Type.Size()
			for i := field.Offset; i < length; i++ {
				if int(row)%t.Align() == 0 && i != field.Offset {
					result = append(result, '\n', ' ', ' ', ' ', ' ')
					for i := 2; i < newRowLength; i++ {
						result = append(result, ' ')
					}
				}
				row++
				result = append(result, used...)
			}

			result = append(result, '\n')
		}

		if row < t.Size() {
			result = result[:len(result)-1]
			for ; row < t.Size(); row++ {
				result = append(result, notUsed...)
			}
			result = append(result, '\n')
		}
	}

	return
}
