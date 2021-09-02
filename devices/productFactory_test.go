package devices

import (
	"strings"
	"testing"
)

func TestGetDeviceFactoryCellPhone(t *testing.T) {
	brand := "iPhone"
	stock := 4

	device, err := GetDeviceFactory(CellPhoneType, brand, stock)
	if err != nil {
		t.Fatalf("A device type '%v' must exist", CellPhoneType)
	}

	if !strings.Contains(device.GetBrand(), brand) {
		t.Error("The device with brand wasn't correct")
	}

	if device.GetStock() != stock {
		t.Errorf("The device with brand %s does not have a correct stock, expected %d and has %d", device.GetBrand(), stock, device.GetStock())
	}

	t.Log("LOG:", device)
}
func TestGetDeviceFactoryCellPhoneChangeBrand(t *testing.T) {
	brand := "iPhone"
	stock := 4

	device, err := GetDeviceFactory(CellPhoneType, brand, stock)
	if err != nil {
		t.Fatalf("A device type '%v' must exist", CellPhoneType)
	}

	device.SetBrand("Motorola")

	if device.GetBrand() == brand {
		t.Error("The device with brand wasn't correct")
	}

	if device.GetStock() != stock {
		t.Errorf("The device with brand %s does not have a correct stock, expected %d and has %d", device.GetBrand(), stock, device.GetStock())
	}

	t.Log("LOG:", device)
}

func TestGetDeviceFactoryTablet(t *testing.T) {
	brand := "Sony"
	stock := 4
	device, err := GetDeviceFactory(TabletType, brand, stock)
	if err != nil {
		t.Fatalf("A device type '%v' must exist", TabletType)
	}

	if !strings.Contains(device.GetBrand(), brand) {
		t.Errorf("The device with brand %s wasn't correct", device.GetBrand())
	}

	if device.GetStock() != stock {
		t.Errorf("The device with brand %s does not have a correct stock, expected %d and has %d", device.GetBrand(), stock, device.GetStock())
	}

	t.Log("LOG:", device)
}

func TestGetDeviceFactoryTabletChangeStock(t *testing.T) {
	brand := "Sony"
	stock := 4
	device, err := GetDeviceFactory(TabletType, brand, stock)
	if err != nil {
		t.Fatalf("A device type '%v' must exist", TabletType)
	}

	device.SetStock(13)

	if !strings.Contains(device.GetBrand(), brand) {
		t.Errorf("The device with brand %s wasn't correct", device.GetBrand())
	}

	if device.GetStock() == stock {
		t.Errorf("The device with brand %s does not have a correct stock, expected %d and has %d", device.GetBrand(), stock, device.GetStock())
	}

	t.Log("LOG:", device)
}

func TestGetDeviceFactorySmartWatch(t *testing.T) {
	brand := "Huawey"
	stock := 6
	device, err := GetDeviceFactory(SmartWatchType, brand, stock)
	if err != nil {
		t.Fatalf("A device type '%v' must exist", SmartWatchType)
	}

	if !strings.Contains(device.GetBrand(), brand) {
		t.Errorf("The device with brand %s wasn't correct", device.GetBrand())
	}

	if device.GetStock() != stock {
		t.Errorf("The device with brand %s does not have a correct stock, expected %d and has %d", device.GetBrand(), stock, device.GetStock())
	}

	t.Log("LOG:", device)
}

func TestGetDeviceFactoryNonExistent(t *testing.T) {
	_, err := GetDeviceFactory(NintendoSwitchType, "Ninteno", 2)
	if err == nil {
		t.Errorf("device type %T not recognized", NintendoSwitchType)
	}
	t.Log("LOG:", err)
}
