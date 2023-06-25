package main

func getLiveUnitCounts(w world) map[int]map[unitKind]int {
	stat := make(map[int]map[unitKind]int)
	for _, u := range w.liveUnits {
		teamStat, exists := stat[u.teamNo]
		if !exists {
			teamStat = make(map[unitKind]int)
			stat[u.teamNo] = teamStat
		}
		_, exists = teamStat[u.kind]
		if !exists {
			teamStat[u.kind] = 0
		}
		teamStat[u.kind]++
	}
	return stat
}
