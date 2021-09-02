package devices

type DeviceType int

const (
	CellPhoneType DeviceType = iota + 1
	TabletType
	SmartWatchType
	NintendoSwitchType
)
