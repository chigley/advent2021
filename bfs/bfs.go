package bfs

import (
	"container/list"

	"github.com/chigley/advent2021"
)

type Node interface {
	IsGoal() bool
	Neighbours() ([]Node, error)
	Key() interface{}
}

func Search(start Node) ([]Node, error) {
	paths, err := search(start, false)
	if err != nil {
		return nil, err
	}
	return paths[0], nil
}

func SearchAll(start Node) ([][]Node, error) {
	return search(start, true)
}

func search(start Node, findAll bool) ([][]Node, error) {
	q := list.New()
	q.PushBack(start)

	parent := map[Node]Node{
		start: nil,
	}

	discovered := map[interface{}]struct{}{
		start.Key(): {},
	}

	var paths [][]Node

	for e := q.Front(); e != nil; e = e.Next() {
		node := e.Value.(Node)

		if node.IsGoal() {
			var path []Node
			for n := node; n != nil; n = parent[n] {
				path = append([]Node{n}, path...)
			}
			paths = append(paths, path[1:])

			if !findAll {
				return paths, nil
			}
			continue
		}

		neighbours, err := node.Neighbours()
		if err != nil {
			return nil, err
		}
		for _, n := range neighbours {
			if _, ok := discovered[n.Key()]; ok {
				continue
			}
			discovered[n.Key()] = struct{}{}
			parent[n] = node
			q.PushBack(n)
		}
	}

	if len(paths) == 0 {
		return nil, advent2021.ErrNoSolution
	}
	return paths, nil
}
