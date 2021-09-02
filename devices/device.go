package devices

type Device struct {
	brand string
	stock int
}

func (d *Device) GetBrand() string {
	return d.brand
}

func (d *Device) SetBrand(brand string) {
	d.brand = brand
}

func (d *Device) GetStock() int {
	return d.stock
}

func (d *Device) SetStock(stock int) {
	d.stock = stock
}
