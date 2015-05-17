package kbdbacklight

import "github.com/godbus/dbus"

// Get the maximum brightness level for the keyboard backlight.
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

// Get the brightness level of the keyboard backlight.
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

// Set the brightness level of the keyboard backlight.
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

// The keyboard backlight brightness level has changed.
func SignalBrightnessChanged() (ch chan *dbus.Signal, call *dbus.Call, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower.KbdBacklight,type=signal,member=BrightnessChanged")
	if call.Err != nil {

		return nil, nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}
