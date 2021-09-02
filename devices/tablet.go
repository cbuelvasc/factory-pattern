package devices

import (
	"fmt"
)

type Tablet struct {
	Device
}

func (d *Tablet) String() string {
	return fmt.Sprintf("Tablet with brand: %s, and stock %d\n", d.GetBrand(), d.GetStock())
}

func NewTablet(brand string, stock int) IDevice {
	return &Tablet{
		Device: Device{
			brand: brand,
			stock: stock,
		},
	}
}
