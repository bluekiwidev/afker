package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Example for a copper pickaxe mining cobblestone
	// Set hotbar spaces that have copper pickaxes (inclusive)
	const min = 1         // Where the pickaxe line starts (hotbar slot 1-9)
	const max = 9         // Where it should end (hotbar slot 1-9)
	const sleepTime = 144 // How long it takes to break a pickaxe or until it switches to the next
	// If mining cobblestone with a cobblestone generator, remember to account for the time it takes to replace the cobblestone
	fmt.Println("Afker v1.0")
	fmt.Println("Move your mouse to your position")
	fmt.Println("Starting in: 10 seconds")
	time.Sleep(10 * time.Second)

	currentSpace := min

	// Ensure we start with the first hotbar slot selected
	fmt.Println("Selecting hotbar slot:", currentSpace)
	robotgo.KeyTap(strconv.Itoa(currentSpace))
	time.Sleep(100 * time.Millisecond) // Brief delay to ensure key press registers

	robotgo.MouseDown("left")
	defer robotgo.MouseUp("left")
	fmt.Println("Mouse is down - mining with hotbar slot:", currentSpace)
	defer fmt.Println("Mouse is up")

	// Process all pickaxes from min to max
	for currentSpace <= max {
		fmt.Println("Waiting for pickaxe to break...")
		time.Sleep(sleepTime * time.Second) // ~190 cobblestone (give or take for lag and such)
		currentSpace++
		// Switch to the next hotbar slot if there's one available
		if currentSpace <= max {
			toMove := strconv.Itoa(currentSpace)
			fmt.Println("Switching to hotbar slot:", toMove)
			robotgo.KeyTap(toMove)
		}
	}

}
