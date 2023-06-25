package main

import (
	"log"
	"math"
)

type cellCondition interface {
	verifiedBy(place cell) (bool, *unit)
}

type targetChecker struct {
	teamNo int
}

func (tc *targetChecker) String() string {
	return "Target acquisition"
}

func (tc *targetChecker) verifiedBy(place cell) (bool, *unit) {
	for _, u := range place.content {
		if u.isLive && u.teamNo != tc.teamNo {
			return true, u
		}
	}
	return false, nil
}

type moveChecker struct {
	field    *world
	who      *unit
	lastPath []coordinates
}

func (tc *moveChecker) String() string {
	return "Destination Finding"
}

func (mc *moveChecker) verifiedBy(place cell) (bool, *unit) {
	for _, u := range place.content {
		if u.isLive && u.teamNo != mc.who.teamNo {
			return mc.pathExists(place.location), u
		}
	}
	return false, nil
}

// can be optimized by precomputing all the paths and later on
// just checking for slot availability on each segment
func (mc *moveChecker) pathExists(goal coordinates) bool {
	log.Println("\tChecking path between", mc.who.position, "and", goal)
	dist := make(map[int64]int)
	prev := make(map[int64]int64)
	Q := make([]int64, 0)
	for v := range mc.field.cells {
		dist[v] = math.MaxInt
		prev[v] = -1
		Q = append(Q, v)
	}
	dist[mc.field.terrain[mc.who.position.row][mc.who.position.col].id] = 0
	for len(Q) != 0 {
		md := math.MaxInt
		mx := 0
		for ix, cd := range Q {
			if dist[cd] < md {
				md = dist[cd]
				mx = ix
			}
		}
		u := Q[mx]
		Q = append(Q[:mx], Q[mx+1:]...)
		ns := mc.field.slotsPerCell - mc.who.kind.slotCount
		for _, v := range Q {
			if isNeighbour(mc.field.cells[v].location, mc.field.cells[u].location) &&
				(mc.field.cells[v].occupiedSlots <= ns) {
				alt := dist[u] + 1
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}
	dest := mc.field.terrain[goal.row][goal.col].id
	temp := make([]int64, 0)
	if prev[dest] != -1 {
		cur := dest
		for cur != -1 {
			temp = append(temp, cur)
			cur = prev[cur]
		}
		pl := len(temp)
		mc.lastPath = make([]coordinates, pl)
		for ix, vid := range temp {
			mc.lastPath[pl-ix-1] = mc.field.cells[vid].location
		}
		log.Println("\tFound Path", mc.lastPath)
	} else {
		log.Println("\tNo Path")
	}
	return prev[dest] != -1
}

type damageInflicter struct {
	field     *world
	teamNo    int // team of the unit inflicting the damage
	hitPoints int
	isSingle  bool
}

func (tc *damageInflicter) String() string {
	return "Damage Distribution"
}

func (d *damageInflicter) verifiedBy(place cell) (bool, *unit) {
	for _, u := range place.content {
		if u.isLive && u.teamNo != d.teamNo {
			u.hitPoints -= max(0, d.hitPoints)
			if u.hitPoints < 1 {
				log.Println("Team", u.teamNo, "'s Unit", u.id, "in", u.position, "was killed")
				u.isLive = false
				delete(d.field.terrain[place.location.row][place.location.col].content, u.id)
				d.field.terrain[place.location.row][place.location.col].occupiedSlots -= u.kind.slotCount
				delete(d.field.liveUnits, u.id)
			}
			if d.isSingle {
				return true, u
			}
		}
	}
	return false, nil
}
