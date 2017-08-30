package set

// StringSet is not routine-safe
type StringSet struct {
	setElems map[string]struct{}
}

// NewStringSet returns a new string set
func NewStringSet(elems ...string) *StringSet {
	setElems := make(map[string]struct{})
	for _, e := range elems {
		setElems[e] = struct{}{}
	}
	return &StringSet{
		setElems: setElems,
	}
}

// Elems returns the elems of the string set
func (s *StringSet) Elems() []string {
	elems := make([]string, 0, len(s.setElems))
	for k := range s.setElems {
		elems = append(elems, k)
	}
	return elems
}

// Append appends the elems into the string set and return itself
func (s *StringSet) Append(elems ...string) *StringSet {
	for _, e := range elems {
		s.setElems[e] = struct{}{}
	}
	return s
}

// Union returns a new string set which holds the union elems of both set
func (s *StringSet) Union(p *StringSet) *StringSet {
	unionElems := make(map[string]struct{})
	for k, v := range s.setElems {
		unionElems[k] = v
	}

	for k, v := range p.setElems {
		unionElems[k] = v
	}

	return &StringSet{
		setElems: unionElems,
	}
}

//Intersect returns a new string set which holds the elems intersect of both set
func (s *StringSet) Intersect(p *StringSet) *StringSet {
	intersectElems := make(map[string]struct{})
	for k, v := range s.setElems {
		if _, ok := p.setElems[k]; ok {
			intersectElems[k] = v
		}
	}

	return &StringSet{
		setElems: intersectElems,
	}
}

//Difference returns a new string set which holds the elems in s but not in p
func (s *StringSet) Difference(p *StringSet) *StringSet {
	diffElems := make(map[string]struct{})
	for k, v := range s.setElems {
		if _, ok := p.setElems[k]; !ok {
			diffElems[k] = v
		}
	}
	return &StringSet{
		setElems: diffElems,
	}
}
