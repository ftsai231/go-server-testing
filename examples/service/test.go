package service

import (
	"fmt"

	"github.com/paypal/gatt"
)

func NewTestService() *gatt.Service {
	n := 0
	var str string = "testing string"
	s := gatt.NewService(gatt.MustParseUUID("09fc95c0-c111-11e3-9904-0002a5d5c51c"))
	s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51c")).HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			fmt.Fprintf(rsp, "count: %d", n)
			fmt.Fprintf(rsp, str)
			fmt.Println( "count: %d", n)

			n++
		})

	//c := gatt.NewCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51d"), s, 0x01 | 0x02 | 0x08, 0, 0)
	//d := gatt.NewDescriptor(gatt.MustParseUUID("2901"), 0,  c)
	//d.SetValue([]byte("Hello World"))
	//
	//
	//
	//c.AddDescriptor(gatt.UUID16(0x2904)).SetValue([]byte{4, 1, 39, 173, 1, 0, 0})
	//
	//
	//s.AddCharacteristic(gatt.MustParseUUID("11fac9e0-c111-11e3-9246-0002a5d5c51d")).HandleReadFunc(
	//	func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
	//		fmt.Fprintf(rsp, str)
	//		fmt.Println(str)
	//
	//	})

	return s
}
