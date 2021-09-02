package devices

import "fmt"

func GetDeviceFactory(deviceType DeviceType, name string, stock int) (IDevice, error) {
	switch deviceType {
	case CellPhoneType:
		return NewCellPhone(name, stock), nil
	case TabletType:
		return NewTablet(name, stock), nil
	case SmartWatchType:
		return NewSmartWatch(name, stock), nil
	default:
		return nil, fmt.Errorf("device type %v not recognized", deviceType)
	}
}
