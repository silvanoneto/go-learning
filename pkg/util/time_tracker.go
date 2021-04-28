// Package util time_tracker has facilities to track the execution time
package util

import (
	"fmt"
	"time"
)

// TimeTrack prints a measure of time based on a start time
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
