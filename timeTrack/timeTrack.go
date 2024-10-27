package timeTrack

import (
	"log"
	"time"
)

func TimeTrack(start time.Time, name string) int64 {
	elapsed := time.Since(start)
	log.Printf("%s zajęło %s", name, elapsed)
	return elapsed.Nanoseconds()
}
