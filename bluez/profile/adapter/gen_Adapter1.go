// WARNING: generated code, do not edit!
// Copyright © 2019 luca capra
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package adapter



import (
  "sync"
  "github.com/muka/go-bluetooth/bluez"
  "reflect"
  "github.com/fatih/structs"
  "github.com/muka/go-bluetooth/util"
  "github.com/godbus/dbus"
  "fmt"
)

var Adapter1Interface = "org.bluez.Adapter1"


// NewAdapter1 create a new instance of Adapter1
//
// Args:
// 	objectPath: [variable prefix]/{hci0,hci1,...}
func NewAdapter1(objectPath dbus.ObjectPath) (*Adapter1, error) {
	a := new(Adapter1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: Adapter1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Adapter1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}

// NewAdapter1FromAdapterID create a new instance of Adapter1
// adapterID: ID of an adapter eg. hci0
func NewAdapter1FromAdapterID(adapterID string) (*Adapter1, error) {
	a := new(Adapter1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: Adapter1Interface,
			Path:  dbus.ObjectPath(fmt.Sprintf("/org/bluez/%s", adapterID)),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Adapter1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


// Adapter1 Adapter hierarchy

type Adapter1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Adapter1Properties
}

// Adapter1Properties contains the exposed properties of an interface
type Adapter1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	// Class The Bluetooth class of device.
  // This property represents the value that is either
  // automatically configured by DMI/ACPI information
  // or provided as static configuration.
	Class uint32

	// UUIDs List of 128-bit UUIDs that represents the available
  // local services.
	UUIDs []string

	// Discoverable Switch an adapter to discoverable or non-discoverable
  // to either make it visible or hide it. This is a global
  // setting and should only be used by the settings
  // application.
  // If the DiscoverableTimeout is set to a non-zero
  // value then the system will set this value back to
  // false after the timer expired.
  // In case the adapter is switched off, setting this
  // value will fail.
  // When changing the Powered property the new state of
  // this property will be updated via a PropertiesChanged
  // signal.
  // For any new adapter this settings defaults to false.
	Discoverable bool

	// Pairable Switch an adapter to pairable or non-pairable. This is
  // a global setting and should only be used by the
  // settings application.
  // Note that this property only affects incoming pairing
  // requests.
  // For any new adapter this settings defaults to true.
	Pairable bool

	// PairableTimeout The pairable timeout in seconds. A value of zero
  // means that the timeout is disabled and it will stay in
  // pairable mode forever.
  // The default value for pairable timeout should be
  // disabled (value 0).
	PairableTimeout uint32

	// Address The Bluetooth device address.
	Address string

	// AddressType The Bluetooth  Address Type. For dual-mode and BR/EDR
  // only adapter this defaults to "public". Single mode LE
  // adapters may have either value. With privacy enabled
  // this contains type of Identity Address and not type of
  // address used for connection.
  // Possible values:
  // "public" - Public address
  // "random" - Random address
	AddressType string

	// Name The Bluetooth system name (pretty hostname).
  // This property is either a static system default
  // or controlled by an external daemon providing
  // access to the pretty hostname configuration.
	Name string

	// Alias The Bluetooth friendly name. This value can be
  // changed.
  // In case no alias is set, it will return the system
  // provided name. Setting an empty string as alias will
  // convert it back to the system provided name.
  // When resetting the alias with an empty string, the
  // property will default back to system name.
  // On a well configured system, this property never
  // needs to be changed since it defaults to the system
  // name and provides the pretty hostname. Only if the
  // local name needs to be different from the pretty
  // hostname, this property should be used as last
  // resort.
	Alias string

	// Powered Switch an adapter on or off. This will also set the
  // appropriate connectable state of the controller.
  // The value of this property is not persistent. After
  // restart or unplugging of the adapter it will reset
  // back to false.
	Powered bool

	// DiscoverableTimeout The discoverable timeout in seconds. A value of zero
  // means that the timeout is disabled and it will stay in
  // discoverable/limited mode forever.
  // The default value for the discoverable timeout should
  // be 180 seconds (3 minutes).
	DiscoverableTimeout uint32

	// Discovering Indicates that a device discovery procedure is active.
	Discovering bool

	// Modalias Local Device ID information in modalias format
  // used by the kernel and udev.
	Modalias string

}

func (p *Adapter1Properties) Lock() {
	p.lock.Lock()
}

func (p *Adapter1Properties) Unlock() {
	p.lock.Unlock()
}


// SetClass set Class value
func (a *Adapter1) SetClass(v uint32) error {
	return a.SetProperty("Class", v)
}

// GetClass get Class value
func (a *Adapter1) GetClass() (uint32, error) {
	v, err := a.GetProperty("Class")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}

// SetUUIDs set UUIDs value
func (a *Adapter1) SetUUIDs(v []string) error {
	return a.SetProperty("UUIDs", v)
}

// GetUUIDs get UUIDs value
func (a *Adapter1) GetUUIDs() ([]string, error) {
	v, err := a.GetProperty("UUIDs")
	if err != nil {
		return []string{}, err
	}
	return v.Value().([]string), nil
}

// SetDiscoverable set Discoverable value
func (a *Adapter1) SetDiscoverable(v bool) error {
	return a.SetProperty("Discoverable", v)
}

// GetDiscoverable get Discoverable value
func (a *Adapter1) GetDiscoverable() (bool, error) {
	v, err := a.GetProperty("Discoverable")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetPairable set Pairable value
func (a *Adapter1) SetPairable(v bool) error {
	return a.SetProperty("Pairable", v)
}

// GetPairable get Pairable value
func (a *Adapter1) GetPairable() (bool, error) {
	v, err := a.GetProperty("Pairable")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetPairableTimeout set PairableTimeout value
func (a *Adapter1) SetPairableTimeout(v uint32) error {
	return a.SetProperty("PairableTimeout", v)
}

// GetPairableTimeout get PairableTimeout value
func (a *Adapter1) GetPairableTimeout() (uint32, error) {
	v, err := a.GetProperty("PairableTimeout")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}

// SetAddress set Address value
func (a *Adapter1) SetAddress(v string) error {
	return a.SetProperty("Address", v)
}

// GetAddress get Address value
func (a *Adapter1) GetAddress() (string, error) {
	v, err := a.GetProperty("Address")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetAddressType set AddressType value
func (a *Adapter1) SetAddressType(v string) error {
	return a.SetProperty("AddressType", v)
}

// GetAddressType get AddressType value
func (a *Adapter1) GetAddressType() (string, error) {
	v, err := a.GetProperty("AddressType")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetName set Name value
func (a *Adapter1) SetName(v string) error {
	return a.SetProperty("Name", v)
}

// GetName get Name value
func (a *Adapter1) GetName() (string, error) {
	v, err := a.GetProperty("Name")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetAlias set Alias value
func (a *Adapter1) SetAlias(v string) error {
	return a.SetProperty("Alias", v)
}

// GetAlias get Alias value
func (a *Adapter1) GetAlias() (string, error) {
	v, err := a.GetProperty("Alias")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetPowered set Powered value
func (a *Adapter1) SetPowered(v bool) error {
	return a.SetProperty("Powered", v)
}

// GetPowered get Powered value
func (a *Adapter1) GetPowered() (bool, error) {
	v, err := a.GetProperty("Powered")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetDiscoverableTimeout set DiscoverableTimeout value
func (a *Adapter1) SetDiscoverableTimeout(v uint32) error {
	return a.SetProperty("DiscoverableTimeout", v)
}

// GetDiscoverableTimeout get DiscoverableTimeout value
func (a *Adapter1) GetDiscoverableTimeout() (uint32, error) {
	v, err := a.GetProperty("DiscoverableTimeout")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}

// SetDiscovering set Discovering value
func (a *Adapter1) SetDiscovering(v bool) error {
	return a.SetProperty("Discovering", v)
}

// GetDiscovering get Discovering value
func (a *Adapter1) GetDiscovering() (bool, error) {
	v, err := a.GetProperty("Discovering")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetModalias set Modalias value
func (a *Adapter1) SetModalias(v string) error {
	return a.SetProperty("Modalias", v)
}

// GetModalias get Modalias value
func (a *Adapter1) GetModalias() (string, error) {
	v, err := a.GetProperty("Modalias")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}


// Close the connection
func (a *Adapter1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Adapter1 object path
func (a *Adapter1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Interface return Adapter1 interface
func (a *Adapter1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Adapter1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Adapter1Properties to map
func (a *Adapter1Properties) ToMap() (map[string]interface{}, error) {
	return structs.Map(a), nil
}

// FromMap convert a map to an Adapter1Properties
func (a *Adapter1Properties) FromMap(props map[string]interface{}) (*Adapter1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Adapter1Properties
func (a *Adapter1Properties) FromDBusMap(props map[string]dbus.Variant) (*Adapter1Properties, error) {
	s := new(Adapter1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// GetProperties load all available properties
func (a *Adapter1) GetProperties() (*Adapter1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Adapter1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Adapter1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Adapter1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Adapter1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Adapter1) WatchProperties() (chan *bluez.PropertyChanged, error) {

	channel, err := a.client.Register(a.Path(), a.Interface())
	if err != nil {
		return nil, err
	}

	ch := make(chan *bluez.PropertyChanged)

	go (func() {
		for {

			if channel == nil {
				break
			}

			sig := <-channel

			if sig == nil {
				return
			}

			if sig.Name != bluez.PropertiesChanged {
				continue
			}
			if sig.Path != a.Path() {
				continue
			}

			iface := sig.Body[0].(string)
			changes := sig.Body[1].(map[string]dbus.Variant)

			for field, val := range changes {

				// updates [*]Properties struct when a property change
				s := reflect.ValueOf(a.Properties).Elem()
				// exported field
				f := s.FieldByName(field)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						x := reflect.ValueOf(val.Value())
						a.Properties.Lock()
						f.Set(x)
						a.Properties.Unlock()
					}
				}

				propChanged := &bluez.PropertyChanged{
					Interface: iface,
					Name:      field,
					Value:     val.Value(),
				}
				ch <- propChanged
			}

		}
	})()

	return ch, nil
}

func (a *Adapter1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	ch <- nil
	close(ch)
	return nil
}




//StartDiscovery This method starts the device discovery session. This
// includes an inquiry procedure and remote device name
// resolving. Use StopDiscovery to release the sessions
// acquired.
// This process will start creating Device objects as
// new devices are discovered.
// During discovery RSSI delta-threshold is imposed.
// Possible errors: org.bluez.Error.NotReady
// org.bluez.Error.Failed
func (a *Adapter1) StartDiscovery() error {
	
	return a.client.Call("StartDiscovery", 0, ).Store()
	
}

//StopDiscovery This method will cancel any previous StartDiscovery
// transaction.
// Note that a discovery procedure is shared between all
// discovery sessions thus calling StopDiscovery will only
// release a single session.
// Possible errors: org.bluez.Error.NotReady
// org.bluez.Error.Failed
// org.bluez.Error.NotAuthorized
func (a *Adapter1) StopDiscovery() error {
	
	return a.client.Call("StopDiscovery", 0, ).Store()
	
}

//RemoveDevice This removes the remote device object at the given
// path. It will remove also the pairing information.
// Possible errors: org.bluez.Error.InvalidArguments
// org.bluez.Error.Failed
func (a *Adapter1) RemoveDevice(device dbus.ObjectPath) error {
	
	return a.client.Call("RemoveDevice", 0, device).Store()
	
}

//SetDiscoveryFilter This method sets the device discovery filter for the
// caller. When this method is called with no filter
// parameter, filter is removed.
// Parameters that may be set in the filter dictionary
// include the following:
// array{string} UUIDs
// Filter by service UUIDs, empty means match
// _any_ UUID.
// When a remote device is found that advertises
// any UUID from UUIDs, it will be reported if:
// - Pathloss and RSSI are both empty.
// - only Pathloss param is set, device advertise
// TX pwer, and computed pathloss is less than
// Pathloss param.
// - only RSSI param is set, and received RSSI is
// higher than RSSI param.
// int16 RSSI
// RSSI threshold value.
// PropertiesChanged signals will be emitted
// for already existing Device objects, with
// updated RSSI value. If one or more discovery
// filters have been set, the RSSI delta-threshold,
// that is imposed by StartDiscovery by default,
// will not be applied.
// uint16 Pathloss
// Pathloss threshold value.
// PropertiesChanged signals will be emitted
// for already existing Device objects, with
// updated Pathloss value.
// string Transport (Default "auto")
// Transport parameter determines the type of
// scan.
// Possible values:
// "auto"	- interleaved scan
// "bredr"	- BR/EDR inquiry
// "le"	- LE scan only
// If "le" or "bredr" Transport is requested,
// and the controller doesn't support it,
// org.bluez.Error.Failed error will be returned.
// If "auto" transport is requested, scan will use
// LE, BREDR, or both, depending on what's
// currently enabled on the controller.
// bool DuplicateData (Default: true)
// Disables duplicate detection of advertisement
// data.
// When enabled PropertiesChanged signals will be
// generated for either ManufacturerData and
// ServiceData everytime they are discovered.
// When discovery filter is set, Device objects will be
// created as new devices with matching criteria are
// discovered regardless of they are connectable or
// discoverable which enables listening to
// non-connectable and non-discoverable devices.
// When multiple clients call SetDiscoveryFilter, their
// filters are internally merged, and notifications about
// new devices are sent to all clients. Therefore, each
// client must check that device updates actually match
// its filter.
// When SetDiscoveryFilter is called multiple times by the
// same client, last filter passed will be active for
// given client.
// SetDiscoveryFilter can be called before StartDiscovery.
// It is useful when client will create first discovery
// session, to ensure that proper scan will be started
// right after call to StartDiscovery.
// Possible errors: org.bluez.Error.NotReady
// org.bluez.Error.NotSupported
// org.bluez.Error.Failed
func (a *Adapter1) SetDiscoveryFilter(filter map[string]interface{}) error {
	
	return a.client.Call("SetDiscoveryFilter", 0, filter).Store()
	
}

//GetDiscoveryFilters Return available filters that can be given to
// SetDiscoveryFilter.
// Possible errors: None
func (a *Adapter1) GetDiscoveryFilters() ([]string, error) {
	
	var val0 []string
	err := a.client.Call("GetDiscoveryFilters", 0, ).Store(&val0)
	return val0, err	
}

//ConnectDevice This method connects to device without need of
// performing General Discovery. Connection mechanism is
// similar to Connect method from Device1 interface with
// exception that this method returns success when physical
// connection is established. After this method returns,
// services discovery will continue and any supported
// profile will be connected. There is no need for calling
// Connect on Device1 after this call. If connection was
// successful this method returns object path to created
// device object.
// Parameters that may be set in the filter dictionary
// include the following:
// string Address
// The Bluetooth device address of the remote
// device. This parameter is mandatory.
// string AddressType
// The Bluetooth device Address Type. This is
// address type that should be used for initial
// connection. If this parameter is not present
// BR/EDR device is created.
// Possible values:
// "public" - Public address
// "random" - Random address
// Possible errors: org.bluez.Error.InvalidArguments
// org.bluez.Error.AlreadyExists
// org.bluez.Error.NotSupported
// org.bluez.Error.NotReady
// org.bluez.Error.Failed
func (a *Adapter1) ConnectDevice(properties map[string]interface{}) (dbus.ObjectPath, error) {
	
	var val0 dbus.ObjectPath
	err := a.client.Call("ConnectDevice", 0, properties).Store(&val0)
	return val0, err	
}

