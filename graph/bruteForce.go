package graph

import (
	"log"
	"math"
	"projekt1/timeTrack"
	"strconv"
	"time"
)

// Główna funkcja rozwiązująca problem TSP metodą brute force
func TSPBruteForce(g Graph, startVertex int, times *[]int64) (int, []int) {
	// Mierzenie czasu rozpoczęcia funkcji
	startTime := time.Now()
	defer func() {
		*times = append(*times, timeTrack.TimeTrack(startTime, "brute force, number of vertices: "+strconv.Itoa(g.GetVertexCount())))
	}()

	vertexCount := g.GetVertexCount()

	// Log rozpoczęcia przetwarzania metodą brute force
	log.Println("Rozpoczęcie Brute Force dla wierzchołka początkowego:", startVertex, "z liczba wierzchołków:", vertexCount)

	vertices := make([]int, 0)
	for i := 0; i < vertexCount; i++ {
		if i != startVertex {
			vertices = append(vertices, i)
		}
	}

	// Zmienna przechowująca minimalny koszt trasy
	minPathCost := math.MaxInt
	var bestPath []int

	permutations := 0

	// Generowanie wszystkich permutacji wierzchołków
	permute(vertices, func(permutation []int) {
		updateBestPath(g, startVertex, permutation, &minPathCost, &bestPath)
	}, &permutations)

	// Log zakończenia przetwarzania metodą brute force
	log.Println("Zakończono Brute Force dla wierzchołka początkowego:", startVertex, "Minimalny koszt:", minPathCost)
	log.Println("Wszystkich permutacji:", permutations)

	return minPathCost, bestPath
}

// Funkcja do generowania permutacji z użyciem funkcji callback
func permute(vertices []int, callback func([]int), permutations *int) {
	permuteRecursive(vertices, 0, callback, permutations)
}

// Funkcja pomocnicza do generowania permutacji
func permuteRecursive(vertices []int, level int, callback func([]int), permutations *int) {
	if level == len(vertices)-1 {
		*permutations++
		//if *permutations%1000 == 0 {
		//	//log.Println("Utworzono permutację:", *permutations) // Logowanie utworzenia nowej permutacji
		//}
		callback(append([]int{}, vertices...)) // Tworzenie kopii permutacji
	} else {
		for i := level; i < len(vertices); i++ {
			vertices[level], vertices[i] = vertices[i], vertices[level]
			permuteRecursive(vertices, level+1, callback, permutations)
			vertices[level], vertices[i] = vertices[i], vertices[level]
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
		//log.Println("Znaleziono nową najlepszą trasę o koszcie:", pathCost, "\nŚcieżka:", path) // Logowanie nowej lepszej trasy
		*minPathCost = pathCost
		*bestPath = append([]int{}, path...) // Kopiowanie najlepszej trasy
	}
}
