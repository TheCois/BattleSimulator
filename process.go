package main

import (
	"log"
	"math/rand"
)

func runRound(w *world, rnd *rand.Rand) bool {
	hits := make([]hitZone, 0)
	moves := make([]move, 0)
	log.Println("Running a round with:", len(w.liveUnits), "Live units")
	for _, u := range w.liveUnits {
		log.Println("Considering unit", u.id, "(", u.kind.name, ")", "at position", u.position)
		v := map[int]bool{u.position.row*w.width + u.position.col: true}
		nc := []coordinates{u.position}
		foundTarget := false
		tc := targetChecker{u.teamNo}
		var t *unit
		if u.roundsSinceLastAttack >= u.kind.attackRate {
			log.Println("\tcan attack")
			for r := 1; (r <= u.kind.attackRange) && !foundTarget; r++ {
				foundTarget, nc, t = scanWithMap(w, v, nc, &tc)
				if !foundTarget {
					for _, c := range nc {
						v[c.row*w.width+c.col] = true
					}
				}
			}
		}
		if foundTarget && (rnd.Intn(100) <= u.kind.attackAccuracy) {
			log.Println("\tfound target AND can hit it")
			hits = append(hits,
				hitZone{
					u.teamNo,
					t,
					u.kind.attackArea,
					u.kind.attackDamage})
			u.roundsSinceLastAttack = 0
		} else if u.kind.moveSpeed > 0 {
			log.Println("\tNo target found, but can move")
			u.roundsSinceLastAttack++
			foundDestination := false
			mc := moveChecker{w, u, nil}
			for r := 1; (r <= u.kind.perception) && !foundDestination; r++ {
				foundDestination, nc, _ = scanWithMap(w, v, nc, &mc)
				if !foundDestination {
					for _, c := range nc {
						v[c.row*w.width+c.col] = true
					}
				}
			}
			if foundDestination {
				np := mc.lastPath[max(1, u.kind.moveSpeed)]
				log.Println("Preparing move for", u.id, "from", u.position, "to", np)
				moves = append(moves, move{u, np})
			} else {
				log.Println("\tNo potential destination found")
			}
		}
	}
	for _, hit := range hits {
		log.Println("Effecting damage", hit.damage, " made on Unit", hit.target.id, "over a radius", hit.radius, "to units of team", 3-hit.teamNo)
		shot := damageInflicter{w, hit.teamNo, hit.damage, hit.radius == 0}
		v := map[int]bool{hit.target.position.row*w.width + hit.target.position.col: true}
		nc := []coordinates{hit.target.position}
		var done bool
		for r := 0; r <= hit.radius; r++ {
			done, nc, _ = scanWithMap(w, v, nc, &shot)
			if done {
				break
			}
			for _, c := range nc {
				v[c.row*w.width+c.col] = true
			}
		}
	}
	for _, move := range moves {
		if move.obj.isLive {
			log.Println("Moving Unit", move.obj.id, "from", move.obj.position, "to", move.to)
			delete(w.terrain[move.obj.position.row][move.obj.position.col].content, move.obj.id)
			w.terrain[move.to.row][move.to.col].content[move.obj.id] = move.obj
			w.terrain[move.obj.position.row][move.obj.position.col].occupiedSlots -= move.obj.kind.slotCount
			w.terrain[move.to.row][move.to.col].occupiedSlots += move.obj.kind.slotCount
			move.obj.position = move.to
		}
	}
	ret := false
	lt := -1
	for _, l := range w.liveUnits {
		if lt == -1 {
			lt = l.teamNo
		} else if lt != l.teamNo {
			ret = true
			break
		}
	}
	return ret
}

func scanWithArray(w *world, visited []coordinates, current []coordinates, predicate cellCondition) (bool, []coordinates, *unit) {
	next := make([]coordinates, 0)
	for _, c := range current {
		var p cell
		p = w.terrain[c.row][c.col]
		if predicate != nil {
			m, u := predicate.verifiedBy(p)
			if m {
				return true, next, u
			}
		}
		for i := 0; i < 6; i++ {
			neigh := neighbour(c, i)
			if isValidPosition(w, neigh) && notInLinear(visited, neigh) {
				next = append(next, neigh)
			}
		}
	}
	return false, next, nil
}

func scanWithMap(w *world, visited map[int]bool, current []coordinates, predicate cellCondition) (bool, []coordinates, *unit) {
	if duplicatesIn(current) {
		panic("duplicates in the current array of coordinates")
	}

	next := make([]coordinates, 0)
	for _, c := range current {
		var p cell
		p = w.terrain[c.row][c.col]
		if predicate != nil {
			m, u := predicate.verifiedBy(p)
			if m {
				return true, next, u
			}
		}
		for i := 0; i < 6; i++ {
			neigh := neighbour(c, i)
			if isValidPosition(w, neigh) &&
				notInConstant(visited, neigh, w.width) &&
				notInLinear(next, neigh) {
				next = append(next, neigh)
			}
		}
	}
	return false, next, nil
}
