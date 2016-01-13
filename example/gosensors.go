package main

import (
	"fmt"

	"github.com/md14454/gosensors"
)

func main() {
	gosensors.Init()

	chips := gosensors.GetDetectedChips()
	fmt.Printf("Found %d chips\n", len(chips))

	for i := 0; i < len(chips); i++ {
		fmt.Printf("Found <Chip prefix=%s bus=%d addr=%d path=%s>\n", chips[i].Prefix, chips[i].Bus, chips[i].Addr, chips[i].Path)
	}

	gosensors.Cleanup()
}
