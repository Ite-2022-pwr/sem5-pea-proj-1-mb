package control

import (
	"log"
	"os"
	"projekt1/graph"
	"runtime"
	"strconv"
)

// 0 - brute force, 1 - dynamic programming, 2 - branch and bound
func RunSingleTest(roundsForSize, maxSize, minSize, chosenTest int, fileName string) {
	timesMatrix := make([][]int64, maxSize-minSize+1)
	for i := 0; i < maxSize-minSize+1; i++ {
		timesMatrix[i] = make([]int64, roundsForSize)
	}

	for size := minSize; size <= maxSize; size++ {
		for round := 0; round < roundsForSize; round++ {
			log.Println("Size:", size, "Round:", round)
			g := new(graph.AdjMatrixGraph)
			graph.GenerateRandomGraph(g, size, -1, 100)
			startVertex := 0
			times := make([]int64, 0)
			var path []int
			switch chosenTest {
			case 0:
				_, path = graph.TSPBruteForce(g, startVertex, &times)
			case 1:
				_, path = graph.TSPDynamicProgramming(g, startVertex, &times)
			case 2:
				_, path = graph.TSPBranchAndBound(g, startVertex, &times)
			default:
				return
			}
			log.Println(g.PathWithWeightsToString(path))
			timesMatrix[size-minSize][round] = times[0]
			runtime.GC()
		}
	}
	saveTimesToCSVFile(timesMatrix, fileName)
}

func BNBTests(roundsForSize, maxSize, minSize int, fileName string) {
	timesMatrix := make([][]int64, maxSize-minSize+1)
	for i := 0; i < maxSize-minSize+1; i++ {
		timesMatrix[i] = make([]int64, roundsForSize)
	}

	for size := minSize; size <= maxSize; size++ {
		for round := 0; round < roundsForSize; round++ {
			g := new(graph.AdjMatrixGraph)
			graph.GenerateRandomGraph(g, size, -1, 100)
			startVertex := 0
			times := make([]int64, 0)
			_, _ = graph.TSPBranchAndBound(g, startVertex, &times)
			timesMatrix[size-minSize][round] = times[0]
		}
	}
	saveTimesToCSVFile(timesMatrix, fileName)
}

func DPTests(roundsForSize, maxSize, minSize int, fileName string) {
	timesMatrix := make([][]int64, maxSize-minSize+1)
	for i := 0; i < maxSize-minSize+1; i++ {
		timesMatrix[i] = make([]int64, roundsForSize)
	}

	for size := minSize; size <= maxSize; size++ {
		for round := 0; round < roundsForSize; round++ {
			g := new(graph.AdjMatrixGraph)
			graph.GenerateRandomGraph(g, size, -1, 100)
			startVertex := 0
			times := make([]int64, 0)
			_, _ = graph.TSPDynamicProgramming(g, startVertex, &times)
			timesMatrix[size-minSize][round] = times[0]
		}
	}
	saveTimesToCSVFile(timesMatrix, fileName)
}

func BFTests(roundsForSize, maxSize, minSize int, fileName string) {
	timesMatrix := make([][]int64, maxSize-minSize+1)
	for i := 0; i < maxSize-minSize+1; i++ {
		timesMatrix[i] = make([]int64, roundsForSize)
	}

	for size := minSize; size <= maxSize; size++ {
		for round := 0; round < roundsForSize; round++ {
			g := new(graph.AdjMatrixGraph)
			graph.GenerateRandomGraph(g, size, -1, 100)
			startVertex := 0
			times := make([]int64, 0)
			_, _ = graph.TSPBruteForce(g, startVertex, &times)
			timesMatrix[size-minSize][round] = times[0]
		}
	}
	saveTimesToCSVFile(timesMatrix, fileName)
}

func saveTimesToCSVFile(timesMatrix [][]int64, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	transposedMatrix := transposeTimesMatrix(timesMatrix)

	for i := 0; i < len(transposedMatrix); i++ {
		for j := 0; j < len(transposedMatrix[i]); j++ {
			_, _ = file.WriteString(strconv.FormatInt(transposedMatrix[i][j], 10) + ";")
		}
		_, _ = file.WriteString("\n")
	}

	_, _ = file.WriteString("\n")

}

func transposeTimesMatrix(timesMatrix [][]int64) [][]int64 {
	transposedMatrix := make([][]int64, len(timesMatrix[0]))
	for i := 0; i < len(timesMatrix[0]); i++ {
		transposedMatrix[i] = make([]int64, len(timesMatrix))
	}

	for i := 0; i < len(timesMatrix); i++ {
		for j := 0; j < len(timesMatrix[i]); j++ {
			transposedMatrix[j][i] = timesMatrix[i][j]
		}
	}

	return transposedMatrix
}
