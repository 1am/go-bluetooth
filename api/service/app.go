package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"github.com/google/uuid"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/advertising"
	"github.com/muka/go-bluetooth/bluez/profile/agent"
	"github.com/muka/go-bluetooth/bluez/profile/gatt"
	log "github.com/sirupsen/logrus"
)

var AppInterface = "go.bluetooth"
var AppPath = "/%s/apps/%d"

var UseRandomUUID = false

var baseUUID = "00000000-000%d-1000-8000-00805F9B34FB"

var appCounter = 0

func RandomUUID() (string, error) {
	bUUID := fmt.Sprintf(baseUUID, appCounter)
	UUID, err := uuid.Parse(bUUID)
	if UseRandomUUID {
		UUID, err = uuid.NewRandom()
	}
	return strings.ToUpper(UUID.String()), err
}

type AppOptions struct {
	AdapterID         string
	AgentCaps         string
	AgentSetAsDefault bool
}

// Initialize a new bluetooth service (app)
func NewApp(options AppOptions) (*App, error) {

	app := new(App)

	if options.AdapterID == "" {
		return nil, errors.New("options.AdapterID is required")
	}
	app.adapterID = options.AdapterID
	app.services = make(map[dbus.ObjectPath]*Service)
	app.path = dbus.ObjectPath(
		fmt.Sprintf(
			AppPath,
			app.adapterID,
			appCounter,
		),
	)

	app.advertisement = &advertising.LEAdvertisement1Properties{
		Type: advertising.AdvertisementTypePeripheral,
	}

	app.AgentCaps = agent.CapKeyboardDisplay
	if options.AgentCaps != "" {
		app.AgentCaps = options.AgentCaps
	}

	app.AgentSetAsDefault = options.AgentSetAsDefault

	appCounter += 1

	return app, app.init()
}

// Wrap a bluetooth application exposing services
type App struct {
	path     dbus.ObjectPath
	baseUUID string

	adapterID string
	adapter   *adapter.Adapter1

	agent             agent.Agent1Client
	AgentCaps         string
	AgentSetAsDefault bool

	conn          *dbus.Conn
	objectManager *api.DBusObjectManager
	services      map[dbus.ObjectPath]*Service
	advertisement *advertising.LEAdvertisement1Properties
	gm            *gatt.GattManager1
}

func (app *App) init() error {

	// log.Tracef("Exposing %s", app.Path())

	// log.Trace("Load adapter")
	a, err := adapter.NewAdapter1FromAdapterID(app.adapterID)
	if err != nil {
		return err
	}
	app.adapter = a

	agent1, err := app.createAgent()
	if err != nil {
		return err
	}
	app.agent = agent1

	conn, err := dbus.SystemBus()
	if err != nil {
		return err
	}
	app.conn = conn

	_, err = conn.RequestName(
		AppInterface,
		dbus.NameFlagDoNotQueue&dbus.NameFlagReplaceExisting,
	)
	if err != nil {
		return err
	}

	om, err := api.NewDBusObjectManager(app.DBusConn())
	if err != nil {
		return err
	}
	app.objectManager = om

	return err
}

func (app *App) GetAdapter() *adapter.Adapter1 {
	return app.adapter
}

func (app *App) Run() (err error) {

	log.Tracef("Expose %s (%s)", app.Path(), bluez.ObjectManagerInterface)
	err = app.conn.Export(app.objectManager, app.Path(), bluez.ObjectManagerInterface)
	if err != nil {
		return err
	}

	// Expose children services, chars and descriptors
	children := []introspect.Node{}
	for _, service := range app.GetServices() {
		childPath := strings.ReplaceAll(string(service.Path()), string(app.Path())+"/", "")
		children = append(children, introspect.Node{
			Name: childPath,
		})
		// chars
		for _, char := range service.GetChars() {
			childPath := strings.ReplaceAll(string(char.Path()), string(app.Path())+"/", "")
			children = append(children, introspect.Node{
				Name: childPath,
			})
			// descrs
			for _, descr := range char.GetDescr() {
				childPath := strings.ReplaceAll(string(descr.Path()), string(app.Path())+"/", "")
				children = append(children, introspect.Node{
					Name: childPath,
				})
			}
		}
	}

	node := &introspect.Node{
		Interfaces: []introspect.Interface{
			//Introspect
			introspect.IntrospectData,
			//ObjectManager
			bluez.ObjectManagerIntrospectData,
		},
		Children: children,
	}

	introspectable := introspect.NewIntrospectable(node)
	err = app.conn.Export(
		introspectable,
		app.Path(),
		"org.freedesktop.DBus.Introspectable",
	)

	err = app.ExposeAgent(app.AgentCaps, app.AgentSetAsDefault)
	if err != nil {
		return fmt.Errorf("ExposeAgent: %s", err)
	}

	gm, err := gatt.NewGattManager1FromAdapterID(app.adapterID)
	if err != nil {
		return err
	}
	app.gm = gm

	options := map[string]interface{}{}
	err = gm.RegisterApplication(app.Path(), options)

	return err
}

func (app *App) Close() {

	if app.agent != nil {

		err := agent.RemoveAgent(app.agent)
		if err != nil {
			log.Warnf("RemoveAgent: %s", err)
		}

		// err =
		app.agent.Release()
		// if err != nil {
		// 	log.Warnf("Agent1.Release: %s", err)
		// }
	}

	if app.gm != nil {
		err1 := app.gm.UnregisterApplication(app.Path())
		if err1 != nil {
			log.Warnf("GattManager1.UnregisterApplication: %s", err1)
		}
	}
}
