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

func FormatDurationFromNanoseconds(nanos int64) string {
	// Konwersja nanosekund na time.Duration
	d := time.Duration(nanos)

	// Obliczanie poszczególnych jednostek czasu
	hours := int64(d.Hours())
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60
	milliseconds := int64(d.Milliseconds()) % 1000
	microseconds := int64(d.Microseconds()) % 1000
	nanoseconds := nanos % 1000

	// Formatowanie czasu na czytelny format
	return fmt.Sprintf("%02d:%02d:%02d.%03dms %03dµs %03dns", hours, minutes, seconds, milliseconds, microseconds, nanoseconds)
}
