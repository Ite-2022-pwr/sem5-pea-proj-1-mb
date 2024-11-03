package main

import (
	"fmt"
	"log"
	"projekt1/graph"
	"projekt1/timeTrack"
)

// Przykład użycia
func main() {

	// wyłączenie logów
	//log.SetOutput(io.Discard)

	//verticesAmountPTR := flag.Int("vertices", 10, "Number of vertices")
	//doBrunteForcePTR := flag.Bool("brute-force", false, "Do brute force")

	//logToFilePTR := flag.Bool("log-to-file", false, "Log to file")
	//flag.Parse()
	//
	//if *logToFilePTR {
	//	//save all logs to file
	//	dateString := time.Now().Format("2006-01-02_15:04:05")
	//	logFileName := dateString + ".log"
	//	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//
	//	if err != nil {
	//		fmt.Println("Error opening file:", err)
	//	} else {
	//		defer func(f *os.File) {
	//			err := f.Close()
	//			if err != nil {
	//				fmt.Println("Error closing file:", err)
	//			}
	//		}(f)
	//		multi := io.MultiWriter(f, os.Stdout)
	//		log.SetOutput(multi)
	//	}
	//}

	// Tworzenie grafu pustego
	g := new(graph.AdjMatrixGraph)

	// Wczytanie grafu z pliku
	err1 := graph.LoadGraphFromFile("17.txt", g)
	if err1 != nil {
		return
	}
	//graph.GenerateRandomGraph(g, *verticesAmountPTR, -1, 100)
	//graph.GenerateRandomGraph(g, 21, -1, 10000)
	times := make([]int64, 0)

	fmt.Println(g.ToString())

	// Wywołanie algorytmu brute force dla problemu TSP
	startVertex := 0
	minCostDP, bestPathDP := graph.TSPDynamicProgramming(g, startVertex, &times)
	minCostBNB, bestPathBNB := graph.TSPBranchAndBound(g, startVertex, &times)
	//minCostBF, bestPathBF := graph.TSPBruteForce(g, startVertex, &times)

	//if *doBrunteForcePTR {
	//	minCostBF, bestPathBF := graph.TSPBruteForce(g, startVertex, &times)
	//	log.Printf("Brute force: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBF, g.PathWithWeightsToString(bestPathBF))
	//}
	log.Printf("Dynamic programming: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostDP, g.PathWithWeightsToString(bestPathDP))
	log.Printf("Branch and bound: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBNB, g.PathWithWeightsToString(bestPathBNB))
	//log.Printf("Brute force: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBF, g.PathWithWeightsToString(bestPathBF))
	fmt.Println(g.ToString())

	//print times formatted to human readable
	log.Println("Times:")
	for i, t := range times {
		log.Printf("%d: %s\n", i, timeTrack.FormatDurationFromNanoseconds(t))
	}

	//log.Println(times)
	//log.Println(g.ToString())

	err2 := graph.SaveGraphToFile(g, "test2.txt")
	if err2 != nil {
		return
	}
}
