package usb

import (
	"time"

	"github.com/bhavpreet/goodTimer/devices/driver"
	"github.com/bhavpreet/goodTimer/devices/timy2/sim"
)

type USBReader interface {
	driver.Reader
}

func NewTimy2SimDeviceReader() driver.Reader {
	return new(defaultTimy2USBSimReader)
}

func NewTimy2SimReader() USBReader {
	return new(defaultTimy2USBSimReader)
}

type defaultTimy2USBSimReader struct{
	tsim sim.Timy2Sim
}

func (d *defaultTimy2USBSimReader) Initialize(interface{}) error {
	d.tsim = sim.NewTimy2Sim()
	return nil
}

func (d *defaultTimy2USBSimReader) SubscribeToImpulses(done chan bool) (chan string, error) {
	return d.tsim.GenerateImpulses(done, time.Second), nil
}
