package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	logfile, err := os.Create("app.log")

	if err != nil {
		log.Fatal(err)
	}

	defer func(logfile *os.File) {
		err := logfile.Close()
		if err != nil {

		}
	}(logfile)
	log.SetOutput(logfile)

	var generator BattleGenerator

	generator = DefaultTinyGenerator{1234}
	battleField := generator.create()

	fmt.Println("Starting Units")
	ts := getLiveUnitCounts(battleField)
	for team, ks := range ts {
		fmt.Println("Team", team)
		for kind, uc := range ks {
			fmt.Println(kind.name, ":", uc)
		}
	}

	startTime := time.Now()
	fmt.Println("Start time:", startTime)
	round := 1
	r1 := rand.New(rand.NewSource(123))

	for runRound(&battleField, r1) {
		log.Println("Round:", round, "Elapsed time:", time.Since(startTime))
		// TODO display moves made (by each team)
		// TODO display shots fired (by each team and unit kind)
		// TODO display targets hit within each team
		// TODO display deaths in each team
		// TODO display counts for each team
		round++
	}

	endTime := time.Now()
	fmt.Println("End time  :", endTime, "Elapsed", time.Since(startTime))
	fmt.Println("Units still alive")
	ts = getLiveUnitCounts(battleField)
	for team, ks := range ts {
		fmt.Println("Team", team)
		for kind, uc := range ks {
			fmt.Println(kind.name, ":", uc)
		}
	}
}
