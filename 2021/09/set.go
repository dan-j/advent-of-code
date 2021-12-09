package main

import "strings"

type Set map[string]struct{}

func NewSet(inputs ...string) Set {
	set := make(Set, 7)
	for _, input := range inputs {
		for _, r := range input {
			set[string(r)] = struct{}{}
		}
	}

	return set
}

func (s Set) Subtract(set Set) Set {
	res := make(Set, len(s))
	for k, v := range s {
		if _, ok := set[k]; !ok {
			res[k] = v
		}
	}

	return res
}

func (s Set) String() string {
	var sb strings.Builder
	for k := range s {
		sb.WriteString(k)
	}

	return sb.String()
}

func (s Set) Character() string {
	if len(s) != 1 {
		panic("Set.Character(): contains more than 1 element")
	}

	for k := range s {
		return k
	}

	panic("Set is empty")
}

func (s Set) Equals(set Set) bool {
	if len(s) != len(set) {
		return false
	}

	for k := range set {
		if _, ok := s[k]; !ok {
			return false
		}
	}
	return true
}
