package dev

import (
	"github.com/Hodapp87/ble"
	"github.com/Hodapp87/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
