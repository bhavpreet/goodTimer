package timer

// USB Serial program to get data from a Micro-controller

import (
	"fmt"
	"log"

	"github.com/bhavpreet/goodTimer/devices/driver"
	"github.com/bhavpreet/goodTimer/devices/timy2/usb"
	"github.com/bhavpreet/goodTimer/parser"
	"github.com/timshannon/bolthold"
)

// Channel on which impulses are send to the parser
var impulseChan = make(chan string, 128)

type Timer struct {
	driver.Reader
	cfg     *TimerConfig
	store   *bolthold.Store
	process func(*parser.Impulse) error
}

type TimerType string

const (
	SIMULATOR TimerType = "SIMULATOR"
	ESP32               = "ESP32"
	TIMY2               = "TIMY2"
)

type TimerConfig struct {
	TimerType TimerType
}

func NewTimer(cfg *TimerConfig, processor func(ii *parser.Impulse) error) (*Timer, error) {
	t := new(Timer)
	t.cfg = cfg
	t.process = processor
	return t, nil
}

func (t *Timer) Run() error {
	switch t.cfg.TimerType {
	case SIMULATOR:
		t.Reader = usb.NewTimy2SimReader()
		t.Initialize(nil)
	case ESP32:
		cfg := usb.GetESP32USBConfig()
		// cfg := usb.GetTimy2USBConfig()

		t.Reader = usb.NewTimy2Reader()
		if err := t.Initialize(cfg); err != nil {
			return err
		}
	case TIMY2:
		cfg := usb.GetTimy2USBConfig()

		t.Reader = usb.NewTimy2Reader()
		if err := t.Initialize(cfg); err != nil {
			return err
		}
	default:
		return fmt.Errorf("No config found %v", t.cfg.TimerType)
	}

	impulseChan, stiClose, err := t.SubscribeToImpulses()
	if err != nil {
		return err
	}
	defer stiClose()

	ic, close, err := parser.ParseImpulse(impulseChan)
	if err != nil {
		return err
	}
	defer close()
	for {
		ii := <-ic
		if ii == nil {
			log.Println("Error occured got nil")
			break
		}
		if ii.IsValidImpulse() {
			fmt.Println("XX", ii)
			err = t.process(ii)
			if err != nil {
				fmt.Println("Error while processing impulse: ", err)
			}
		}
	}
	return fmt.Errorf("Some error occurred")
}

// func (t *Timer) scanForImpulse() error {
// 	var err error
// 	timy2 := usb.NewTimy2SimReader()
// 	timy2.Initialize(nil)
// 	// timy2 := usb.NewTimy2SimDeviceReader()
// 	// if err = timy2.Initialize(nil); err != nil {
// 	// 	return err
// 	// }

// 	// cfg := usb.GetESP32USBConfig()
// 	// // cfg := usb.GetTimy2USBConfig()

// 	// timy2 := usb.NewTimy2Reader()
// 	// if err := timy2.Initialize(cfg); err != nil {
// 	// 	return err
// 	// }

// 	impulseChan, stiClose, err := timy2.SubscribeToImpulses()
// 	if err != nil {
// 		return err
// 	}
// 	defer stiClose()

// 	ic, close, err := parser.ParseImpulse(impulseChan)
// 	if err != nil {
// 		return err
// 	}
// 	defer close()
// 	for {
// 		ii := <-ic
// 		if ii == nil {
// 			log.Println("Error occured got nil")
// 			break
// 			// return fmt.Errorf("Error occured got nil")
// 		}
// 		if ii.IsValidImpulse() {
// 			fmt.Println("XX", ii)
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	for {
// 		if err := scanForImpulse(); err != nil {
// 			log.Printf("ERR: %v", err)
// 		}
// 		log.Println("Sleeping 5 secs..")
// 		time.Sleep(5 * time.Second)
// 	}
// }
