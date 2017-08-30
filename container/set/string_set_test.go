package set

import (
	"testing"
)

func TestStringSet(t *testing.T) {
	s := NewStringSet("hello", "world", "this", "is", "qiniu", "cloud")
	p := NewStringSet("hello", "qiniu", "is", "a", "company")

	unionRet := s.Union(p)
	t.Logf("%v", unionRet.Elems())

	diffRet := s.Difference(p)
	t.Logf("%v", diffRet.Elems())

	interRet := s.Intersect(p)
	t.Logf("%v", interRet.Elems())
}
