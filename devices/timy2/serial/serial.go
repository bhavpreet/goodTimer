package serial

import (
	"time"

	"github.com/bhavpreet/goodTimer/devices/driver"
	"github.com/bhavpreet/goodTimer/devices/timy2/sim"
)

type SerialConfig struct {
	SerialPort string
	BaudRate int
}

type SerialReader interface {
	driver.Reader
	InitializeSerial(sconfig *SerialConfig) error
}

func NewTimy2SimDeviceReader() driver.Reader {
	return new(defaultTimySimReader)
}

func NewTimy2SimReader() SerialReader {
	return new(defaultTimySimReader)
}

type defaultTimySimReader struct{
	tsim sim.Timy2Sim
}

func (d *defaultTimySimReader) Initialize(interface{}) error {
	d.tsim = sim.NewTimy2Sim()
	return nil
}

func (d *defaultTimySimReader) InitializeSerial(sconfig *SerialConfig) error {
	 d.tsim = sim.NewTimy2Sim()
	return nil
}

func (d *defaultTimySimReader) SubscribeToImpulses(done chan bool) (chan string, error) {
	return d.tsim.GenerateImpulses(done, time.Second), nil
}
