package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Example for a copper pickaxe mining cobblestone
	//set hotabar spaces that have copper pickaxe (inclusive)
	const min = 3         //where the picaxe line starts
	const max = 6         //where it should end
	const sleepTime = 370 //how long it takes to break a picaxe or until it switches to the next
	//If mining cobblestone with a cobblestone generator, remember to account for the time it takes to replace the cobblestone
	fmt.Println("Afker v1.0")
	fmt.Println("Move your mouse to your posistion")
	fmt.Println("Starting in: 10 seconds")
	time.Sleep(10 * time.Second)

	currentSpace := min

	robotgo.MouseDown("left")
	defer robotgo.MouseUp("left")
	fmt.Println("Mouse is down")
	defer fmt.Println("Mouse is up")

	for currentSpace <= max {
		fmt.Println("Waiting for picaxe to break")
		time.Sleep(sleepTime * time.Second) //~190 cobblestone (give or take for lag and such)
		fmt.Println("Mousing up")
		robotgo.MouseUp("left")
		toMove := strconv.Itoa(currentSpace + 1)
		fmt.Println("moving to space:", toMove)
		robotgo.KeyTap(toMove)
		robotgo.MouseDown("left")
		fmt.Println("Mouse is down")
		currentSpace++
	}

}
