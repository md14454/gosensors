package gosensors

// #cgo LDFLAGS: -lsensors
// #include <stdlib.h>
// #include <stdio.h>
// #include <sensors/sensors.h>
import "C"

import (
	"log"
	"unsafe"
)

type Bus struct {
	Type int16
	Nr   int16
}

type Chip struct {
	Prefix string
	Bus    Bus
	Addr   int32
	Path   string
}

func Init() {
	filename := C.CString("/etc/sensors3.conf")
	defer C.free(unsafe.Pointer(filename))
	mode := C.CString("r")
	defer C.free(unsafe.Pointer(mode))

	fp, err := C.fopen(filename, mode)
	defer C.fclose(fp)

	if fp == nil {
		log.Fatal(err)
	}

	C.sensors_init(fp)
}

func Cleanup() {
	C.sensors_cleanup()
}

func GetDetectedChips() []Chip {
	var chips []Chip

	var count C.int = 0

	for {
		resp := C.sensors_get_detected_chips(nil, &count)

		if resp == nil {
			break
		}

		var bus Bus
		bus.Type = int16(resp.bus._type)
		bus.Nr = int16(resp.bus.nr)

		var chip Chip
		chip.Prefix = C.GoString(resp.prefix)
		chip.Bus = bus
		chip.Addr = int32(resp.addr)
		chip.Path = C.GoString(resp.path)

		chips = append(chips, chip)

	}

	return chips
}
