package main

import (
	"fmt"

	"github.com/md14454/gosensors"
)

func main() {
	gosensors.Init()
	defer gosensors.Cleanup()

	chips := gosensors.GetDetectedChips()

	for i := 0; i < len(chips); i++ {
		fmt.Println(chips[i])
		fmt.Println("")
	}
}
