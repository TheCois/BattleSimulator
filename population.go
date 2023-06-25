package main

import (
	"fmt"
	"math/rand"
)

func (a *terrainArea) String() string {
	return a.name
}

func generateBand(w world, ur int, lr int, lc int, rc int) terrainArea {
	if (lr > ur) || (lc > rc) {
		panic("boundaries of the band are inconsistent")
	}
	if (lc < 0) || (lr < 0) || (rc >= w.width) || (ur >= w.length) {
		panic("band is out of the range of the field")
	}
	ids := make([]int64, 0)
	for r := lr; r <= ur; r++ {
		for c := lc; c <= rc; c++ {
			ids = append(ids, w.terrain[r][c].id)
		}
	}
	res := terrainArea{
		nextId(),
		w,
		ids,
		fmt.Sprintf("Band[(%d,%d),(%d, %d)]", lc, lr, rc, ur)}
	return res
}

func generateZone(w world, x int, y int, radius int) terrainArea {
	if (x-radius < 0) || (x+radius >= w.width) || (y-radius < 0) || (y+radius >= w.length) {
		panic("boundaries of the band are inconsistent")
	}
	if (x < 0) || (y < 0) || (x >= w.width) || (y >= w.length) {
		panic("zone centre is out of the field")
	}
	ids := make([]int64, 0)

	v := []coordinates{{x, y}}
	nc := []coordinates{{x, y}}
	for r := 1; r <= radius; r++ {
		_, nc, _ = scanWithArray(&w, v, nc, nil)
		v = append(v, nc...)
	}
	for _, p := range v {
		ids = append(ids, w.terrain[p.row][p.col].id)
	}
	res := terrainArea{
		nextId(),
		w,
		ids,
		fmt.Sprintf("Zone[(%d,%d),%d]", x, y, radius)}
	return res
}

func union(fa terrainArea, sa terrainArea) terrainArea {
	if &fa.parent != &sa.parent {
		panic("Both Areas must have the same parent world")
	}
	values := make(map[int64]bool)
	for _, key := range fa.cellIds {
		values[key] = true
	}
	for _, key := range sa.cellIds {
		values[key] = true
	}
	output := make([]int64, 0, len(values))
	for val := range values {
		output = append(output, val)
	}

	res := terrainArea{
		nextId(),
		fa.parent,
		output,
		fmt.Sprintf("Union(%v, %v)", fa, sa)}
	return res
}

func intersection(fa terrainArea, sa terrainArea) terrainArea {
	if &fa.parent != &sa.parent {
		panic("Both Areas must have the same parent world")
	}
	values := make(map[int64]bool)
	for _, key := range fa.cellIds {
		values[key] = true
	}
	output := make([]int64, 0, len(values))
	for _, key := range sa.cellIds {
		_, ok := values[key]
		if ok {
			output = append(output, key)
		}
	}
	res := terrainArea{
		nextId(),
		fa.parent,
		output,
		fmt.Sprintf("Intersection(%v, %v)", fa, sa)}
	return res
}

func difference(fa terrainArea, sa terrainArea) terrainArea {
	if &fa.parent != &sa.parent {
		panic("Both Areas must have the same parent world")
	}
	values := make(map[int64]bool)
	for _, key := range sa.cellIds {
		values[key] = true
	}
	output := make([]int64, 0, len(values))
	for _, key := range fa.cellIds {
		_, ok := values[key]
		if !ok {
			output = append(output, key)
		}
	}
	res := terrainArea{
		nextId(),
		fa.parent,
		output,
		fmt.Sprintf("Difference(%v, %v)", fa, sa)}
	return res
}

// adds a maximum of n unit
// of unit kind uk
// from team tn
// to the world w
// returning the exact number of units created
func (a *terrainArea) fillWith(n int, uk unitKind, tn int, fk FillKind, rnd *rand.Rand) int {
	avgPerCell := n / len(a.cellIds)
	created := 0
	for _, cid := range a.cellIds {
		cell := a.parent.cells[cid]
		availSlots := a.parent.slotsPerCell - cell.occupiedSlots
		availUnits := availSlots / uk.slotCount
		var numUnits int
		switch fk {
		case Pack:
			numUnits = min(n, availUnits)
			createAndInsert(numUnits, uk, tn, cell, a.parent)
			break
		case Random:
			numUnits = min(n, availUnits, rnd.Intn(avgPerCell))
			createAndInsert(numUnits, uk, tn, cell, a.parent)
			break
		case Uniform:
			numUnits = min(n, availUnits, avgPerCell)
			createAndInsert(numUnits, uk, tn, cell, a.parent)
			break
		}
		n -= numUnits
		created += numUnits
	}
	return created
}

func createAndInsert(num int, uk unitKind, team int, c cell, w world) {
	for i := 0; i < num; i++ {
		u := unit{
			uk,
			nextId(),
			uk.hitPoints,
			uk.attackRate + 1,
			c.location,
			true,
			team}
		c.content[u.id] = &u
		c.occupiedSlots += u.kind.slotCount
		w.liveUnits[u.id] = &u
	}
}
