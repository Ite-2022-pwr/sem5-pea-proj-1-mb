package timeTrack

import (
	"fmt"
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) int64 {
	elapsed := time.Since(start)
	fmt.Printf("%s zajęło %s\n", name, elapsed)
	log.Printf("%s zajęło %s", name, elapsed)
	return elapsed.Nanoseconds()
}
