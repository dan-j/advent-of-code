package main

import "strings"

type Path []string

func NewPath(v string) Path {
	if v == "" {
		return Path{}
	}
	return strings.Split(v, ",")
}

func (p Path) CanVisit(v string) bool {
	if strings.ToUpper(v) == v {
		return true
	}

	for _, p := range p {
		if p == v {
			return false
		}
	}

	return true
}

func (p Path) CanVisitPart2(v string) bool {
	if strings.ToUpper(v) == v || v == End {
		return true
	}

	if v == Start {
		return false
	}

	visited := make(map[string]int, len(p))
	var hasDouble bool
	for _, component := range p {
		if strings.ToUpper(component) == component {
			continue
		}

		visited[component]++
		if visited[component] > 1 {
			hasDouble = true
		}
	}

	_, ok := visited[v]
	if !ok {
		// haven't visited this yet
		return true
	}

	return !hasDouble
}

func (p Path) Equal(other Path) bool {
	return p.String() == other.String()
}

func (p Path) String() string {
	return strings.Join(p, ",")
}

type Paths []Path

func (p Paths) Len() int {
	return len(p)
}

func (p Paths) Less(i, j int) bool {
	return p[i].String() < p[j].String()
}

func (p Paths) Swap(i, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}
