package graph

import (
	"log"
	"math"
	"projekt1/timeTrack"
	"strconv"
	"time"
)

// Główna funkcja rozwiązująca problem komiwojażera metodą brute-force.
func TSPBruteForce(g Graph, startVertex int, times *[]int64) (int, []int) {
	startTime := time.Now()
	defer func() {
		*times = append(*times, timeTrack.TimeTrack(startTime, "brute-force, liczba wierzchołków: "+strconv.Itoa(g.GetVertexCount())))
	}()

	vertexCount := g.GetVertexCount()

	log.Println("Rozpoczęcie Brute-Force dla wierzchołka początkowego:", startVertex, "z liczbą wierzchołków:", vertexCount)

	minPathCost := math.MaxInt             // Inicjalizacja minimalnego kosztu.
	currentPath := make([]int, 0)          // Aktualna ścieżka.
	visited := make([]bool, vertexCount)   // Tablica odwiedzonych wierzchołków.
	bestPath := make([]int, vertexCount+1) // Najlepsza znaleziona ścieżka (z powrotem do startu).

	// Oznaczamy wierzchołek początkowy jako odwiedzony i dodajemy go do aktualnej ścieżki.
	visited[startVertex] = true
	currentPath = append(currentPath, startVertex)

	// Rozpoczynamy rekurencyjne przeszukiwanie wszystkich możliwych ścieżek.
	bruteForce(g, startVertex, visited, 0, &minPathCost, currentPath, bestPath)

	log.Println("Zakończono Brute-Force dla wierzchołka początkowego:", startVertex, "Minimalny koszt:", minPathCost)

	return minPathCost, bestPath
}

// Rekurencyjna funkcja przeszukująca wszystkie możliwe ścieżki.
func bruteForce(g Graph, currentVertex int, visited []bool, currentCost int, minPathCost *int, currentPath, bestPath []int) {
	vertexCount := g.GetVertexCount()

	// Jeśli odwiedziliśmy wszystkie wierzchołki, sprawdzamy powrót do wierzchołka startowego.
	if len(currentPath) == vertexCount {
		// Pobieramy krawędź z ostatniego wierzchołka do wierzchołka startowego.
		edge := g.GetEdge(currentVertex, currentPath[0])
		if edge.Weight != g.GetNoEdgeValue() {
			totalCost := currentCost + edge.Weight
			// Sprawdzamy, czy całkowity koszt jest mniejszy od dotychczasowego minimalnego kosztu.
			if totalCost < *minPathCost {
				*minPathCost = totalCost
				// Tworzymy tymczasową ścieżkę dodając powrót do wierzchołka startowego.
				tempPath := append(currentPath, currentPath[0])
				// Kopiujemy aktualną ścieżkę jako najlepszą znalezioną.
				copy(bestPath, tempPath)
			}
		}
		return
	}

	// Przechodzimy przez wszystkie wierzchołki grafu.
	for nextVertex := 0; nextVertex < vertexCount; nextVertex++ {
		if !visited[nextVertex] {
			// Sprawdzamy, czy istnieje krawędź z bieżącego wierzchołka do nextVertex.
			edge := g.GetEdge(currentVertex, nextVertex)
			if edge.Weight != g.GetNoEdgeValue() {
				// Oznaczamy nextVertex jako odwiedzony i dodajemy go do aktualnej ścieżki.
				visited[nextVertex] = true
				currentPath = append(currentPath, nextVertex)
				newCost := currentCost + edge.Weight

				// Rekurencyjne wywołanie dla nextVertex.
				bruteForce(g, nextVertex, visited, newCost, minPathCost, currentPath, bestPath)

				// Cofamy zmiany (backtracking): usuwamy nextVertex z aktualnej ścieżki i oznaczamy go jako nieodwiedzonego.
				visited[nextVertex] = false
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
	}
}
