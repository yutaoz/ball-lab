package main

import (
	"ballstats/pkg/fileutils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func displayGames(allData []string, score1 int, score2 int) int {
	total := 0
	for _, game := range allData {
		tokens := strings.Split(game, ",")
		//fmt.Println(tokens[0] + " | " + tokens[3] + " | " + tokens[4])
		scorehome, _ := strconv.Atoi(tokens[3])
		scoreaway, _ := strconv.Atoi(tokens[4])

		if scorehome == score1 && scoreaway == score2 || scoreaway == score1 && scorehome == score2 {
			total += 1
			fmt.Println(tokens[0] + " | " + tokens[1] + " | " + tokens[2] + " | " + tokens[3] + " | " + tokens[4])
		}

	}

	return total
}

func main() {
	startTime := time.Now()
	fileSlice := fileutils.GetFiles("..\\data")
	var allData []string
	for _, v := range fileSlice {
		//fmt.Println(v)

		lines := fileutils.ReadFileSync(v)
		allData = append(allData, lines...)
	}

	elapsed := time.Since(startTime)

	_ = displayGames(allData, 98, 90)
	//fmt.Println(allData)
	fmt.Println("Solution 2 took:", elapsed)
}
