package abrecovery

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type ABRecovery struct {
	abRecovery // interface com.deepin.ABRecovery
	proxy.Object
}

func NewABRecovery(conn *dbus.Conn) *ABRecovery {
	obj := new(ABRecovery)
	obj.Object.Init_(conn, "com.deepin.ABRecovery", "/com/deepin/ABRecovery")
	return obj
}

type abRecovery struct{}

func (v *abRecovery) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*abRecovery) GetInterfaceName_() string {
	return "com.deepin.ABRecovery"
}

// method CanBackup

func (v *abRecovery) GoCanBackup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanBackup", flags, ch)
}

func (*abRecovery) StoreCanBackup(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *abRecovery) CanBackup(flags dbus.Flags) (can bool, err error) {
	return v.StoreCanBackup(
		<-v.GoCanBackup(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanRestore

func (v *abRecovery) GoCanRestore(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanRestore", flags, ch)
}

func (*abRecovery) StoreCanRestore(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *abRecovery) CanRestore(flags dbus.Flags) (can bool, err error) {
	return v.StoreCanRestore(
		<-v.GoCanRestore(flags, make(chan *dbus.Call, 1)).Done)
}

// method StartBackup

func (v *abRecovery) GoStartBackup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartBackup", flags, ch)
}

func (v *abRecovery) StartBackup(flags dbus.Flags) error {
	return (<-v.GoStartBackup(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StartRestore

func (v *abRecovery) GoStartRestore(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartRestore", flags, ch)
}

func (v *abRecovery) StartRestore(flags dbus.Flags) error {
	return (<-v.GoStartRestore(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal JobEnd

func (v *abRecovery) ConnectJobEnd(cb func(kind string, success bool, errMsg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "JobEnd", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".JobEnd",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var kind string
		var success bool
		var errMsg string
		err := dbus.Store(sig.Body, &kind, &success, &errMsg)
		if err == nil {
			cb(kind, success, errMsg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property BackingUp b

func (v *abRecovery) BackingUp() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "BackingUp",
	}
}

// property Restoring b

func (v *abRecovery) Restoring() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Restoring",
	}
}

// property ConfigValid b

func (v *abRecovery) ConfigValid() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ConfigValid",
	}
}
