package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
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

	startTime := time.Now()
	fmt.Println("Start time:", startTime)
	round := 1
	for runRound(&battleField, r1) {
		fmt.Println("Round:", round, "Elapsed time:", time.Since(startTime))
		// TODO display moves made (by each team)
		// TODO display shots fired (by each team and unit kind)
		// TODO display targets hit within each team
		// TODO display deaths in each team
		// TODO display counts for each team
		round++
	}
	fmt.Println("End time  :", time.Now())
}
