package service

import (
	"fmt"
	"github.com/paypal/gatt"
)

func NewTestService() *gatt.Service {
	n := 0
	s := gatt.NewService(gatt.MustParseUUID("09fc95c0-c111-11e3-9904-0002a5d5c51c"))
	s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51c")).HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			fmt.Fprintf(rsp, "count: %d", n)
			fmt.Println( "count: %d", n)
			n++
		})

	c := gatt.NewCharacteristic(gatt.MustParseUUID("5435D20C-7086-484A-B506-9234873070EA"), s, 0x01 | 0x02 | 0x08, 0, 0)
	d := gatt.NewDescriptor(gatt.MustParseUUID("2901"), 0,  nil)
	d.SetValue([]byte("Hello World"))
	c.AddDescriptor(gatt.MustParseUUID("2901"))

	s.AddCharacteristic(gatt.MustParseUUID("5435D20C-7086-484A-B506-9234873070EA"))

	return s
}
