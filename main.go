package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"projekt1/graph"
	"time"
)

// Przykład użycia
func main() {

	// wyłączenie logów
	//log.SetOutput(io.Discard)

	verticesAmountPTR := flag.Int("vertices", 10, "Number of vertices")
	mtThreadsPowTwoStartPTR := flag.Int("mt-threads-powers-of-two-start", 2, "for example 2^2 threads")
	mtThreadsPowMultPTR := flag.Int("mt-threads-powers-of-two-mult", 1, "for example 2^2 threads then 2^3 threads to 2^n threads")

	logToFilePTR := flag.Bool("log-to-file", false, "Log to file")
	flag.Parse()

	if *logToFilePTR {
		//save all logs to file
		dateString := time.Now().Format("2006-01-02_15:04:05")
		logFileName := dateString + ".log"
		f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println("Error opening file:", err)
		} else {
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					fmt.Println("Error closing file:", err)
				}
			}(f)
			multi := io.MultiWriter(f, os.Stdout)
			log.SetOutput(multi)
		}
	}

	// Tworzenie grafu pustego
	g := new(graph.AdjMatrixGraph)

	// Wczytanie grafu z pliku
	//err1 := graph.LoadGraphFromFile("17.txt", g)
	//if err1 != nil {
	//	return
	//}
	graph.GenerateRandomGraph(g, *verticesAmountPTR, -1, 1000)

	times := make([]int64, 0)

	// Wywołanie algorytmu brute force dla problemu TSP
	startVertex := 0
	minCostBF, bestPathBF := graph.TSPBruteForce(g, startVertex, &times)
	minCostBFMT, bestPathBFMT := []int{}, [][]int{}
	for i := *mtThreadsPowTwoStartPTR; i < *mtThreadsPowTwoStartPTR+*mtThreadsPowMultPTR; i++ {
		for j := 1; j <= 4; j++ {
			minCostBFMTTMP, bestPathBFMTTMP := graph.TSPBruteForceMT(g, startVertex, &times, 1<<uint(i), j)
			minCostBFMT = append(minCostBFMT, minCostBFMTTMP)
			bestPathBFMT = append(bestPathBFMT, bestPathBFMTTMP)
		}
	}
	minCostBNB, bestPathBNB := graph.TSPBranchAndBound(g, startVertex, &times)
	minCostDP, bestPathDP := graph.TSPDynamicProgramming(g, startVertex, &times)

	log.Println(times)

	log.Printf("Brute force: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBF, g.PathWithWeightsToString(bestPathBF))
	for i := 0; i < len(minCostBFMT); i++ {
		log.Printf("Brute force MT: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBFMT[i], g.PathWithWeightsToString(bestPathBFMT[i]))
	}
	log.Printf("Branch and bound: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBNB, g.PathWithWeightsToString(bestPathBNB))
	log.Printf("Dynamic programming: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostDP, g.PathWithWeightsToString(bestPathDP))
	log.Println(g.ToString())

	err2 := graph.SaveGraphToFile(g, "test2.txt")
	if err2 != nil {
		return
	}
}
