package main

import (
	"ballstats/pkg/fileutils"
	"fmt"
	"strings"
)

func checkMatchup(yearData []string, t1 string, t2 string) {
	for _, game := range yearData {
		gameData := strings.Split(game, ",")
		homeTeam := gameData[1]
		awayTeam := gameData[2]
		homeScore := gameData[3]
		awayScore := gameData[4]

		// _, ok := winner[homeTeam]

		// if !ok {
		// 	winner[homeTeam] = 0
		// 	winner[awayTeam] = 0
		// }

		if (homeTeam == t1 || homeTeam == t2) && (awayTeam == t1 || awayTeam == t2) {
			fmt.Println("Home - " + homeTeam + "| " + homeScore + " : " + awayScore + " |" + awayTeam + " - Away")
		}

	}
}
func main() {
	fileSlice := fileutils.GetFilesYear("..\\data", 2024)
	var yearData []string
	for _, v := range fileSlice {
		print(v + "\n")
		lines := fileutils.ReadFileSync(v)
		yearData = append(yearData, lines...)

	}

	checkMatchup(yearData, "Boston Celtics", "Indiana Pacers")

}
