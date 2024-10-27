package graph

import (
	"math"
	"projekt1/timeTrack"
	"time"
)

// Główna funkcja rozwiązująca problem TSP metodą brute force
func TSPBruteForce(g Graph, startVertex int, times *[]int64) (int, []int) {
	startTime := time.Now()
	defer func() {
		*times = append(*times, timeTrack.TimeTrack(startTime, "name"))
	}()
	vertexCount := g.GetVertexCount()
	vertices := make([]int, 0)
	for i := 0; i < vertexCount; i++ {
		if i != startVertex {
			vertices = append(vertices, i)
		}
	}

	// Zmienna przechowująca minimalny koszt trasy
	minPathCost := math.MaxInt
	var bestPath []int

	// Generowanie wszystkich permutacji wierzchołków
	permute(vertices, func(permutation []int) {
		updateBestPath(g, startVertex, permutation, &minPathCost, &bestPath)
	})

	return minPathCost, bestPath
}

// Funkcja do generowania permutacji z użyciem funkcji callback
func permute(vertices []int, callback func([]int)) {
	permuteRecursive(vertices, 0, callback)
}

// Funkcja pomocnicza do generowania permutacji
func permuteRecursive(vertices []int, l int, callback func([]int)) {
	if l == len(vertices)-1 {
		callback(append([]int{}, vertices...)) // Tworzenie kopii permutacji
	} else {
		for i := l; i < len(vertices); i++ {
			vertices[l], vertices[i] = vertices[i], vertices[l]
			permuteRecursive(vertices, l+1, callback)
			vertices[l], vertices[i] = vertices[i], vertices[l]
		}
	}
}

// Nowa funkcja updateBestPath - używana jako callback w permute
func updateBestPath(g Graph, startVertex int, permutation []int, minPathCost *int, bestPath *[]int) {
	// Dodanie wierzchołka początkowego na początku i końcu trasy
	path := append([]int{startVertex}, append(permutation, startVertex)...)

	// Obliczenie kosztu ścieżki przy użyciu metody interfejsu Graph
	pathCost := g.CalculatePathWeight(path)

	// Sprawdzanie, czy znaleziono trasę o niższym koszcie
	if pathCost < *minPathCost {
		*minPathCost = pathCost
		*bestPath = append([]int{}, path...) // Kopiowanie najlepszej trasy
	}
}
