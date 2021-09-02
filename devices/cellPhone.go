package devices

import (
	"fmt"
)

type CellPhone struct {
	Device
}

func (d *CellPhone) String() string {
	return fmt.Sprintf("CellPhone with brand: %s, and stock %d\n", d.GetBrand(), d.GetStock())
}

func NewCellPhone(brand string, stock int) IDevice {
	return &CellPhone{
		Device: Device{
			brand: brand,
			stock: stock,
		},
	}
}
