package main

import (
	"math"
)

var id int64

func max(values ...int) int {
	res := math.MinInt
	for _, v := range values {
		if v > res {
			res = v
		}
	}
	return res
}

func min(values ...int) int {
	res := math.MaxInt
	for _, v := range values {
		if v < res {
			res = v
		}
	}
	return res
}

func neighbour(c coordinates, i int) coordinates {
	dx := []int{0, 1, 1, 0, -1, -1}
	dyOdd := []int{1, 1, 0, -1, 0, 1}
	dyEven := []int{1, 0, -1, -1, -1, 0}

	if c.col%2 == 1 {
		return coordinates{c.col + dx[i], c.row + dyOdd[i]}
	} else {
		return coordinates{c.col + dx[i], c.row + dyEven[i]}
	}
}

func isNeighbour(p1, p2 coordinates) bool {
	if (math.Abs(float64(p1.col-p2.col)) > 1) || (math.Abs(float64(p1.row-p2.row)) > 1) {
		return false
	}
	return true
}

func isValidPosition(w *world, p coordinates) bool {
	return p.col >= 0 && p.row >= 0 && p.col < w.width && p.row < w.length
}

func notInLinear(v []coordinates, n coordinates) bool {
	for _, cp := range v {
		if (cp.col == n.col) && (cp.row == n.row) {
			return false
		}
	}
	return true
}

func notInConstant(v map[int]bool, n coordinates, dim int) bool {
	ix := n.row*dim + n.col
	_, ok := v[ix]
	return !ok
}

func nextId() int64 {
	id++
	return id
}

func duplicatesIn(a []coordinates) bool {
	const W = 10000
	m := make(map[int]bool)
	for _, c := range a {
		_, in := m[c.row*W+c.col]
		if !in {
			m[c.row*W+c.col] = true
		} else {
			return true
		}
	}
	return false
}
