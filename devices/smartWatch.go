package devices

import (
	"fmt"
)

type SmartWatch struct {
	Device
}

func (d *SmartWatch) String() string {
	return fmt.Sprintf("SmartWatch with brand: %s, and stock %d\n", d.GetBrand(), d.GetStock())
}

func NewSmartWatch(brand string, stock int) IDevice {
	return &SmartWatch{
		Device: Device{
			brand: brand,
			stock: stock,
		},
	}
}
