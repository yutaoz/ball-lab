package main

import (
	"ballstats/pkg/fileutils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	dataChan := make(chan []string, len(fileSlice))

	for _, v := range fileSlice {
		//fmt.Println(v)
		wg.Add(1)
		go fileutils.ReadFile(v, &wg, dataChan)
	}

	go func() {
		wg.Wait()
		close(dataChan)
	}()

	var allData []string
	for data := range dataChan {
		allData = append(allData, data...)
	}
	sort.Strings(allData)
	elapsed := time.Since(startTime)
	total := displayGames(allData, 113, 98)
	fmt.Println(total)
	fmt.Println("Solution 1 took:", elapsed)
}
