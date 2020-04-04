package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	hours := flag.Int("hours", 1, "hours")
	mins := flag.Int("mins", -1, "mins")
	freq := flag.Int("freq", 60, "move the mouse every N seconds")

	flag.Usage = func() {
		fmt.Println("Usage of jitter:")
		fmt.Println("  -freq int")
		fmt.Println("  	move the mouse every N seconds (default 60)")
		fmt.Println("  -hours int")
		fmt.Println("  	hours (default 1)")
		fmt.Println("  -mins int")
		fmt.Println("  	mins (default -1)")
	}

	flag.Parse()

	// sanity check the args
	if *mins < 0 {
		*mins = 0
	}

	if *hours < 0 {
		*hours = 0
	}

	durationStr := fmt.Sprintf("%dh%dm", *hours, *mins)
	jitterDuration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Errorf("Invalid input - bad duration string '%s'\n", durationStr)
	}

	fmt.Printf("Jittering the mouse every %d seconds for %s...\n", *freq, jitterDuration)

	ticker := time.NewTicker(time.Duration(*freq) * time.Second)
	defer ticker.Stop()

	stopJittering := make(chan bool)
	go func() {
		time.Sleep(jitterDuration)
		stopJittering <- true
	}()

	var origX, origY, newX, newY, delta int
	delta = 5 // move 5 pixels diagonally by default
	for {
		select {
		case <-stopJittering:
			fmt.Println("Time's up! Stopping mouse jitter.")
			return
		case <-ticker.C:
			origX, origY = robotgo.GetMousePos()

			// Try to move diagonally down first.
			robotgo.MoveSmooth(origX+delta, origY+delta)

			newX, newY = robotgo.GetMousePos()
			if newX == origX && newY == origY {
				// Oops! We are probably in a corner and need to go in a different direction
				robotgo.MoveSmooth(origX-delta, origY-delta)
			}

			// Go back to our original starting position.
			robotgo.MoveSmooth(origX, origY)
		}
	}
}
