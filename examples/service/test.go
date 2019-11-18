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
			//fmt.Fprintf(rsp, "count: %d", n)
			fmt.Printf( "count: %d", n)
			n++
		})

	return s
}
