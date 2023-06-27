package main

import (
	"log"
	"math/rand"
	"time"
)

type BattleGenerator interface {
	create() world
}

// generates an empty World of w times l cells
// s is the number of slots held in each cell
// should precompute the neighbours of each cell
func generateWorld(w int, l int, s int) world {
	wo := world{
		length:       l,
		width:        w,
		terrain:      make([][]cell, l),
		cells:        make(map[int64]cell),
		liveUnits:    make(map[int64]*unit),
		slotsPerCell: s}
	log.Println("Generating a field of", l, "rows and", w, "columns")
	for i := 0; i < l; i++ {
		wo.terrain[i] = make([]cell, w)
		for j := 0; j < w; j++ {
			ni := nextId()
			c := cell{
				ni,
				make(map[int64]*unit, 0),
				coordinates{j, i},
				0,
				make([]int64, 0),
			}
			wo.terrain[i][j] = c
			wo.cells[ni] = c
		}
	}
	return wo
}

// DefaultMediumGenerator
// Flat Uniform Map with 150 x 1000 Cells
// 1,000,000 Zombies
// 20 Tanks
// 2000 soldiers
// 10 Howitzers
// 200 Special Operators
//
//goland:noinspection GoUnusedExportedType
type DefaultMediumGenerator struct {
	randomizerSeed int64
}

func (g *DefaultMediumGenerator) create() world {
	if g.randomizerSeed == 0 {
		g.randomizerSeed = time.Now().UnixNano()
	}
	s1 := rand.NewSource(g.randomizerSeed)
	r1 := rand.New(s1)

	battleField := generateWorld(1000, 150, 64)

	ttZ := generateBand(battleField, 50, 1, 1, 900)
	ttZ.fillWith(1000000, zombieKind, 2, Pack, r1)

	ttT := generateBand(battleField, 100, 99, 100, 200)
	ttT.fillWith(20, tankKind, 1, Uniform, r1)

	ttI := generateBand(battleField, 98, 98, 100, 200)
	ttI.fillWith(2000, infantryKind, 1, Uniform, r1)

	ttH := generateBand(battleField, 90, 90, 100, 200)
	ttH.fillWith(10, howitzerKind, 1, Uniform, r1)

	ttO := generateBand(battleField, 89, 89, 100, 200)
	ttO.fillWith(200, operatorKind, 1, Random, r1)

	return battleField
}

// DefaultTinyGenerator
// Flat Uniform Map with 10 x 10 Cells
// 4 Zombies
// 1 Archer
// 5 peasants
type DefaultTinyGenerator struct {
	randomizerSeed int64
}

func (g DefaultTinyGenerator) create() world {
	if g.randomizerSeed == 0 {
		g.randomizerSeed = time.Now().UnixNano()
	}
	s1 := rand.NewSource(g.randomizerSeed)
	r1 := rand.New(s1)

	battleField := generateWorld(10, 10, 64)

	ttZ := generateBand(battleField, 1, 1, 4, 7)
	ttZ.fillWith(20, zombieKind, 2, Uniform, r1)

	ttT := generateBand(battleField, 8, 8, 5, 5)
	ttT.fillWith(1, archerKind, 1, Uniform, r1)

	ttI := generateBand(battleField, 6, 6, 5, 6)
	ttI.fillWith(5, peasantKind, 1, Uniform, r1)

	return battleField
}
