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

package media



import (
  "sync"
  "github.com/muka/go-bluetooth/bluez"
  "reflect"
  "github.com/fatih/structs"
  "github.com/muka/go-bluetooth/util"
  "github.com/godbus/dbus"
)

var MediaItem1Interface = "org.bluez.MediaItem1"


// NewMediaItem1 create a new instance of MediaItem1
//
// Args:
// 	servicePath: unique name
// 	objectPath: freely definable
func NewMediaItem1(servicePath string, objectPath dbus.ObjectPath) (*MediaItem1, error) {
	a := new(MediaItem1)
	a.propertiesSignal = make(chan *dbus.Signal)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: MediaItem1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaItem1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}

// NewMediaItem1Controller create a new instance of MediaItem1
//
// Args:
// 	objectPath: [variable	prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX/playerX/itemX
func NewMediaItem1Controller(objectPath dbus.ObjectPath) (*MediaItem1, error) {
	a := new(MediaItem1)
	a.propertiesSignal = make(chan *dbus.Signal)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: MediaItem1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaItem1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


// MediaItem1 MediaItem1 hierarchy

type MediaItem1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	Properties 				*MediaItem1Properties
}

// MediaItem1Properties contains the exposed properties of an interface
type MediaItem1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	// Player Player object path the item belongs to
	Player dbus.ObjectPath

	// Type Item type
  // Possible values: "video", "audio", "folder"
	Type string

	// Title Item title name
  // Available if property Type is "audio"
  // or "video"
	Title string

	// Duration Item duration in milliseconds
  // Available if property Type is "audio"
  // or "video"
	Duration uint32

	// Album Item album name
  // Available if property Type is "audio"
  // or "video"
	Album string

	// Genre Item genre name
  // Available if property Type is "audio"
  // or "video"
	Genre string

	// NumberOfTracks Item album number of tracks in total
  // Available if property Type is "audio"
  // or "video"
	NumberOfTracks uint32

	// Name Item displayable name
	Name string

	// FolderType Folder type.
  // Possible values: "mixed", "titles", "albums", "artists"
  // Available if property Type is "Folder"
	FolderType string

	// Playable Indicates if the item can be played
  // Available if property Type is "folder"
	Playable bool

	// Metadata Item metadata.
  // Possible values:
	Metadata map[string]interface{}

	// Artist Item artist name
  // Available if property Type is "audio"
  // or "video"
	Artist string

	// Number Item album number
  // Available if property Type is "audio"
  // or "video"
	Number uint32

}

func (p *MediaItem1Properties) Lock() {
	p.lock.Lock()
}

func (p *MediaItem1Properties) Unlock() {
	p.lock.Unlock()
}

// Close the connection
func (a *MediaItem1) Close() {
	
	a.unregisterSignal()
	
	a.client.Disconnect()
}

// Path return MediaItem1 object path
func (a *MediaItem1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Interface return MediaItem1 interface
func (a *MediaItem1) Interface() string {
	return a.client.Config.Iface
}


// ToMap convert a MediaItem1Properties to map
func (a *MediaItem1Properties) ToMap() (map[string]interface{}, error) {
	return structs.Map(a), nil
}

// FromMap convert a map to an MediaItem1Properties
func (a *MediaItem1Properties) FromMap(props map[string]interface{}) (*MediaItem1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an MediaItem1Properties
func (a *MediaItem1Properties) FromDBusMap(props map[string]dbus.Variant) (*MediaItem1Properties, error) {
	s := new(MediaItem1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// GetProperties load all available properties
func (a *MediaItem1) GetProperties() (*MediaItem1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *MediaItem1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *MediaItem1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *MediaItem1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *MediaItem1) unregisterSignal() {
	if a.propertiesSignal == nil {
		a.propertiesSignal <- nil
	}
}

// WatchProperties updates on property changes
func (a *MediaItem1) WatchProperties() (chan *bluez.PropertyChanged, error) {

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

				// updates [*]Properties struct
				props := a.Properties

				s := reflect.ValueOf(props).Elem()
				// exported field
				f := s.FieldByName(field)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						x := reflect.ValueOf(val.Value())
						props.Lock()
						f.Set(x)
						props.Unlock()
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

func (a *MediaItem1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	ch <- nil
	close(ch)
	return nil
}





//Play Play item
// Possible Errors: org.bluez.Error.NotSupported
// org.bluez.Error.Failed
func (a *MediaItem1) Play() error {
	
	return a.client.Call("Play", 0, ).Store()
	
}

//AddtoNowPlaying Add item to now playing list
// Possible Errors: org.bluez.Error.NotSupported
// org.bluez.Error.Failed
func (a *MediaItem1) AddtoNowPlaying() error {
	
	return a.client.Call("AddtoNowPlaying", 0, ).Store()
	
}

