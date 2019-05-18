package dev

import (
	"github.com/Hodapp87/ble"
	"github.com/Hodapp87/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
