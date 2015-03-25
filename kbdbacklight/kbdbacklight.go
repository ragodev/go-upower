package kbdbacklight

import "github.com/godbus/dbus"

func GetMaxBrightness() (b int64, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.GetMaxBrightness", 0)
	if call.Err != nil {

		return 0, call.Err
	}

	if err := call.Store(&b); err != nil {

		return 0, err
	}

	return
}

func GetBrightness() (b int64, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.GetBrightness", 0)
	if call.Err != nil {

		return 0, call.Err
	}

	if err := call.Store(&b); err != nil {

		return 0, err
	}

	return
}

func SetBrightness(b int64) (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.SetBrightness", 0, b)
	if call.Err != nil {

		return 0, call.Err
	}

	return
}
