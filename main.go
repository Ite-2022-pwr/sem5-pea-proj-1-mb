package main

import (
	"fmt"
	"projekt1/graph"
)

// Przykład użycia
func main() {

	// wyłączenie logów
	//log.SetOutput(ioutil.Discard)

	// Tworzenie grafu pustego
	g := new(graph.AdjMatrixGraph)

	// Wczytanie grafu z pliku
	//err := graph.LoadGraphFromFile("17.txt", g)
	//if err != nil {
	//	return
	//}
	graph.GenerateRandomGraph(g, 5, -1, 10)

	times := make([]int64, 0)

	// Wywołanie algorytmu brute force dla problemu TSP
	startVertex := 0
	minCostBF, bestPathBF := graph.TSPBruteForce(g, startVertex, &times)
	minCostBNB, bestPathBNB := graph.TSPBranchAndBound(g, startVertex, &times)
	minCostDP, bestPathDP := graph.TSPDynamicProgramming(g, startVertex, &times)

	fmt.Println(times)

	fmt.Printf("Brute force: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBF, g.PathWithWeightsToString(bestPathBF))
	fmt.Printf("Branch and bound: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBNB, g.PathWithWeightsToString(bestPathBNB))
	fmt.Printf("Dynamic programming: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostDP, g.PathWithWeightsToString(bestPathDP))
	fmt.Println(g.ToString())

	err := graph.SaveGraphToFile(g, "test2.txt")
	if err != nil {
		return
	}
}
