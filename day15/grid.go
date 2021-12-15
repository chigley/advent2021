package day15

import (
	"container/heap"
	"math"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/priority_queue"
)

type Grid map[advent2021.XY]int

func NewGrid(in []string) Grid {
	g := make(Grid, len(in)*len(in[0]))
	for y, l := range in {
		for x, r := range l {
			g[advent2021.XY{X: x, Y: y}] = int(r - '0')
		}
	}
	return g
}

func (g Grid) ShortestPath(from, to advent2021.XY) int {
	dist := make(map[advent2021.XY]int)

	var pq priority_queue.PriorityQueue
	heap.Init(&pq)
	heap.Push(&pq, priority_queue.NewItem(from, 0))

	for pq.Len() > 0 {
		curPos := heap.Pop(&pq).(*priority_queue.Item).Value().(advent2021.XY)

		for _, nPos := range curPos.Neighbours(false) {
			cost, ok := g[nPos]
			if !ok {
				// Neighbour is out of bounds
				continue
			}

			curBestToNeighbour, ok := dist[nPos]

			altViaCur := dist[curPos] + cost

			if !ok || altViaCur < curBestToNeighbour {
				dist[nPos] = altViaCur

				// We could actually already have nPow in the queue at this
				// point, but with a higher priority. It's not convenient to
				// change the priority of that, so we just insert a duplicate.
				//
				// ¯\_(ツ)_/¯
				heap.Push(&pq, priority_queue.NewItem(nPos, altViaCur))
			}
		}
	}

	return dist[to]
}

func (g Grid) BottomRight() advent2021.XY {
	maxX, maxY := math.MinInt, math.MinInt
	for pos := range g {
		maxX = advent2021.Max(maxX, pos.X)
		maxY = advent2021.Max(maxY, pos.Y)
	}
	return advent2021.XY{X: maxX, Y: maxY}
}
