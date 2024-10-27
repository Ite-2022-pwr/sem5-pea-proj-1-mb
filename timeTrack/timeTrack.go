package timeTrack

import (
	"fmt"
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) int64 {
	elapsed := time.Since(start)
	log.Printf("%s zajęło %s", name, elapsed)
	fmt.Printf("%s zajęło %s\n", name, elapsed)
	return elapsed.Nanoseconds()
}
