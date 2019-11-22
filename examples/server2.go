// +build

package main

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	//"github.com/paypal/gatt/examples/service"
)

func main() {
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)

	// A mandatory handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			// Setup GAP and GATT services for Linux implementation.
			// OS X doesn't export the access of these services.

			// A simple count service for demo.
			//var s1 = service.NewTestService()

			n:= 0

			s := gatt.NewService(gatt.MustParseUUID("09fc95c0-c111-11e3-9904-0002a5d5c51c"))
			s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51c")).HandleReadFunc(
				func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
					fmt.Fprintf(rsp, "count: %d", n)
					fmt.Println( "count: %d", n)
					n++
				})

			c := gatt.NewCharacteristic(gatt.MustParseUUID("5435D20C-7086-484A-B506-9234873070EA"), s, 0x01 | 0x02, 0, 0)
			dd := gatt.NewDescriptor(gatt.MustParseUUID("2901"), 2901,  c)
			dd.SetValue([]byte("Hello World"))
			c.AddDescriptor(dd.UUID())




			s.AddCharacteristic(gatt.MustParseUUID("5435D20C-7086-484A-B506-9234873070EA")).HandleReadFunc(
				func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
					fmt.Println( "(Println) Characteristic Name: " + c.Name())
					fmt.Println("Descriptor UUID:" + dd.UUID().String())
					fmt.Println("Descriptor Name: " + dd.Name())
					fmt.Println("value byte array: ", []byte("Hello World"))
				})

			d.AddService(s)

			// Advertise device name and service's UUIDs.
			d.AdvertiseNameAndServices("Fred's project", []gatt.UUID{s.UUID()})

			// Advertise as an OpenBeacon iBeacon
			d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)

		default:
		}
	}

	d.Init(onStateChanged)
	select {}
}

