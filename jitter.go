package main

import (
	"flag"
	"fmt"
	"math"
	"time"

	"github.com/go-vgo/robotgo"
)

func circle(x, y, radius, angleStep int) {
	var angle, newX, newY int
	for angle < 360 {
		// calculate new position
		newX = x + int(float64(radius)*math.Cos(float64(-angle)*math.Pi/180.0))
		newY = y - int(float64(radius)*math.Sin(float64(-angle)*math.Pi/180.0))

		// ...then move to it
		robotgo.MoveSmooth(newX, newY)

		angle = angle + angleStep
	}
}

func diagonal(origX, origY, delta int) {
	// Try to move diagonally down first.
	robotgo.MoveSmooth(origX+delta, origY+delta)

	newX, newY := robotgo.GetMousePos()
	if newX == origX && newY == origY {
		// Oops! We are probably in a corner and need to go in a different direction.
		robotgo.MoveSmooth(origX-delta, origY-delta)
	}

}

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

	var origX, origY, delta int
	delta = 10 // move 10 pixels diagonally by default
	for {
		select {
		case <-stopJittering:
			fmt.Println("Time's up! Stopping mouse jitter.")
			return
		case <-ticker.C:
			origX, origY = robotgo.GetMousePos()
			if origX > delta && origY > delta {
				// We're far enough from the top left corner to do a circle.
				circle(origX, origY, delta, 45)
			} else {
				// We're close to a corner - move diagonally only.
				diagonal(origX, origY, delta)
			}

			// Go back to our original starting position.
			robotgo.MoveSmooth(origX, origY)
		}
	}
}
