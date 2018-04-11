package hash

import (
	"fmt"
	"testing"
)

func TestStringIntAccMap(t *testing.T) {
	m := NewStringIntAccMap()
	m.Append("apple", 10)
	m.Append("orange", 20)
	m.Append("apple", 5)
	m.Append("banana", 1)
	m.Append("apple", 1)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
