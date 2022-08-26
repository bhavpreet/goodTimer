package serial

import (
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2/sim"
)

type SerialReader interface {
	Initialize(serialPort string, baudRate int) error
	SubscribeToImpulses(done chan bool) (chan string, error)
}

func NewTimy2SimReader() SerialReader {
	return new(defaultTimySimReader)
}

type defaultTimySimReader struct{
	tsim sim.Timy2Sim
}

func (d *defaultTimySimReader) Initialize(serialPort string, baudRate int) error {
	 d.tsim = sim.NewTimy2Sim()
	return nil
}

func (d *defaultTimySimReader) SubscribeToImpulses(done chan bool) (chan string, error) {
	return d.tsim.GenerateImpulses(done, time.Second), nil
}
