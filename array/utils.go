package array

import (
	"fmt"
	"reflect"

	"github.com/iVitaliya/colors-go"
)

const (
	_LOG = iota
	INFO
	DEBUG
	WARNING
	ERROR
)

const maxInt int = int(^uint(0) >> 1)

func print(state int, text ...string) {
	var (
		st    string
		open  = colors.BrightBlack("[")
		close = colors.BrightBlack("]")
	)

	go func(_state int) {
		switch _state {
		case INFO:
			st = open + colors.BrightBlue("INFO") + close
			break
		case DEBUG:
			st = open + colors.Green("DEBUG") + close
			break
		case WARNING:
			st = open + colors.Dim(colors.BrightYellow("WARNING")) + close
			break
		case ERROR:
			st = open + colors.Red("ERROR") + close
			break
		}
	}(state)
}

func appendValue[T any](arr []T, value T) []T {
	_arr := arr

	_arr = append(_arr, value)

	return _arr
}

func flattenArray[T comparable](input interface{}, depth int) []T {
	if depth <= 0 {
		return []T{input.(T)}
	}

	v := reflect.ValueOf(input)

	if v.Kind() != reflect.Slice {
		return []T{input.(T)}
	}

	var result []T

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()

		if reflect.TypeOf(elem).Kind() == reflect.Slice {
			result = append(result, flattenArray[T](elem, depth-1)...)
		} else {
			result = append(result, elem.(T))
		}
	}

	return result
}

func searchIndex[T any](arr []T, term T) int {
	var trm string = fmt.Sprint(term)

	for i, v := range arr {
		val := fmt.Sprint(v)

		if trm == val {
			return i
		}
	}

	return -1
}
