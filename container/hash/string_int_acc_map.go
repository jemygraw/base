package hash

import (
	"sort"
)

type StringIntAccMap map[string]int

func NewStringIntAccMap() StringIntAccMap {
	m := make(map[string]int)
	return m
}

func (m StringIntAccMap) Append(key string, value int) {
	if _, ok := m[key]; ok {
		m[key] += value
	} else {
		m[key] = value
	}
}

func (m StringIntAccMap) Get(key string) (value int) {
	if v, ok := m[key]; ok {
		value = v
	} else {
		value = 0
	}
	return
}

func (m StringIntAccMap) Keys() []string {
	keys := make([]string, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m StringIntAccMap) SortedKeys() []string {
	keys := m.Keys()
	sort.Strings(keys)
	return keys
}
