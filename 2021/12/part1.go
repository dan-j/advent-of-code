package main

import (
	"fmt"
	"sort"
)

func Part1(input string) int {
	type State struct {
		Path string
		Node *Tree
	}

	tree := ParseTree(input)

	var paths []Path
	path := Path{}
	queue := map[State]struct{}{
		State{
			Path: Start,
			Node: tree,
		}: {},
	}

	for len(queue) != 0 {
		for state := range queue {
			delete(queue, state)

			path = NewPath(state.Path)

			for _, child := range state.Node.Children {
				if !path.CanVisit(child.Value) {
					continue
				}

				childPath := append(path, child.Value)
				if child.Value == End {
					paths = append(paths, childPath)
					continue
				}

				queue[State{
					Path: childPath.String(),
					Node: child,
				}] = struct{}{}
			}
		}
	}

	sort.Sort(Paths(paths))
	for _, p := range paths {
		fmt.Println(p.String())
	}
	return len(paths)
}
