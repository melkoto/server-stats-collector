package main

import (
	"fmt"
	"os"
	"time"

	"github.com/melkoto/server-stats-collector/internal/checker"
	"github.com/melkoto/server-stats-collector/internal/fetcher"
	"github.com/melkoto/server-stats-collector/internal/parser"
)

func main() {
	url := "http://srv.msk01.gigacorp.local/_stats"
	errorCount := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		statsStr, err := fetcher.FetchStats(url)
		if err != nil {
			errorCount++
			if errorCount >= 3 {
				fmt.Println("Unable to fetch server statistic.")
				os.Stdout.Sync()
			}
			continue
		}

		arr, err := parser.ParseStats(statsStr)
		if err != nil {
			errorCount++
			if errorCount >= 3 {
				fmt.Println("Unable to fetch server statistic.")
				os.Stdout.Sync()
			}
			continue
		}

		errorCount = 0

		loadAvg := arr[0]
		totalMem := arr[1]
		usedMem := arr[2]
		totalDisk := arr[3]
		usedDisk := arr[4]
		totalNet := arr[5]
		usedNet := arr[6]

		if msg := checker.CheckLoad(loadAvg); msg != "" {
			fmt.Println(msg)
			os.Stdout.Sync()
		}
		if msg := checker.CheckMemory(totalMem, usedMem); msg != "" {
			fmt.Println(msg)
			os.Stdout.Sync()
		}
		if msg := checker.CheckDisk(totalDisk, usedDisk); msg != "" {
			fmt.Println(msg)
			os.Stdout.Sync()
		}
		if msg := checker.CheckNetwork(totalNet, usedNet); msg != "" {
			fmt.Println(msg)
			os.Stdout.Sync()
		}
	}
}
