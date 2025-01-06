package array

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/iVitaliya/logger-go"
)

type Array[T comparable] struct {
	array []T
}

// New returns a new empty array
//
// Example:
//
//	arr := array.New[int]()
//	// arr is now an empty array of type []int
func New[T comparable]() *Array[T] {
	return &Array[T]{
		array: []T{},
	}
}

// NewWithEntries creates a new array and appends the given entries to it.
// It is different from the New function in that it takes an array of elements
// and appends them to the new array, rather than returning an empty array.
// The returned array is a new array with the given entries, and not the same
// as the original array.
//
// Example:
//
//	arr := array.NewWithEntries[int]([]int{1, 2, 3})
//	// arr is now an array of type []int with elements 1, 2, 3
func NewWithEntries[T comparable](entries []T) *Array[T] {
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
func FromIter[T comparable](arr []T, fn func(value T)) []T {
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
func (array *Array[T]) Every(fn func(value T) bool) bool {
	for _, v := range array.array {
		return fn(v)
	}

	return false
}

// Fill fills all the elements of the array from a start index to an end index with a static value.
// The start index is inclusive, and the end index is exclusive.
// If the start index is negative, it is treated as an offset from the end of the array.
// If the end index is negative, it is treated as an offset from the end of the array.
// If the start index is greater than the end index, the elements are filled in reverse order.
// If the start index is out of range, the function will return the array unchanged.
// If the end index is out of range, the function will return the array unchanged.
func (array *Array[T]) Fill(element T, start int, end ...int) []T {
	var (
		endIndex int
		lastInd  = len(array.array) - 1
	)
	if end[0] > lastInd {
		endIndex = lastInd
	} else {
		endIndex = end[0]
	}

	if start < 0 || start >= len(array.array) {
		logger.Error("Start index out of range")
		return array.array
	}

	if endIndex < 0 || endIndex >= len(array.array) {
		logger.Error("End index out of range")
		return array.array
	}

	for i := start; i < endIndex; i++ {
		array.array[i] = element
	}

	return array.array
}

// Filter creates a new array with all elements that pass the test implemented by the provided function.
// It takes a callback function as an argument, which is called once for each element present in the array.
// The callback function takes one argument, the element value, and returns true if the element passes the test, false otherwise.
// The returned array is a filtered version of the original array.
// The elements are copied in the same order as they appear in the original array.
func (array *Array[T]) Filter(fn func(value T) bool) []T {
	var result []T
	for _, item := range array.array {
		if fn(item) {
			result = append(result, item)
		}
	}

	array.array = result
	return result
}

// Find returns the first element in the array that satisfies the provided testing function.
// The function iterates through the array and returns the element and true if found.
// If no element is found, it returns a zero value of type T and false.
// The testing function takes one argument, the element value, and returns true if the element passes the test, false otherwise.
func (array *Array[T]) Find(fn func(value T) bool) (T, bool) {
	for _, element := range array.array {
		if fn(element) {
			return element, true
		}
	}

	return *new(T), false
}

// FindIndex returns the index of the first element in the array that satisfies the provided testing function.
// If no element is found, it returns -1.
// The testing function takes one argument, the element value, and returns true if the element passes the test, false otherwise.
// The index returned is the index of the element in the original array.
func (array *Array[T]) FindIndex(fn func(value T) bool) (int, bool) {
	for i, element := range array.array {
		if fn(element) {
			return i, true
		}
	}

	return -1, false
}

// FindLast returns the last element in the array that satisfies the provided testing function.
// The function iterates through the array in reverse order, and returns the element and true if found.
// If no element is found, it returns a zero value of type T and false.
// The testing function takes one argument, the element value, and returns true if the element passes the test, false otherwise.
func (array *Array[T]) FindLast(fn func(value T) bool) (T, bool) {
	for i := len(array.array); i > 0; i-- {
		if fn(array.array[i-1]) {
			return array.array[i-1], true
		}
	}

	return *new(T), false
}

// FindLastIndex returns the index of the last element in the array that satisfies the provided testing function.
// If no element is found, it returns -1.
// The testing function takes one argument, the element value, and returns true if the element passes the test, false otherwise.
// The function iterates through the array in reverse order, and the index returned is the index of the element in the original array.
func (array *Array[T]) FindLastIndex(fn func(value T) bool) (int, bool) {
	for i := len(array.array); i > 0; i-- {
		if fn(array.array[i-1]) {
			return i - 1, true
		}
	}

	return -1, false
}

// Flat returns a new array with the elements of the given array flattened to a single depth.
// The depth parameter specifies the maximum recursion depth.
// If the depth is 0, the elements of the array are returned unchanged.
// If the depth is 1, the elements of the array are flattened one level deep.
// If the depth is 2 or more, the elements of the array are flattened as many levels deep as the depth.
// If the depth is negative, the elements of the array are flattened as many levels deep as possible.
func (array *Array[T]) Flat(depth int) []T {
	arr := array.array

	return flattenArray[T](arr, depth)
}

// FlatMap applies the given function to each element of the array,
// returning a new array with the results of each function call
// flattened into a single array.
// It is similar to the Map function, but the callback function can
// return more than one value, and the returned values are flattened
// into a single array.
func (array *Array[T]) FlatMap(fn func(value T) []T) []T {
	var result []T

	for _, item := range array.array {
		result = append(result, fn(item)...)
	}

	return result
}

// ForEach calls the provided function once for each element present in the array in ascending order.
func (array *Array[T]) ForEach(fn func(value T)) {
	for _, item := range array.array {
		fn(item)
	}
}

// Includes determines whether the array includes a certain element, returning true or false as appropriate.
func (array *Array[T]) Includes(search_term T) bool {
	for _, v := range array.array {
		if v == search_term {
			return true
		}
	}

	return false
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

// Join joins all elements of an array into a string.
//
// The separator parameter is a string that is inserted between each pair of adjacent elements.
// If the separator is an empty string, all elements are joined without any characters in between.
//
// The string conversions of all elements are made using the fmt package's Sprint function.
//
// Panics if the number of elements is large enough that the length of the resulting string would overflow a int.
func (array *Array[T]) Join(separator string) string {
	switch len(array.array) {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(array.array[0])
	}

	var n int
	if len(separator) > 0 {
		if len(separator) >= maxInt/(len(array.array)-1) {
			panic("Join output length overflow")
		}

		n += len(separator) * (len(array.array) - 1)
	}

	if reflect.TypeOf(array.array[0]).Kind() == reflect.String {
		for _, item := range array.array {
			elem := fmt.Sprint(item)
			if len(elem) > maxInt-n {
				panic("Join output length overflow")
			}

			n += len(elem)
		}
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprint(array.array[0]))

	for _, s := range array.array[1:] {
		b.WriteString(separator)
		b.WriteString(fmt.Sprint(s))
	}

	return b.String()
}

// Keys returns a slice of the keys of the array.
// The keys are the indices of the elements in the array.
// The order of the keys is the same as the order of the elements in the array.
// The returned slice is a view over the keys of the array,
// and modifying the returned slice will modify the array.
func (array *Array[T]) Keys() []int {
	var result []int

	for i, _ := range array.array {
		result = append(result, i)
	}

	return result
}

// LastIndexOf returns the index of the last occurrence of the specified element in the array,
// or -1 if it is not present.
// The elements are searched in reverse order, and the index returned is the index of the element in the original array.
func (array *Array[T]) LastIndexOf(search_term T) int {
	for i := len(array.array) - 1; i >= 0; i-- {
		if array.array[i] == search_term {
			return i
		}
	}

	return -1
}

// Map applies the provided function to each element of the array and returns a new array containing the results.
// The provided function takes an element of the array and returns a slice of elements, which are then flattened into the result array.
// This function is similar to FlatMap, allowing the callback to return multiple values that are flattened into the final array.
func (array *Array[T]) Map(fn func(value T) []T) []T {
	var result []T

	for _, item := range array.array {
		result = append(result, fn(item)...)
	}

	return result
}

// Pop removes the last element from the array and returns it.
// If the array is empty, it returns a zero value of type T.
func (array *Array[T]) Pop() T {
	if len(array.array) == 0 {
		return *new(T)
	}

	result := array.array[len(array.array)-1]
	array.array = array.array[:len(array.array)-1]

	return result
}

// Reduce applies a function against an accumulator and each element in the array
// (from left to right) so as to reduce it to a single value.
//
// The callback function takes four arguments:
//   - accumulator: The returned value of the previous callback, or the initial value.
//   - value: The current element being processed in the array.
//   - index: The index of the current element being processed in the array.
//   - array: The array the element belongs to.
//
// The initial value is the first element of the array. If the array is empty, it returns a zero value of type T.
func (array *Array[T]) Reduce(fn func(accumulator T, value T) T) T {
	var result T

	for _, item := range array.array {
		result = fn(result, item)
	}

	return result
}

// ReduceRight applies a function against an accumulator and each element in the array
// (from right to left) so as to reduce it to a single value.
//
// The callback function takes two arguments:
//   - accumulator: The returned value of the previous callback, or the initial value.
//   - value: The current element being processed in the array.
//
// The initial value is the last element of the array. If the array is empty, it returns a zero value of type T.
func (array *Array[T]) ReduceRight(fn func(accumulator T, value T) T) T {
	var result T

	for i := len(array.array) - 1; i >= 0; i-- {
		result = fn(result, array.array[i])
	}

	return result
}

// Reverse reverses the elements of the array in place.
// The first element of the array becomes the last, and the last element becomes the first.
// All other elements are shifted accordingly.
func (array *Array[T]) Reverse() {
	for i, j := 0, len(array.array)-1; i < j; i, j = i+1, j-1 {
		array.array[i], array.array[j] = array.array[j], array.array[i]
	}
}

// Push adds the given value to the end of the array.
// The return value is the new length of the array.
func (array *Array[T]) Push(value T) {
	array.array = append(array.array, value)
}

// Shift removes the first element from the array and returns it.
// If the array is empty, it returns a zero value of type T.
// The remaining elements are shifted, and the length of the array is reduced by one.
func (array *Array[T]) Shift() T {
	if len(array.array) == 0 {
		return *new(T)
	}

	result := array.array[0]
	array.array = array.array[1:]

	return result
}

// Slice returns a shallow copy of a portion of the array from the start index to the end index (exclusive).
// The start index is inclusive, while the end index is exclusive. It panics if the indices are out of range.
func (array *Array[T]) Slice(start int, end int) []T {
	return array.array[start:end]
}

// Some tests whether at least one element in the array passes the test implemented by the provided function.
// The function returns true if at least one element passes the test, and false otherwise.
// It calls the provided function once for each element present in the array until it finds one where the function returns true.
// If such an element is found, the Some method immediately returns true.
// Otherwise, if the callback function returns false for all elements, Some returns false.
func (array *Array[T]) Some(fn func(value T) bool) bool {
	for _, item := range array.array {
		if fn(item) {
			return true
		}
	}

	return false
}

// Sort sorts the elements of the array in place according to the provided comparison function.
// The comparison function determines the order of the elements. It takes two arguments, i and j,
// which are elements of the array, and returns true if the element i should come before the element j.
// This method modifies the original array and does not return a new array.
func (array *Array[T]) Sort(fn func(i, j T) bool) {
	sort.Slice(array.array, func(i, j int) bool {
		return fn(array.array[i], array.array[j])
	})
}

// Splice changes the content of the array by removing or replacing existing elements and/or adding new elements in place.
// The start parameter is the index at which to start changing the array.
// The end parameter is the index at which to stop changing the array.
// The value parameter is the element to be added to the array.
// The return value is the modified array.
func (array *Array[T]) Splice(start, end int, value T) []T {
	return append(array.array[:start], append([]T{value}, array.array[end:]...)...)
}

// ToReverse returns a new array with the elements of the original array in reverse order.
// The original array remains unchanged, and the returned array contains the same elements
// but in reversed sequence.
func (array *Array[T]) ToReverse() []T {
	var result []T

	for i := len(array.array) - 1; i >= 0; i-- {
		result = append(result, array.array[i])
	}

	return result
}

// ToSorted returns a new array with the elements of the original array sorted according to the provided comparison function.
// The original array remains unchanged, and the returned array contains the same elements
// but in the sorted sequence.
// The comparison function determines the order of the elements. It takes two arguments, a and b,
// which are elements of the array, and returns true if the element a should come before the element b.
func (array *Array[T]) ToSorted(fn func(a, b T) bool) []T {
	copySlice := make([]T, len(array.array))
	copy(copySlice, array.array)

	sort.Slice(copySlice, func(i, j int) bool {
		return fn(copySlice[i], copySlice[j])
	})

	return copySlice
}

// ToSpliced returns a new array with elements added, removed, or replaced based on the specified parameters.
// The 'start' parameter determines the index at which to begin changing the array. If negative, it is treated
// as an offset from the end of the array. The 'deleteCount' parameter specifies the number of elements to remove
// from the array starting at the 'start' index. If 'deleteCount' is negative, it is set to zero. If 'start' or
// 'start + deleteCount' exceeds the array bounds, they are clamped appropriately. The 'items' parameter allows
// for new elements to be added to the array at the 'start' index. The original array remains unchanged.
func (array *Array[T]) ToSpliced(start, deleteCount int, items ...T) []T {
	// Ensure start is within bounds.
	if start < 0 {
		start = 0
		start += len(array.array)
	}
	if start > len(array.array) {
		start = len(array.array)
	}

	// Ensure deleteCount is valid.
	if deleteCount < 0 {
		deleteCount = 0
	}
	if start+deleteCount > len(array.array) {
		deleteCount = len(array.array) - start
	}

	result := make([]T, 0, len(array.array)-deleteCount+len(items))

	result = append(result, array.array[:start]...)
	result = append(result, items...)
	result = append(result, array.array[start+deleteCount:]...)

	return result
}

// ToString returns a string representation of the array, using the fmt package's Sprint function.
// It is the same as calling fmt.Sprintf("%v", array.array).
func (array *Array[T]) ToString() string {
	return fmt.Sprintf("%v", array.array)
}

// Unshift adds one or more elements to the beginning of the array and returns the new array, along with its length.
// The elements are inserted in the same order as they appear in the parameters.
// The returned array is a new array with the same elements as the original array, but with the new elements added at the beginning.
// The original array remains unchanged.
func (array *Array[T]) Unshift(elements ...T) ([]T, int) {
	result := append(elements, array.array...)

	return result, len(result)
}

// Values returns the elements of the array as a slice.
// The returned slice is a view over the same elements as the array.
// Modifying the returned slice will modify the array.
func (array *Array[T]) Values() []T {
	return array.array
}

// With returns a new array with the value at the given index replaced with the given value.
// If the index is out of range, it returns an error.
// The returned array is a new array with the same elements as the original array, but with the element at the given index replaced.
// The original array remains unchanged.
func (array *Array[T]) With(index int, value T) ([]T, error) {
	if index < 0 || index >= len(array.array) {
		return nil, fmt.Errorf("index out of range")
	}

	newSlice := make([]T, len(array.array))
	copy(newSlice, array.array)

	newSlice[index] = value

	return newSlice, nil
}
