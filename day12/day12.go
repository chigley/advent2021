package day12

import (
	"fmt"
	"strings"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/bfs"
)

const (
	CaveStart = "start"
	CaveEnd   = "end"
)

func Part1(cs CaveSystem) (int, error) {
	return numPaths(cs, false)
}

func Part2(cs CaveSystem) (int, error) {
	return numPaths(cs, true)
}

func numPaths(cs CaveSystem, allowedSecondVisit bool) (int, error) {
	start := SearchNode{
		CaveSystem:         cs,
		Path:               []string{CaveStart},
		VisitedSmall:       nil,
		AllowedSecondVisit: allowedSecondVisit,
	}

	paths, err := bfs.SearchAll(&start)
	if err != nil {
		return 0, fmt.Errorf("day12: bfs: %w", err)
	}

	return len(paths), nil
}

type SearchNode struct {
	CaveSystem         CaveSystem
	Path               []string
	VisitedSmall       Caves
	AllowedSecondVisit bool
}

func (n *SearchNode) IsGoal() bool {
	return len(n.Path) > 0 && n.Path[len(n.Path)-1] == CaveEnd
}

func (n *SearchNode) Neighbours() ([]bfs.Node, error) {
	if len(n.Path) == 0 {
		return nil, nil
	}

	currentCave := n.Path[len(n.Path)-1]

	var neighbours []bfs.Node
	for neighbourCave := range n.CaveSystem[currentCave] {
		// Don't re-visit the start
		if neighbourCave == CaveStart {
			continue
		}

		// Don't re-visit a small cave, unless we haven't done so yet
		nextAllowedSecondVisit := n.AllowedSecondVisit
		if advent2021.IsLower(neighbourCave) {
			if _, ok := n.VisitedSmall[neighbourCave]; ok {
				if n.AllowedSecondVisit {
					nextAllowedSecondVisit = false
				} else {
					continue
				}
			}
		}

		// Clone n.Path and append neighbourCave
		nextPath := append([]string(nil), n.Path...)
		nextPath = append(nextPath, neighbourCave)

		// Clone n.VisitedSmall and add neighbourCave
		nextVisitedSmall := make(Caves, len(n.VisitedSmall)+1)
		for k, v := range n.VisitedSmall {
			nextVisitedSmall[k] = v
		}
		nextVisitedSmall[neighbourCave] = struct{}{}

		neighbours = append(neighbours, &SearchNode{
			CaveSystem:         n.CaveSystem,
			Path:               nextPath,
			VisitedSmall:       nextVisitedSmall,
			AllowedSecondVisit: nextAllowedSecondVisit,
		})
	}
	return neighbours, nil
}

func (n *SearchNode) Key() interface{} {
	return n.String()
}

func (n *SearchNode) String() string {
	return strings.Join(n.Path, ",")
}
