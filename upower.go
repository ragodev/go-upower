// Package upower provides an implementation of the Freedesktop UPower Specification
// using the DBus API.
package upower

import "github.com/godbus/dbus"

// Properties
const (
	DaemonVersion = "upower.DaemonVersion"
	CanSuspend    = "upower.CanSuspend"
	CanHibernate  = "upower.CanHibernate"
	OnBattery     = "upower.OnBattery"
	OnLowBattery  = "upower.OnLowBattery"
	LidIsClosed   = "upower.LidIsClosed"
	LidIsPresent  = "upower.LidIsPresent"
)

// Enumerate all power objects on the system.
func EnumerateDevices() (devices []dbus.ObjectPath, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.EnumerateDevices", 0)
	if call.Err != nil {

		return nil, call.Err
	}

	if err := call.Store(&devices); err != nil {

		return nil, err
	}

	return
}

// This method tells UPower that the Suspend() or Hibernate() method is about to be called.
func AboutToSleep() (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.AboutToSleep", 0)
	if call.Err != nil {

		return call.Err
	}

	return
}

// Suspends the computer into a low power state.
func Suspend() (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.Suspend", 0)
	if call.Err != nil {

		return call.Err
	}

	return
}

// Check if the caller has (or can get) the PolicyKit privilege to call Suspend.
func SuspendAllowed() (ok bool, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.SuspendAllowed", 0)
	if call.Err != nil {

		return false, call.Err
	}

	if err := call.Store(&ok); err != nil {

		return false, err
	}

	return
}

// Hibernates the computer into a low power state.
func Hibernate() (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.Hibernate", 0)
	if call.Err != nil {

		return call.Err
	}

	return
}

// Check if the caller has (or can get) the PolicyKit privilege to call Hibernate.
func HibernateAllowed() (ok bool, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.HibernateAllowed", 0)
	if call.Err != nil {

		return false, call.Err
	}

	if err := call.Store(&ok); err != nil {

		return false, err
	}

	return
}

// Get UPower property
func GetProperty(p string) (v dbus.Variant, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	v, err = obj.GetProperty(p)
	if err != nil {

		return
	}

	return
}

// Emitted when a device is added.
func SignalDeviceAdded() (ch chan *dbus.Signal, call *dbus.Call, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=DeviceAdded")
	if call.Err != nil {

		return nil, nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}

// Emitted when a device is removed.
func SignalDeviceRemoved() (ch chan *dbus.Signal, call *dbus.Call, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=DeviceRemoved")
	if call.Err != nil {

		return nil, nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}

// Emitted when a device changed.
func SignalDeviceChanged() (ch chan *dbus.Signal, call *dbus.Call, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=DeviceChanged")
	if call.Err != nil {

		return nil, nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}

// Emitted when one or more properties on the object changes.
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

// This signal is sent when the session is about to be suspended or hibernated.
func SignalSleeping() (ch chan *dbus.Signal, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=Sleeping")
	if call.Err != nil {

		return nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}

// This signal is sent when the session has just returned from Suspend() or Hibernate().
func SignalResuming() (ch chan *dbus.Signal, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=Resuming")
	if call.Err != nil {

		return nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}
