package main

import (
	"fmt"
	"math/rand"
)

type Partition string

const (
	PartitionA Partition = "A"
	PartitionB Partition = "B"
)

type Firmware struct {
	Version string
	Valid   bool
}

type Device struct {
	ActivePartition Partition
	PartitionA      Firmware
	PartitionB      Firmware
	LastKnownGood   Partition
}

func (d *Device) GetInactivePartition() Partition {
	if d.ActivePartition == PartitionA {
		return PartitionB
	}
	return PartitionA
}

func (d *Device) InstallFirmware(part Partition, fw Firmware) {
	switch part {
	case PartitionA:
		d.PartitionA = fw
	case PartitionB:
		d.PartitionB = fw
	}
	fmt.Printf("Firmware v%s installed to partition %s\n", fw.Version, part)
}

func (d *Device) Boot() {
	activeFirmware := d.getActiveFirmware()
	fmt.Printf("Booting partition %s with firmware v%s...\n", d.ActivePartition, activeFirmware.Version)

	if activeFirmware.Valid {
		fmt.Println("Boot successful.")
		d.LastKnownGood = d.ActivePartition
	} else {
		fmt.Println("Boot failed! Rolling back to last known good partition...")
		d.ActivePartition = d.LastKnownGood
	}
}

func (d *Device) getActiveFirmware() Firmware {
	if d.ActivePartition == PartitionA {
		return d.PartitionA
	}
	return d.PartitionB
}

func main() {
	device := Device{
		ActivePartition: PartitionA,
		PartitionA:      Firmware{Version: "1.0", Valid: true},
		LastKnownGood:   PartitionA,
	}

	// Simulate update
	newFirmware := Firmware{Version: "2.0", Valid: rand.Intn(2) == 1}
	inactive := device.GetInactivePartition()
	device.InstallFirmware(inactive, newFirmware)
	device.ActivePartition = inactive

	// Attempt to boot
	device.Boot()
}
