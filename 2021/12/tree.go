package main

import "strings"

const (
	Start = "start"
	End   = "end"
)

type Tree struct {
	Value    string
	Children []*Tree
}

func NewTree(value string) *Tree {
	return &Tree{
		Value: value,
	}
}

func (t *Tree) AddChild(child *Tree) {
	child.Children = append(child.Children, t)
	t.Children = append(t.Children, child)
}

func ParseTree(input string) *Tree {
	m := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		s := strings.Split(line, "-")
		m[s[0]], m[s[1]] = append(m[s[0]], s[1]), append(m[s[1]], s[0])
	}

	root := &Tree{
		Value: Start,
	}

	nodeMap := map[string]*Tree{
		Start: root,
	}

	for _, child := range m[Start] {
		nodeMap[child] = NewTree(child)
		root.AddChild(nodeMap[child])
	}

	delete(m, Start)

	for len(m) != 0 {
		for parent, children := range m {
			parentNode, ok := nodeMap[parent]
			if !ok {
				continue
			}

			for _, c := range children {
				cnode, ok := nodeMap[c]
				if !ok {
					cnode = NewTree(c)
					nodeMap[c] = cnode
				}

				existing := false
				for _, existingChild := range parentNode.Children {
					if existingChild.Value == cnode.Value {
						existing = true
						break
					}
				}

				if !existing {
					parentNode.AddChild(cnode)
				}
			}

			delete(m, parent)
		}
	}

	return root
}
