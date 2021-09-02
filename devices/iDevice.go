package devices

type IDevice interface {
	GetBrand() string
	SetBrand(brand string)
	GetStock() int
	SetStock(stock int)
	String() string
}
