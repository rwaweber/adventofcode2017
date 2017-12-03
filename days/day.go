package days

import (
	"container/list"
	"strconv"
)

// The Day interface represents a contract that
// an adventofcode problem must fulfill.
type Day interface {
	Run() string
}

// StringToIntSlice converts a string of digits into a slice of integers.
func StringToIntSlice(intString string) (intSlice []int) {
	for _, char := range intString {
		convertedInt, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, convertedInt)
	}
	return
}

// StringToIntList converts a string of digits into a list of integers.
func StringToIntList(intString string) (intList list.List) {
	for _, char := range intString {
		convertedInt, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		intList.PushBack(convertedInt)
	}
	return
}
