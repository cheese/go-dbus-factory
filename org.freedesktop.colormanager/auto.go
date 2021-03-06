package colormanager

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

type Manager struct {
	manager // interface org.freedesktop.ColorManager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.ColorManager", "/org/freedesktop/ColorManager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.ColorManager"
}

// method GetDevices

func (v *manager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*manager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDevicesByKind

func (v *manager) GoGetDevicesByKind(flags dbus.Flags, ch chan *dbus.Call, kind string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevicesByKind", flags, ch, kind)
}

func (*manager) StoreGetDevicesByKind(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetDevicesByKind(flags dbus.Flags, kind string) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevicesByKind(
		<-v.GoGetDevicesByKind(flags, make(chan *dbus.Call, 1), kind).Done)
}

// method FindDeviceById

func (v *manager) GoFindDeviceById(flags dbus.Flags, ch chan *dbus.Call, device_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindDeviceById", flags, ch, device_id)
}

func (*manager) StoreFindDeviceById(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) FindDeviceById(flags dbus.Flags, device_id string) (object_path dbus.ObjectPath, err error) {
	return v.StoreFindDeviceById(
		<-v.GoFindDeviceById(flags, make(chan *dbus.Call, 1), device_id).Done)
}

// method FindSensorById

func (v *manager) GoFindSensorById(flags dbus.Flags, ch chan *dbus.Call, sensor_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindSensorById", flags, ch, sensor_id)
}

func (*manager) StoreFindSensorById(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) FindSensorById(flags dbus.Flags, sensor_id string) (object_path dbus.ObjectPath, err error) {
	return v.StoreFindSensorById(
		<-v.GoFindSensorById(flags, make(chan *dbus.Call, 1), sensor_id).Done)
}

// method FindDeviceByProperty

func (v *manager) GoFindDeviceByProperty(flags dbus.Flags, ch chan *dbus.Call, key string, value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindDeviceByProperty", flags, ch, key, value)
}

func (*manager) StoreFindDeviceByProperty(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) FindDeviceByProperty(flags dbus.Flags, key string, value string) (object_path dbus.ObjectPath, err error) {
	return v.StoreFindDeviceByProperty(
		<-v.GoFindDeviceByProperty(flags, make(chan *dbus.Call, 1), key, value).Done)
}

// method FindProfileById

func (v *manager) GoFindProfileById(flags dbus.Flags, ch chan *dbus.Call, profile_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindProfileById", flags, ch, profile_id)
}

func (*manager) StoreFindProfileById(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) FindProfileById(flags dbus.Flags, profile_id string) (object_path dbus.ObjectPath, err error) {
	return v.StoreFindProfileById(
		<-v.GoFindProfileById(flags, make(chan *dbus.Call, 1), profile_id).Done)
}

// method FindProfileByProperty

func (v *manager) GoFindProfileByProperty(flags dbus.Flags, ch chan *dbus.Call, key string, value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindProfileByProperty", flags, ch, key, value)
}

func (*manager) StoreFindProfileByProperty(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) FindProfileByProperty(flags dbus.Flags, key string, value string) (object_path dbus.ObjectPath, err error) {
	return v.StoreFindProfileByProperty(
		<-v.GoFindProfileByProperty(flags, make(chan *dbus.Call, 1), key, value).Done)
}

// method FindProfileByFilename

func (v *manager) GoFindProfileByFilename(flags dbus.Flags, ch chan *dbus.Call, filename string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindProfileByFilename", flags, ch, filename)
}

func (*manager) StoreFindProfileByFilename(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) FindProfileByFilename(flags dbus.Flags, filename string) (object_path dbus.ObjectPath, err error) {
	return v.StoreFindProfileByFilename(
		<-v.GoFindProfileByFilename(flags, make(chan *dbus.Call, 1), filename).Done)
}

// method GetStandardSpace

func (v *manager) GoGetStandardSpace(flags dbus.Flags, ch chan *dbus.Call, standard_space string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetStandardSpace", flags, ch, standard_space)
}

func (*manager) StoreGetStandardSpace(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) GetStandardSpace(flags dbus.Flags, standard_space string) (object_path dbus.ObjectPath, err error) {
	return v.StoreGetStandardSpace(
		<-v.GoGetStandardSpace(flags, make(chan *dbus.Call, 1), standard_space).Done)
}

// method GetProfiles

func (v *manager) GoGetProfiles(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProfiles", flags, ch)
}

func (*manager) StoreGetProfiles(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetProfiles(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetProfiles(
		<-v.GoGetProfiles(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSensors

func (v *manager) GoGetSensors(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSensors", flags, ch)
}

func (*manager) StoreGetSensors(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetSensors(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetSensors(
		<-v.GoGetSensors(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetProfilesByKind

func (v *manager) GoGetProfilesByKind(flags dbus.Flags, ch chan *dbus.Call, kind string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProfilesByKind", flags, ch, kind)
}

func (*manager) StoreGetProfilesByKind(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetProfilesByKind(flags dbus.Flags, kind string) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetProfilesByKind(
		<-v.GoGetProfilesByKind(flags, make(chan *dbus.Call, 1), kind).Done)
}

// method CreateProfileWithFd

func (v *manager) GoCreateProfileWithFd(flags dbus.Flags, ch chan *dbus.Call, profile_id string, scope string, handle dbus.UnixFD, properties map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateProfileWithFd", flags, ch, profile_id, scope, handle, properties)
}

func (*manager) StoreCreateProfileWithFd(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) CreateProfileWithFd(flags dbus.Flags, profile_id string, scope string, handle dbus.UnixFD, properties map[string]string) (object_path dbus.ObjectPath, err error) {
	return v.StoreCreateProfileWithFd(
		<-v.GoCreateProfileWithFd(flags, make(chan *dbus.Call, 1), profile_id, scope, handle, properties).Done)
}

// method CreateProfile

func (v *manager) GoCreateProfile(flags dbus.Flags, ch chan *dbus.Call, profile_id string, scope string, properties map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateProfile", flags, ch, profile_id, scope, properties)
}

func (*manager) StoreCreateProfile(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) CreateProfile(flags dbus.Flags, profile_id string, scope string, properties map[string]string) (object_path dbus.ObjectPath, err error) {
	return v.StoreCreateProfile(
		<-v.GoCreateProfile(flags, make(chan *dbus.Call, 1), profile_id, scope, properties).Done)
}

// method CreateDevice

func (v *manager) GoCreateDevice(flags dbus.Flags, ch chan *dbus.Call, device_id string, scope string, properties map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateDevice", flags, ch, device_id, scope, properties)
}

func (*manager) StoreCreateDevice(call *dbus.Call) (object_path dbus.ObjectPath, err error) {
	err = call.Store(&object_path)
	return
}

func (v *manager) CreateDevice(flags dbus.Flags, device_id string, scope string, properties map[string]string) (object_path dbus.ObjectPath, err error) {
	return v.StoreCreateDevice(
		<-v.GoCreateDevice(flags, make(chan *dbus.Call, 1), device_id, scope, properties).Done)
}

// method DeleteDevice

func (v *manager) GoDeleteDevice(flags dbus.Flags, ch chan *dbus.Call, object_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteDevice", flags, ch, object_path)
}

func (v *manager) DeleteDevice(flags dbus.Flags, object_path dbus.ObjectPath) error {
	return (<-v.GoDeleteDevice(flags, make(chan *dbus.Call, 1), object_path).Done).Err
}

// method DeleteProfile

func (v *manager) GoDeleteProfile(flags dbus.Flags, ch chan *dbus.Call, object_path dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteProfile", flags, ch, object_path)
}

func (v *manager) DeleteProfile(flags dbus.Flags, object_path dbus.ObjectPath) error {
	return (<-v.GoDeleteProfile(flags, make(chan *dbus.Call, 1), object_path).Done).Err
}

// signal Changed

func (v *manager) ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Changed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Changed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceAdded

func (v *manager) ConnectDeviceAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceRemoved

func (v *manager) ConnectDeviceRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal DeviceChanged

func (v *manager) ConnectDeviceChanged(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "DeviceChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".DeviceChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProfileAdded

func (v *manager) ConnectProfileAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProfileAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProfileAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProfileRemoved

func (v *manager) ConnectProfileRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProfileRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProfileRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SensorAdded

func (v *manager) ConnectSensorAdded(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SensorAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SensorAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SensorRemoved

func (v *manager) ConnectSensorRemoved(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SensorRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SensorRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ProfileChanged

func (v *manager) ConnectProfileChanged(cb func(object_path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ProfileChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ProfileChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var object_path dbus.ObjectPath
		err := dbus.Store(sig.Body, &object_path)
		if err == nil {
			cb(object_path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property DaemonVersion s

func (v *manager) DaemonVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DaemonVersion",
	}
}

// property SystemVendor s

func (v *manager) SystemVendor() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SystemVendor",
	}
}

// property SystemModel s

func (v *manager) SystemModel() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SystemModel",
	}
}

type Profile struct {
	profile // interface org.freedesktop.ColorManager.Profile
	proxy.Object
}

func NewProfile(conn *dbus.Conn, path dbus.ObjectPath) (*Profile, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Profile)
	obj.Object.Init_(conn, "org.freedesktop.ColorManager", path)
	return obj, nil
}

type profile struct{}

func (v *profile) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*profile) GetInterfaceName_() string {
	return "org.freedesktop.ColorManager.Profile"
}

// method SetProperty

func (v *profile) GoSetProperty(flags dbus.Flags, ch chan *dbus.Call, property_name string, property_value string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProperty", flags, ch, property_name, property_value)
}

func (v *profile) SetProperty(flags dbus.Flags, property_name string, property_value string) error {
	return (<-v.GoSetProperty(flags, make(chan *dbus.Call, 1), property_name, property_value).Done).Err
}

// method InstallSystemWide

func (v *profile) GoInstallSystemWide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InstallSystemWide", flags, ch)
}

func (v *profile) InstallSystemWide(flags dbus.Flags) error {
	return (<-v.GoInstallSystemWide(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Changed

func (v *profile) ConnectChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Changed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Changed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ProfileId s

func (v *profile) ProfileId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ProfileId",
	}
}

// property Title s

func (v *profile) Title() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Title",
	}
}

// property Metadata a{ss}

func (v *profile) Metadata() PropProfileMetadata {
	return PropProfileMetadata{
		Impl: v,
	}
}

type PropProfileMetadata struct {
	Impl proxy.Implementer
}

func (p PropProfileMetadata) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Metadata", &value)
	return
}

func (p PropProfileMetadata) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Metadata", cb0)
}

// property Qualifier s

func (v *profile) Qualifier() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Qualifier",
	}
}

// property Format s

func (v *profile) Format() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Format",
	}
}

// property Kind s

func (v *profile) Kind() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Kind",
	}
}

// property Colorspace s

func (v *profile) Colorspace() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Colorspace",
	}
}

// property HasVcgt b

func (v *profile) HasVcgt() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasVcgt",
	}
}

// property IsSystemWide b

func (v *profile) IsSystemWide() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsSystemWide",
	}
}

// property Filename s

func (v *profile) Filename() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Filename",
	}
}

// property Created x

func (v *profile) Created() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Created",
	}
}

// property Scope s

func (v *profile) Scope() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Scope",
	}
}

// property Owner u

func (v *profile) Owner() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Owner",
	}
}

// property Warnings as

func (v *profile) Warnings() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Warnings",
	}
}
