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
	q := list.New()
	q.PushBack(start)

	parent := map[Node]Node{
		start: nil,
	}

	discovered := map[interface{}]struct{}{
		start.Key(): {},
	}

	for e := q.Front(); e != nil; e = e.Next() {
		node := e.Value.(Node)

		if node.IsGoal() {
			var path []Node
			for n := node; n != nil; n = parent[n] {
				path = append([]Node{n}, path...)
			}
			return path[1:], nil
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

	return nil, advent2021.ErrNoSolution
}
