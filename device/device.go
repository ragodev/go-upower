package device

import (
	"time"

	"github.com/godbus/dbus"
)

// Value shared by multiple properties
const Unknown = 0

// Possible power sources
const (
	LinePower = 1
	Battery   = 2
	UPS       = 3
	Monitor   = 4
	Mouse     = 5
	Keyboard  = 6
	PDA       = 7
	Phone     = 8
)

// Possible states
const (
	Charging         = 1
	Discharging      = 2
	Empty            = 3
	FullyCharged     = 4
	PendingCharge    = 5
	PendingDischarge = 6
)

// Possible technologys
const (
	LithiumIon           = 1
	LithiumPolymer       = 2
	LithiumIronPhosphate = 3
	LeadAcid             = 4
	NickelCadmium        = 5
	NickelMetalHydride   = 6
)

// Properties
const (
	NativePath       = "NativePath"
	Vendor           = "Vendor"
	Model            = "Model"
	Serial           = "Serial"
	UpdateTime       = "UpdateTime"
	Type             = "Type"
	PowerSupply      = "PowerSupply"
	HasHistory       = "HasHistory"
	HasStatistics    = "HasStatistics"
	Online           = "Online"
	Energy           = "Energy"
	EnergyEmpty      = "EnergyEmpty"
	EnergyFull       = "EnergyFull"
	EnergyFullDesign = "EnergyFullDesign"
	EnergyRate       = "EnergyRate"
	Voltage          = "Voltage"
	TimeToEmpty      = "TimeToEmpty"
	TimeToFull       = "TimeToFull"
	Percentage       = "Percentage"
	IsPresent        = "IsPresent"
	State            = "State"
	IsRechargeable   = "IsRechargeable"
	Capacity         = "Capacity"
	Technology       = "Technology"
	RecallNotice     = "RecallNotice"
	RecallVendor     = "RecallVendor"
	RecallUrl        = "RecallUrl"
)

func Refresh() (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.Refresh", 0)
	if call.Err != nil {

		return call.Err
	}

	return
}

/*
func GetHistory() (mode string, timespan uint32, resolution uint32, data, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.GetHistory", 0)
	if call.Err != nil {

		return nil, call.Err
	}

	if err := call.Store(&mode, &timespan, &resolution, &data); err != nil {

		return nil, err
	}

	return
}

func GetStatistics() (mode string, data, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.GetStatistics", 0)
	if call.Err != nil {

		return nil, call.Err
	}

	if err := call.Store(&mode & data); err != nil {

		return nil, err
	}

	return
}*/

type Properties struct {
	NativePath       string
	Vendor           string
	Model            string
	Serial           string
	UpdateTime       uint64
	Type             uint32
	PowerSupply      bool
	HasHistory       bool
	HasStatistics    bool
	Online           bool
	Energy           float64
	EnergyEmpty      float64
	EnergyFull       float64
	EnergyFullDesign float64
	EnergyRate       float64
	Voltage          float64
	TimeToEmpty      time.Duration
	TimeToFull       time.Duration
	Percentage       float64
	IsPresent        bool
	State            uint32
	IsRechargeable   bool
	Capacity         float64
	Technology       uint32
	RecallNotice     bool
	RecallVendor     string
	RecallUrl        string
}

func GetAllProperties(dev dbus.ObjectPath) (p *Properties, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", dev)

	call := obj.Call("org.freedesktop.DBus.Properties.GetAll", 0, "org.freedesktop.UPower")
	if call.Err != nil {

		return nil, call.Err
	}

	p = &Properties{}
	m := map[string]dbus.Variant{}

	if err := call.Store(&m); err != nil {

		return nil, err
	}

	p.NativePath = m[NativePath].Value().(string)
	p.Vendor = m[Vendor].Value().(string)
	p.Model = m[Model].Value().(string)
	p.Serial = m[Serial].Value().(string)
	p.UpdateTime = m[UpdateTime].Value().(uint64)
	p.Type = m[Type].Value().(uint32)
	p.PowerSupply = m[PowerSupply].Value().(bool)
	p.HasHistory = m[HasHistory].Value().(bool)
	p.HasStatistics = m[HasStatistics].Value().(bool)
	p.Online = m[Online].Value().(bool)
	p.Energy = m[Energy].Value().(float64)
	p.EnergyEmpty = m[EnergyEmpty].Value().(float64)
	p.EnergyFull = m[EnergyFull].Value().(float64)
	p.EnergyFullDesign = m[EnergyFullDesign].Value().(float64)
	p.EnergyRate = m[EnergyRate].Value().(float64)
	p.Voltage = m[Voltage].Value().(float64)
	p.TimeToEmpty = time.Duration(time.Duration(m[TimeToEmpty].Value().(int64)) * time.Second)
	p.TimeToFull = time.Duration(time.Duration(m[TimeToFull].Value().(int64)) * time.Second)
	p.Percentage = m[Percentage].Value().(float64)
	p.IsPresent = m[IsPresent].Value().(bool)
	p.State = m[State].Value().(uint32)
	p.IsRechargeable = m[IsRechargeable].Value().(bool)
	p.Capacity = m[Capacity].Value().(float64)
	p.Technology = m[Technology].Value().(uint32)
	p.RecallNotice = m[RecallNotice].Value().(bool)
	p.RecallVendor = m[RecallVendor].Value().(string)
	p.RecallUrl = m[RecallUrl].Value().(string)

	return
}

// Get UPower device property
func GetProperty(dev dbus.ObjectPath, p string) (v dbus.Variant, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", dev)

	v, err = obj.GetProperty(p)
	if err != nil {

		return
	}

	return
}

// Some value on the power source changed.
func SignalChanged() (ch chan *dbus.Signal, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=Changed")
	if call.Err != nil {

		return nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}
