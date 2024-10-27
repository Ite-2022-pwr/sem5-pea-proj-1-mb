package graph

import (
	"log"
	"math"
	"projekt1/timeTrack"
	"strconv"
	"time"
)

// Główna funkcja rozwiązująca problem TSP metodą programowania dynamicznego (Held-Karp)
func TSPDynamicProgramming(g Graph, startVertex int, times *[]int64) (int, []int) {
	// Mierzenie czasu rozpoczęcia funkcji
	startTime := time.Now()
	defer func() {
		*times = append(*times, timeTrack.TimeTrack(startTime, "dynamic programming, number of vertices: "+strconv.Itoa(g.GetVertexCount())))
	}()

	vertexCount := g.GetVertexCount()
	allVisited := (1 << vertexCount) - 1 // Maskowanie dla wszystkich wierzchołków odwiedzonych

	// Tworzenie mapy przechowującej koszty częściowych rozwiązań
	memo := make([][]int, vertexCount)
	for i := range memo {
		memo[i] = make([]int, 1<<vertexCount)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt // Inicjalizacja maksymalnym kosztem
		}
	}

	// Inicjalizacja wierzchołka początkowego
	memo[startVertex][1<<startVertex] = 0

	// Dynamiczne przeliczanie wartości dla podproblemów
	for subset := 1; subset <= allVisited; subset++ {
		if (subset & (1 << startVertex)) == 0 {
			continue // Pomijamy zbiory, które nie zawierają startVertex
		}

		for currentVertex := 0; currentVertex < vertexCount; currentVertex++ {
			if (subset&(1<<currentVertex)) == 0 || currentVertex == startVertex {
				continue // Pomijamy wierzchołki, które nie są w bieżącym podzbiorze
			}

			previousSubset := subset ^ (1 << currentVertex) // Podzbiór bez bieżącego wierzchołka

			// Szukamy minimalnego kosztu przejścia do bieżącego wierzchołka
			for prevVertex := 0; prevVertex < vertexCount; prevVertex++ {
				if (previousSubset & (1 << prevVertex)) == 0 {
					continue // Pomijamy, jeśli prevVertex nie jest w zbiorze
				}

				edge := g.GetEdge(prevVertex, currentVertex)
				if edge.Weight == g.GetNoEdgeValue() {
					continue // Pomijamy, jeśli nie ma krawędzi
				}

				if memo[prevVertex][previousSubset] == math.MaxInt {
					continue // Pomijamy, jeśli nie ma wartości dla tego podzbioru
				}

				newCost := memo[prevVertex][previousSubset] + edge.Weight
				if newCost < memo[currentVertex][subset] {
					memo[currentVertex][subset] = newCost
					log.Println("Aktualizacja kosztu dla podzbioru:", subset, "i wierzchołka:", currentVertex, "Nowy koszt:", newCost)
				}
			}
		}
	}

	// Znalezienie minimalnej ścieżki powrotnej do wierzchołka startowego
	minCost := math.MaxInt
	lastVertex := -1
	for vertex := 0; vertex < vertexCount; vertex++ {
		if vertex == startVertex {
			continue
		}
		edge := g.GetEdge(vertex, startVertex)
		if edge.Weight == g.GetNoEdgeValue() {
			continue
		}

		if memo[vertex][allVisited] == math.MaxInt {
			continue // Pomijamy, jeśli nie ma rozwiązania dla tego podzbioru
		}

		totalCost := memo[vertex][allVisited] + edge.Weight
		if totalCost < minCost {
			minCost = totalCost
			lastVertex = vertex
		}
		log.Println("Sprawdzono koszt powrotu z wierzchołka:", vertex, "Koszt całkowity:", totalCost)
	}

	// Odtwarzanie najlepszej ścieżki
	if lastVertex == -1 {
		return -1, nil // Jeśli nie znaleziono żadnej ścieżki
	}

	bestPath := []int{startVertex}
	currentSubset := allVisited
	currentVertex := lastVertex

	for currentVertex != startVertex {
		bestPath = append(bestPath, currentVertex)
		previousSubset := currentSubset ^ (1 << currentVertex)

		// Znajdujemy poprzedni wierzchołek na podstawie zapamiętanych kosztów
		for prevVertex := 0; prevVertex < vertexCount; prevVertex++ {
			if (previousSubset & (1 << prevVertex)) == 0 {
				continue
			}

			if memo[prevVertex][previousSubset] == math.MaxInt {
				continue
			}

			if memo[prevVertex][previousSubset]+g.GetEdge(prevVertex, currentVertex).Weight == memo[currentVertex][currentSubset] {
				currentVertex = prevVertex
				currentSubset = previousSubset
				break
			}
		}
	}

	// Dodajemy wierzchołek startowy na koniec trasy, aby wrócić do punktu startu
	bestPath = append(bestPath, startVertex)

	log.Println("Znaleziono najlepszą ścieżkę o koszcie:", minCost, "Ścieżka:", bestPath)

	return minCost, bestPath
}
