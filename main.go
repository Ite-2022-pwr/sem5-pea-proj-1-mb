package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"projekt1/control"
	"time"
)

// Przykład użycia
func main() {

	// wyłączenie logów
	//log.SetOutput(io.Discard)

	//verticesAmountPTR := flag.Int("vertices", 10, "Number of vertices")
	//doBrunteForcePTR := flag.Bool("brute-force", false, "Do brute force")

	logToFilePTR := flag.Bool("log-to-file", false, "Log to file")
	runMenuPTR := flag.Bool("control", true, "Run control, if false run default tests")
	flag.Parse()
	//
	if *logToFilePTR {
		//save all logs to file
		dateString := time.Now().Format("2006-01-02_15:04:05")
		logFileName := dateString + ".log"
		f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println("Error opening file:", err)
		} else {
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					fmt.Println("Error closing file:", err)
				}
			}(f)
			multi := io.MultiWriter(f, os.Stdout)
			log.SetOutput(multi)
		}
	}

	if *runMenuPTR {
		control.Menu()
	} else {
		control.RunSingleTest(100, 13, 6, 1, "DP.csv")
		control.RunSingleTest(100, 13, 6, 2, "BNB.csv")
		control.RunSingleTest(100, 13, 6, 0, "BF.csv")
	}
}
