package control

import (
	"fmt"
	"log"
	"projekt1/graph"
	"projekt1/timeTrack"
)

func Menu() {
	g := new(graph.AdjMatrixGraph)
	for {
		fmt.Println("1. Wczytaj graf z pliku")
		fmt.Println("2. Wygeneruj losowy graf")
		fmt.Println("3. Wyświetl graf")
		fmt.Println("4. Algorytm dynamiczny")
		fmt.Println("5. Algorytm Branch and Bound tylko z górnym ograniczeniem")
		fmt.Println("6. Algorytm Branch and Bound z dolnym ograniczeniem")
		fmt.Println("7. Algorytm Brute Force")
		fmt.Println("8. Zapisz graf do pliku")
		fmt.Println("9. Ustaw wartość braku krawędzi")
		fmt.Println("0. Wyjście")
		fmt.Print("Wybierz opcję: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}
		switch choice {
		case 1:
			fmt.Print("Podaj nazwę pliku: ")
			var fileName string
			_, err := fmt.Scanln(&fileName)
			if err != nil {
				return
			}
			err = graph.LoadGraphFromFile(fileName, g)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Print("Podaj liczbę wierzchołków: ")
			var vertexCount int
			_, err := fmt.Scanln(&vertexCount)
			if err != nil {
				return
			}
			graph.GenerateRandomGraph(g, vertexCount, -1, 100)
		case 3:
			fmt.Printf("wartość braku krawędzi: %d\n", g.GetNoEdgeValue())
			fmt.Println(g.ToString())
		case 4:
			startVertex := 0
			times := make([]int64, 0)
			minCostDP, bestPathDP := graph.TSPDynamicProgramming(g, startVertex, &times)
			log.Println("Ścieżka:", bestPathDP)
			fmt.Printf("Dynamic programming: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostDP, g.PathWithWeightsToString(bestPathDP))
			fmt.Printf("Czas: %s\n", timeTrack.FormatDurationFromNanoseconds(times[0]))
		case 5:
			startVertex := 0
			times := make([]int64, 0)
			minCostBNB, bestPathBNB := graph.TSPBranchAndBound(g, startVertex, &times)
			log.Println("Ścieżka:", bestPathBNB)
			fmt.Printf("Branch and bound: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBNB, g.PathWithWeightsToString(bestPathBNB))
			fmt.Printf("Czas: %s\n", timeTrack.FormatDurationFromNanoseconds(times[0]))
		case 6:
			startVertex := 0
			times := make([]int64, 0)
			minCostNBNB, bestPathNBNB := graph.TSPNewBranchAndBound(g, startVertex, &times)
			log.Println("Ścieżka:", bestPathNBNB)
			fmt.Printf("Branch and bound: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostNBNB, g.PathWithWeightsToString(bestPathNBNB))
			fmt.Printf("Czas: %s\n", timeTrack.FormatDurationFromNanoseconds(times[0]))
		case 7:
			startVertex := 0
			times := make([]int64, 0)
			minCostBF, bestPathBF := graph.TSPBruteForce(g, startVertex, &times)
			log.Println("Ścieżka:", bestPathBF)
			fmt.Printf("Brute force: minimalny koszt: %d, najlepsza ścieżka: %v\n", minCostBF, g.PathWithWeightsToString(bestPathBF))
			fmt.Printf("Czas: %s\n", timeTrack.FormatDurationFromNanoseconds(times[0]))
		case 8:
			fmt.Print("Podaj nazwę pliku: ")
			var fileName string
			_, err := fmt.Scanln(&fileName)
			if err != nil {
				return
			}
			err = graph.SaveGraphToFile(g, fileName)
			if err != nil {
				fmt.Println(err)
			}
		case 9:
			fmt.Println("Podaj wartość braku krawędzi")
			var noEdgeValue int
			_, err := fmt.Scanln(&noEdgeValue)
			if err != nil {
				return
			}
			g.SetNoEdgeValue(noEdgeValue)
		case 0:
			return
		default:
			fmt.Println("Niepoprawna opcja")
		}
	}
}
