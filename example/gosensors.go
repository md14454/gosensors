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
		chip := chips[i]

		fmt.Println(chip)
		fmt.Println("Adapter:", chip.AdapterName())
		fmt.Println("")
	}
}
