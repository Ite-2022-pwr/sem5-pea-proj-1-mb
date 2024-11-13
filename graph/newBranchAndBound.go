package graph

import (
	"container/heap"
	"log"
	"math"
	"projekt1/timeTrack"
	"strconv"
	"time"
)

// Funkcja oblicza początkowe dolne ograniczenie oraz tworzy tablicę minimalnych kosztów krawędzi wychodzących z każdego wierzchołka.
func calculateStartLowerBound(g Graph) (int, []int) {
	lowerBound := 0
	minEdgeLookup := make([]int, g.GetVertexCount())
	for i := 0; i < g.GetVertexCount(); i++ {
		// Dodajemy minimalny koszt krawędzi wychodzącej z wierzchołka i do dolnego ograniczenia.
		lowerBound += g.GetMinEdgeFromWeight(i)
		// Zapamiętujemy minimalny koszt krawędzi wychodzącej z wierzchołka i.
		minEdgeLookup[i] = g.GetMinEdgeFromWeight(i)
	}
	return lowerBound, minEdgeLookup
}

// Funkcja oblicza dolne ograniczenie dla przejścia z bieżącego wierzchołka do następnego.
func calculateLowerBound(g Graph, currentBNBNode BNBNode, nextVertex int, minEdgeLookup []int) int {
	// Aktualizujemy dolne ograniczenie, odejmując minimalny koszt krawędzi wychodzącej z bieżącego wierzchołka
	// i dodając koszt rzeczywistej krawędzi do następnego wierzchołka.
	return currentBNBNode.lowerBound - minEdgeLookup[currentBNBNode.vertex] + g.GetEdge(currentBNBNode.vertex, nextVertex).Weight
}

// Główna funkcja rozwiązująca problem komiwojażera metodą Branch and Bound.
func TSPNewBranchAndBound(g Graph, startVertex int, times *[]int64) (int, []int) {
	log.Println("Rozpoczęcie Branch and Bound dla wierzchołka początkowego:", startVertex, "z liczbą wierzchołków:", g.GetVertexCount())
	startTime := time.Now()
	defer func() {
		*times = append(*times, timeTrack.TimeTrack(startTime, "branch and bound, liczba wierzchołków: "+strconv.Itoa(g.GetVertexCount())))
	}()

	vertexCount := g.GetVertexCount()

	// Obliczamy początkowe dolne ograniczenie oraz minimalne koszty krawędzi wychodzących.
	lowerBound, minEdgeLookup := calculateStartLowerBound(g)
	minPathCost := math.MaxInt
	currentPath := make([]int, 0)                                     // Aktualna ścieżka.
	visited := make([]bool, vertexCount)                              // Tablica odwiedzonych wierzchołków.
	bestPath := make([]int, vertexCount+1)                            // Najlepsza znaleziona ścieżka.
	startNode := BNBNode{vertex: startVertex, lowerBound: lowerBound} // Inicjalizacja początkowego węzła.

	// Rozpoczynamy rekurencyjne przeszukiwanie drzewa rozwiązań.
	newBranchAndBound(g, startNode, visited, &minPathCost, currentPath, bestPath, minEdgeLookup)

	return minPathCost, bestPath
}

// Rekurencyjna funkcja realizująca algorytm Branch and Bound.
func newBranchAndBound(g Graph, currentBNBNode BNBNode, visited []bool, minPathCost *int, currentPath, bestPath, minEdgeLookup []int) {
	// Dodajemy bieżący wierzchołek do aktualnej ścieżki.
	currentPath = append(currentPath, currentBNBNode.vertex)
	// Oznaczamy bieżący wierzchołek jako odwiedzony.
	visited[currentBNBNode.vertex] = true
	// Tworzymy listę nieodwiedzonych węzłów do dalszego przeszukiwania.
	notVisitedBNBNodes := make([]BNBNode, 0)

	// Przechodzimy przez wszystkie wierzchołki grafu.
	for i := 0; i < g.GetVertexCount(); i++ {
		if !visited[i] {
			// Obliczamy dolne ograniczenie dla przejścia do wierzchołka i.
			newLowerBound := calculateLowerBound(g, currentBNBNode, i, minEdgeLookup)
			// Dodajemy nowy węzeł do listy nieodwiedzonych węzłów.
			notVisitedBNBNodes = append(notVisitedBNBNodes, BNBNode{vertex: i, lowerBound: newLowerBound})
		}
	}

	// Tworzymy kopiec z nieodwiedzonych węzłów, aby zawsze wybierać ten z najniższym dolnym ograniczeniem.
	notVisitedBNBNodesHeap := NewBNBNodeHeapByInit(notVisitedBNBNodes)

	// Jeśli nie ma więcej węzłów do odwiedzenia (osiągnięto liść drzewa).
	if notVisitedBNBNodesHeap.Len() == 0 {
		// Obliczamy dolne ograniczenie dla powrotu do wierzchołka startowego.
		returnToStartLowerBound := calculateLowerBound(g, currentBNBNode, currentPath[0], minEdgeLookup)
		// Sprawdzamy, czy znaleziony koszt jest mniejszy od dotychczasowego minimalnego kosztu.
		if returnToStartLowerBound < *minPathCost {
			*minPathCost = returnToStartLowerBound
			// Dodajemy powrót do wierzchołka startowego w aktualnej ścieżce.
			currentPath = append(currentPath, currentPath[0])
			// Kopiujemy aktualną ścieżkę jako najlepszą znalezioną.
			copy(bestPath, currentPath)
		}
	} else {
		// Przechodzimy przez dostępne nieodwiedzone węzły.
		for notVisitedBNBNodesHeap.Len() > 0 {
			// Pobieramy węzeł z najniższym dolnym ograniczeniem.
			nextBNBNode := heap.Pop(notVisitedBNBNodesHeap).(BNBNode)
			// Jeśli dolne ograniczenie jest mniejsze od obecnego minimalnego kosztu, kontynuujemy przeszukiwanie.
			if nextBNBNode.lowerBound < *minPathCost {
				// Rekurencyjne wywołanie dla następnego węzła.
				newBranchAndBound(g, nextBNBNode, visited, minPathCost, currentPath, bestPath, minEdgeLookup)
			}
		}
	}
	// Cofamy oznaczenie bieżącego wierzchołka jako odwiedzonego (backtracking).
	visited[currentBNBNode.vertex] = false
	// Usuwamy bieżący wierzchołek z aktualnej ścieżki.
	currentPath = currentPath[:len(currentPath)-1]
}
