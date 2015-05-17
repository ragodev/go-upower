package qos

import "github.com/godbus/dbus"

// Latency Types
const (
	CpuDma  = "cpu_dma"
	Network = "network"
)

// Latency Values
const (
	// The value -1 means unset and the default is used.
	LatencyDefault = -1
)

func SetMinimumLatency(t string, value int) (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.SetMinimumLatency", 0, t, value)
	if call.Err != nil {

		return call.Err
	}

	return
}

func RequestLatency(t string, value int, persistent bool) (cookie uint32, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.RequestLatency", 0, t, value, persistent)
	if call.Err != nil {

		return call.Err
	}

	if err := call.Store(&cookie); err != nil {

		return nil, err
	}

	return
}

func CancelRequest(t string, cookie uint32) (err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.CancelRequest", 0, t, cookie)
	if call.Err != nil {

		return call.Err
	}

	return
}

func GetLatency(t string) (value int, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.GetLatency", 0, t)
	if call.Err != nil {

		return call.Err
	}

	if err := call.Store(&value); err != nil {

		return nil, err
	}

	return
}

func GetLatencyRequests() (requests map[string]dbus.Variant, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}
	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")

	call := obj.Call("org.freedesktop.UPower.GetLatencyRequests", 0)
	if call.Err != nil {

		return call.Err
	}

	if err := call.Store(&requests); err != nil {

		return nil, err
	}

	return
}

func SignalLatencyChanged() (ch chan *dbus.Signal, call *dbus.Call, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=LatencyChanged")
	if call.Err != nil {

		return nil, nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}

func SignalRequestsChanged() (ch chan *dbus.Signal, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {

		return
	}

	call = conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "sender=org.freedesktop.UPower,type=signal,member=RequestsChanged")
	if call.Err != nil {

		return nil, nil, call.Err
	}

	ch = make(chan *dbus.Signal, 10)
	conn.Signal(ch)

	return
}
