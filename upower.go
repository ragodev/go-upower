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
