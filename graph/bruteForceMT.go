package graph

import (
	"log"
	"math"
	"projekt1/timeTrack"
	"strconv"
	"sync"
	"time"
)

// Główna funkcja rozwiązująca problem TSP metodą brute-force z wykorzystaniem goroutines
func TSPBruteForceMT(g Graph, startVertex int, times *[]int64, goroutines int, prefixLen int) (int, []int) {
	// Mierzenie czasu rozpoczęcia funkcji
	startTime := time.Now()
	defer func() {
		elapsedTime := timeTrack.TimeTrack(startTime, "brute force MT"+strconv.Itoa(goroutines)+"pref"+strconv.Itoa(prefixLen)+", liczba wierzchołków: "+strconv.Itoa(g.GetVertexCount()))
		*times = append(*times, elapsedTime)
		log.Println("Czas wykonania TSPBruteForceMT:", elapsedTime, "ns")
	}()

	vertexCount := g.GetVertexCount()

	// Log rozpoczęcia przetwarzania metodą brute force MT
	log.Println("Rozpoczęcie Brute Force MT dla wierzchołka początkowego:", startVertex, "z liczbą wierzchołków:", vertexCount)

	// Lista wierzchołków bez wierzchołka początkowego
	vertices := make([]int, 0)
	for i := 0; i < vertexCount; i++ {
		if i != startVertex {
			vertices = append(vertices, i)
		}
	}

	// Generowanie wszystkich permutacji prefiksów
	prefixes := [][]int{}
	prefixLength := prefixLen // Można dostosować długość prefiksu w zależności od potrzeb
	generatePrefixesMT(vertices, []int{}, prefixLength, &prefixes)

	log.Println("Liczba wygenerowanych prefiksów:", len(prefixes))
	// log.Println("Prefiksy:", prefixes) // Można odkomentować dla małych grafów

	// Kanał do przesyłania prefiksów
	prefixChan := make(chan []int, len(prefixes))

	// Wypełnienie kanału prefiksami
	for _, prefix := range prefixes {
		prefixChan <- prefix
	}
	close(prefixChan)

	// Zmienna przechowująca minimalny koszt trasy
	minPathCost := math.MaxInt
	var bestPath []int

	// Mutex do ochrony wspólnych zasobów
	var mutex sync.Mutex

	var wg sync.WaitGroup

	// Zmienna do liczenia wszystkich permutacji
	totalPermutations := 0
	var permMutex sync.Mutex

	// Uruchomienie goroutines
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for prefix := range prefixChan {
				log.Println("Goroutine", workerID, "z prefiksem", prefix, "rozpoczęła pracę")

				remainingVertices := subtract(vertices, prefix)

				// Lokalny minimalny koszt i najlepsza ścieżka
				localMinCost := math.MaxInt
				var localBestPath []int

				// Lokalna liczba permutacji
				localPermutations := 0

				// Generowanie permutacji pozostałych wierzchołków
				permuteRecursiveMT(remainingVertices, 0, func(permutation []int) {
					localPermutations++

					// Co pewien krok logujemy postęp
					if localPermutations%1000 == 0 {
						log.Println("Goroutine", workerID, "z prefiksem", prefix, "przetworzyła", localPermutations, "permutacji")
					}

					// Budowanie pełnej ścieżki
					path := append([]int{startVertex}, append(prefix, permutation...)...)
					path = append(path, startVertex)

					// Obliczanie kosztu ścieżki
					pathCost := g.CalculatePathWeight(path)

					// Aktualizacja lokalnego minimalnego kosztu
					if pathCost < localMinCost {
						localMinCost = pathCost
						localBestPath = append([]int{}, path...)
						log.Println("Goroutine", workerID, "z prefiksem", prefix, "znalazła nową lokalnie najlepszą trasę o koszcie:", localMinCost)
					}
				})

				// Aktualizacja globalnej liczby permutacji
				permMutex.Lock()
				totalPermutations += localPermutations
				permMutex.Unlock()

				// Aktualizacja globalnych zmiennych
				mutex.Lock()
				if localMinCost < minPathCost {
					minPathCost = localMinCost
					bestPath = localBestPath
					log.Println("Goroutine", workerID, "z prefiksem", prefix, "znalazła nową globalnie najlepszą trasę o koszcie:", minPathCost)
				}
				mutex.Unlock()

				log.Println("Goroutine", workerID, "z prefiksem", prefix, "zakończyła pracę. Przetworzone permutacje:", localPermutations)
			}
		}(i)
	}

	wg.Wait()

	// Log zakończenia przetwarzania metodą brute force MT
	log.Println("Zakończono Brute Force MT dla wierzchołka początkowego:", startVertex)
	log.Println("Minimalny koszt znalezionej trasy:", minPathCost)
	log.Println("Najlepsza znaleziona ścieżka:", bestPath)
	log.Println("Łączna liczba przetworzonych permutacji:", totalPermutations)

	if minPathCost == math.MaxInt {
		return -1, nil
	}

	return minPathCost, bestPath
}

// Funkcja generująca wszystkie możliwe prefiksy o zadanej długości
func generatePrefixesMT(vertices []int, prefix []int, k int, prefixes *[][]int) {
	if len(prefix) == k || len(prefix) == len(vertices) {
		*prefixes = append(*prefixes, append([]int{}, prefix...))
		return
	}
	for i, v := range vertices {
		newPrefix := append(prefix, v)
		remaining := append([]int{}, vertices[:i]...)
		remaining = append(remaining, vertices[i+1:]...)
		generatePrefixesMT(remaining, newPrefix, k, prefixes)
	}
}

// Funkcja generująca permutacje rekurencyjnie
func permuteRecursiveMT(vertices []int, l int, callback func([]int)) {
	if l == len(vertices)-1 {
		callback(append([]int{}, vertices...)) // Tworzenie kopii permutacji
	} else {
		for i := l; i < len(vertices); i++ {
			vertices[l], vertices[i] = vertices[i], vertices[l]
			permuteRecursiveMT(vertices, l+1, callback)
			vertices[l], vertices[i] = vertices[i], vertices[l]
		}
	}
}

// Funkcja odejmująca elementy b od a
func subtract(a, b []int) []int {
	result := []int{}
	m := make(map[int]bool)
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			result = append(result, v)
		}
	}
	return result
}
