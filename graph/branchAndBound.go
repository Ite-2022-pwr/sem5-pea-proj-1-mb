package graph

import (
	"log"
	"math"
	"projekt1/timeTrack"
	"strconv"
	"time"
)

// Główna funkcja rozwiązująca problem TSP metodą Branch and Bound
func TSPBranchAndBound(g Graph, startVertex int, times *[]int64) (int, []int) {
	log.Println("Rozpoczęcie Branch and Bound dla wierzchołka początkowego:", startVertex, "liczba wierchołków:", g.GetVertexCount())
	// Mierzenie czasu rozpoczęcia funkcji
	startTime := time.Now()
	defer func() {
		*times = append(*times, timeTrack.TimeTrack(startTime, "branch and bound, number of vertices: "+strconv.Itoa(g.GetVertexCount())))
	}()

	vertexCount := g.GetVertexCount()

	// Inicjalizacja zmiennych
	minPathCost := math.MaxInt                // Przechowujemy minimalny koszt trasy
	currentPath := make([]int, vertexCount+1) // Aktualna ścieżka (plus powrót do startowego wierzchołka)
	visited := make([]bool, vertexCount)      // Śledzenie odwiedzonych wierzchołków

	// Startujemy z wierzchołkiem początkowym
	currentPath[0] = startVertex
	visited[startVertex] = true

	// Rozpocznij procedurę Branch and Bound
	bestPath := make([]int, vertexCount+1)
	branchAndBound(g, startVertex, 1, 0, visited, &minPathCost, currentPath, bestPath)

	return minPathCost, bestPath
}

// Funkcja Branch and Bound do przeszukiwania możliwych ścieżek
func branchAndBound(g Graph, currentVertex, level, currentCost int, visited []bool, minPathCost *int, currentPath, bestPath []int) {
	vertexCount := g.GetVertexCount()

	// Logujemy rozpoczęcie przeszukiwania dla nowego poziomu
	//log.Println("Przeszukiwanie na poziomie:", level, "Aktualny wierzchołek:", currentVertex, "Aktualny koszt:", currentCost)

	// Jeśli osiągnęliśmy poziom równy liczbie wierzchołków, sprawdzamy połączenie powrotne
	if level == vertexCount {
		// Sprawdź czy istnieje krawędź z ostatniego wierzchołka do wierzchołka startowego
		edge := g.GetEdge(currentVertex, currentPath[0])
		if edge.Weight != g.GetNoEdgeValue() {
			totalCost := currentCost + edge.Weight

			// Aktualizuj najlepszą ścieżkę, jeśli koszt jest niższy
			if totalCost < *minPathCost {
				*minPathCost = totalCost

				// Kopiujemy aktualną ścieżkę do najlepszej ścieżki
				copy(bestPath, currentPath)

				// Dodajemy na końcu wierzchołek początkowy, aby utworzyć pełny cykl
				bestPath[vertexCount] = currentPath[0]

				// Logujemy znalezienie nowej lepszej ścieżki
				//log.Println("Znaleziono nową lepszą ścieżkę o koszcie:", totalCost, "Ścieżka:", bestPath)
			}
		}
		return
	}

	// Przechodzimy przez wszystkie wierzchołki w celu dalszego przeszukiwania
	for nextVertex := 0; nextVertex < vertexCount; nextVertex++ {
		if !visited[nextVertex] {
			// Sprawdź, czy istnieje krawędź do `nextVertex`
			edge := g.GetEdge(currentVertex, nextVertex)
			if edge.Weight != g.GetNoEdgeValue() {
				// Oznacz wierzchołek jako odwiedzony
				visited[nextVertex] = true
				currentPath[level] = nextVertex

				// Oblicz nowy koszt i sprawdź, czy warto kontynuować
				newCost := currentCost + edge.Weight
				if newCost < *minPathCost {
					// Kontynuuj przeszukiwanie z `nextVertex`
					branchAndBound(g, nextVertex, level+1, newCost, visited, minPathCost, currentPath, bestPath)
				}

				// Odznacz wierzchołek jako odwiedzony
				visited[nextVertex] = false
			}
		}
		//runtime.GC()
	}
}
