package main

type unitKind struct {
	name           string
	hitPoints      int
	perception     int
	attackRange    int // how many cells away can hit
	attackAccuracy int // probability of hitting whatever is aimed at (percentage)
	attackRate     int // number of rounds in between attacks
	attackDamage   int // hitPoints taken from target
	attackArea     int // cells affected by attack (0 for single unit effect)
	moveSpeed      int // max number of cells moved in any round
	slotCount      int // number of slots taken
}

type coordinates struct {
	col int
	row int
}

type move struct {
	obj *unit
	to  coordinates
}

type unit struct {
	kind                  unitKind
	id                    int64
	hitPoints             int
	roundsSinceLastAttack int
	position              coordinates
	isLive                bool
	teamNo                int
}

type hitZone struct {
	teamNo int // team of the unit making the hit
	target *unit
	radius int
	damage int
}

type cell struct {
	id            int64
	content       map[int64]*unit
	location      coordinates
	occupiedSlots int
	neighbours    []int64 //all neighbouring cell ids in clockwise from the top order
}

type world struct {
	length       int
	width        int
	terrain      [][]cell
	cells        map[int64]cell
	slotsPerCell int
	liveUnits    map[int64]*unit
}

type terrainArea struct {
	id      int64
	parent  world
	cellIds []int64
	name    string
}

type FillKind int

const (
	Pack FillKind = iota
	Uniform
	Random
)
