package afker

import (
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

	for currentSpace <= max {
		toMove := "num" + strconv.Itoa(currentSpace+1)
		robotgo.KeyTap(toMove)
		currentSpace++
	}

	time.Sleep(144000) //~190 cobblestone (give or take for lag and such)

}
