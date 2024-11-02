package array

type Array[T comparable] struct {
	array []T
}

func New[T comparable]() *Array[T] {
	return &Array[T]{
		array: []T{},
	}
}

func NewWithEntries[T comparable](entries ...T) *Array[T] {
	var arr []T

	for _, v := range entries {
		arr = appendValue[T](arr, v)
	}

	return &Array[T]{
		array: arr,
	}
}

// From converts a string to an array of its characters.
// The output is a slice of strings, where each element is a single character.
// Example: From("hello") returns ["h", "e", "l", "l", "o"].
func From(str string) []string {
	var strArr []string

	for i := 0; i < len(str); i++ {
		strArr = append(strArr, string(str[i]))
	}

	return strArr
}

// FromIter applies the given function to each element of the given array,
// returning the same array. It is similar to the Array.Map function, but
// does not return a new array.
func FromIter[T comparable](arr []T, fn func(T)) []T {
	for _, v := range arr {
		fn(v)
	}

	return arr
}

// At returns the value at the given index.
func (array *Array[T]) At(index int) T {
	for i, v := range array.array {
		if i == index {
			return v
		}
	}

	return *new(T)
}

// Append adds the given values to the end of the array.
func (array *Array[T]) Append(value ...T) {
	var arr []T = array.array

	for _, v := range value {
		arr = append(arr, v)
	}

	array.array = arr
}

// Concat appends the elements of the given arrays to the end of the array.
// It does not create a new array, but changes the original array.
// The return value is the length of the new array.
func (array *Array[T]) Concat(elements ...[]T) {
	var arr []T = array.array

	for _, v := range elements {
		arr = append(arr, v...)
	}

	array.array = arr
}

// IndexOf returns the index of the first occurrence of the specified element in the array,
// or -1 if it is not present.
func (array *Array[T]) IndexOf(search_term T) int {
	for i, v := range array.array {
		if v == search_term {
			return i
		}
	}

	return -1
}

// CopyWithin copies the elements of the given array from the start index up to but not including the end index into a new array.
// The new array is then returned.
// The elements are copied in the same order as they appear in the original array.
// The start and end indices are both inclusive.
func (array *Array[T]) CopyWithin(start, end int) []T {
	var arr []T

	for i := start; i < end; i++ {
		arr = append(arr, array.array[i-end])
	}

	return arr
}

// Entries returns the elements of the array as a slice.
// The returned slice is a view over the same elements as the array.
// Modifying the returned slice will modify the array.
func (array *Array[T]) Entries() []T {
	return array.array
}

// Every tests whether all elements in the array pass the test implemented by the provided function.
// It returns true if all elements pass the test, and false otherwise.
// It calls the provided function once for each element present in the array until it finds one where falsy is returned.
// If such an element is found, the Every method immediately returns false.
// Otherwise, if the callback function returns a truthy value for all elements, Every returns true.
func (array *Array[T]) Every(fn func(T) bool) bool {
	for _, v := range array.array {
		return fn(v)
	}

	return false
}

func (array *Array[T]) Fill(element T, start, end int) []T {
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fill
}

func test() {
	t := New[string]()

}
