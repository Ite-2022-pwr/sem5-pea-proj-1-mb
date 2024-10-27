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
	graph.GenerateRandomGraph(g, *verticesAmountPTR, -1, 100)

	times := make([]int64, 0)

	// Wywołanie algorytmu brute force dla problemu TSP
	startVertex := 0
	minCostBNB, bestPathBNB := graph.TSPBranchAndBound(g, startVertex, &times)
	minCostDP, bestPathDP := graph.TSPDynamicProgramming(g, startVertex, &times)
	minCostBF, bestPathBF := graph.TSPBruteForce(g, startVertex, &times)

	log.Println(times)

	log.Printf("Brute force: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBF, g.PathWithWeightsToString(bestPathBF))
	log.Printf("Branch and bound: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBNB, g.PathWithWeightsToString(bestPathBNB))
	log.Printf("Dynamic programming: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostDP, g.PathWithWeightsToString(bestPathDP))
	log.Println(g.ToString())

	err2 := graph.SaveGraphToFile(g, "test2.txt")
	if err2 != nil {
		return
	}
}
