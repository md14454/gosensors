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
	bus  *C.struct_sensors_bus_id
}

func (b Bus) String() string {
	if b.Type == -1 {
		return "*"
	} else {
		return C.GoString(C.sensors_get_adapter_name(b.bus))
	}
}

type Chip struct {
	Prefix string
	Bus    Bus
	Addr   int32
	Path   string
	chip   *C.struct_sensors_chip_name
}

func (c Chip) String() string {
	var buffer [200]C.char

	len := C.sensors_snprintf_chip_name(&buffer[0], C.size_t(len(buffer)), c.chip)

	return C.GoStringN(&buffer[0], len)
}

func (c Chip) AdapterName() string {
	return c.Bus.String()
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

		bus := Bus{
			Type: int16(resp.bus._type),
			Nr:   int16(resp.bus.nr),
			bus:  &resp.bus,
		}

		chip := Chip{
			Prefix: C.GoString(resp.prefix),
			Bus:    bus,
			Addr:   int32(resp.addr),
			Path:   C.GoString(resp.path),
			chip:   resp,
		}

		chips = append(chips, chip)

	}

	return chips
}
