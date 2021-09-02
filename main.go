package main

import (
	"fmt"

	factory "github.com/cbuelvasc/factory-pattern/devices"
)

func main() {
	cellPhone, _ := factory.GetDeviceFactory(factory.CellPhoneType, "Samsung", 23)
	tablet, _ := factory.GetDeviceFactory(factory.TabletType, "Sony", 12)
	smartWatch, _ := factory.GetDeviceFactory(factory.SmartWatchType, "Garmin", 2)

	fmt.Println(cellPhone)
	fmt.Println(tablet)
	fmt.Println(smartWatch)
}
