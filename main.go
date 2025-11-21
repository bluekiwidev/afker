package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	//set hotabar spaces that have copper pickaxe (inclusive)
	const min = 1
	const max = 9

	currentSpace := min

	robotgo.MouseDown("left")
	defer robotgo.MouseUp("left")
	fmt.Println("Mouse is down")
	defer fmt.Println("Mouse is up")

	for currentSpace <= max {
		fmt.Println("Waiting for picaxe to break")
		time.Sleep(144 * time.Second) //~190 cobblestone (give or take for lag and such)
		toMove := "num" + strconv.Itoa(currentSpace+1)
		fmt.Println("moving to space: ", toMove)
		robotgo.KeyTap(toMove)
		currentSpace++
	}

}
