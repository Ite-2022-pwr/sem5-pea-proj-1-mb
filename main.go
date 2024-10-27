package main

import (
	"fmt"
	"projekt1/graph"
)

// Przykład użycia
func main() {
	// Tworzenie grafu o 4 wierzchołkach z wartością -1 dla braku krawędzi
	g := new(graph.AdjMatrixGraph)

	graph.GenerateRandomGraph(g, 10, -1, 50)

	times := make([]int64, 0)

	// Wywołanie algorytmu brute force dla problemu TSP
	startVertex := 0
	minCost, bestPath := graph.TSPBruteForce(g, startVertex, &times)

	fmt.Println(times)

	fmt.Printf("Minimalny koszt trasy: %d\n", minCost)
	fmt.Printf("Najlepsza trasa: %v\n", bestPath)

	fmt.Println(g.ToString())

	err := graph.SaveGraphToFile(g, "test2.txt")
	if err != nil {
		return
	}
}
