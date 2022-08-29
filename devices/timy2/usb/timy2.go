package usb

import (
	"bufio"
	"bytes"
	"fmt"
	"log"

	"github.com/bhavpreet/goodTimer/devices/driver"
	"github.com/google/gousb"
)

func NewTimy2DefaultReader() driver.Reader {
	return new(defaultTimy2USBSimReader)
}

func NewTimy2USBReader() USBReader {
	return new(timy2USBReader)
}

func NewTimy2Reader() driver.Reader {
	return NewTimy2USBReader()
}

type timy2USBReader struct {
	cfg     *Timy2USBConfig
	scanner *bufio.Scanner
}

type Timy2USBConfig struct {
	VendorID  gousb.ID
	ProductID gousb.ID
}

func GetTimy2USBConfig() *Timy2USBConfig {
	const TIMY_VEND = 0x0c4a // USB Vendor ID
	const TIMY_PROD = 0x088a // USB Product ID Timy (2)

	return &Timy2USBConfig{
		VendorID:  TIMY_VEND,
		ProductID: TIMY_PROD,
	}
}

func GetESP32USBConfig() *Timy2USBConfig {
	const TIMY_VEND = 0x10c4
	const TIMY_PROD = 0xea60

	return &Timy2USBConfig{
		VendorID:  TIMY_VEND,
		ProductID: TIMY_PROD,
	}
}

func (d *timy2USBReader) Initialize(cfg interface{}) error {

	tCfg, ok := cfg.(*Timy2USBConfig)
	if !ok {
		return fmt.Errorf("Invalid config type, should be *Timy2USBConfig")
	}
	d.cfg = tCfg

	// // Initialize a new Context.
	// ctx := gousb.NewContext()
	// defer ctx.Close()

	// // Open any device with a given VID/PID using a convenience function.
	// // dev, err := ctx.OpenDeviceWithVIDPID(TIMY_VEND, TIMY_PROD)
	// dev, err := ctx.OpenDeviceWithVIDPID(d.cfg.VendorID, d.cfg.ProductID)
	// if err != nil {
	// 	log.Fatalf("Could not open a device: %v", err)
	// }
	// defer dev.Close()

	// // 38400 = 0x9600 ~> 0x00 0x96 in little endian
	// encoding := []byte{0x00, 0x96, 0x00, 0x00, 0x00, 0x00, 0x08}
	// dev.Control(0x21, 0x20, 0, 0, encoding)

	// // Claim the default interface using a convenience function.
	// // The default interface is always #0 alt #0 in the currently active
	// // config.
	// _cfg, err := dev.Config(1)
	// if err != nil {
	// 	log.Printf("ERR: %s.Config(1): %v", dev, err)
	// 	return err
	// }
	// defer _cfg.Close()

	// intf, done, err := dev.DefaultInterface()
	// if err != nil {
	// 	log.Printf("ERR: %s.DefaultInterface(): %v", dev, err)
	// 	return err
	// }
	// defer done()

	// // Read from endpoint
	// ep, err := intf.InEndpoint(0x01)
	// if err != nil {
	// 	log.Printf("ERR: %s.InEndpoint(0x01): %v", intf, err)
	// 	return err
	// }

	// stream, err := ep.NewStream(ep.Endpoints[0x01].MaxPacketSize*10, 1)
	// if err != nil {
	// 	log.Printf("ERR: %s.InEndpoint(0x01).NewStream: %v", intf, err)
	// 	return err
	// }

	// dropCR := // dropCR drops a terminal \r from the data.
	// 	func(data []byte) []byte {
	// 		if len(data) > 0 && data[len(data)-1] == '\r' {
	// 			return data[0 : len(data)-1]
	// 		}
	// 		return data
	// 	}

	// split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 	if atEOF && len(data) == 0 {
	// 		return 0, nil, nil
	// 	}
	// 	if i := bytes.IndexByte(data, '\r'); i >= 0 {
	// 		// We have a full newline-terminated line.
	// 		return i + 1, dropCR(data[0:i]), nil
	// 	}
	// 	// If we're at EOF, we have a final, non-terminated line. Return it.
	// 	if atEOF {
	// 		return len(data), dropCR(data), nil
	// 	}
	// 	// Request more data.
	// 	return 0, nil, nil
	// }

	// scanner := bufio.NewScanner(stream)
	// scanner.Split(split)
	// d.scanner = scanner

	return nil
}

func (d *timy2USBReader) SubscribeToImpulses(done chan bool) (chan string, error) {
	out := make(chan string, 128)
	go func(end chan bool, out chan string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic occurred:", err)
				return
			}
		}()
		// Initialize a new Context.
		ctx := gousb.NewContext()
		defer ctx.Close()

		// Open any device with a given VID/PID using a convenience function.
		// dev, err := ctx.OpenDeviceWithVIDPID(TIMY_VEND, TIMY_PROD)
		dev, err := ctx.OpenDeviceWithVIDPID(d.cfg.VendorID, d.cfg.ProductID)
		if err != nil {
			log.Fatalf("Could not open a device: %v", err)
		}
		defer dev.Close()

		// 38400 = 0x9600 ~> 0x00 0x96 in little endian
		// 9600  = 0x2580 ~> 0x80 0x25
		encoding := []byte{0x80, 0x25, 0x00, 0x00, 0x00, 0x00, 0x08}
		dev.Control(0x21, 0x20, 0, 0, encoding)

		// Claim the default interface using a convenience function.
		// The default interface is always #0 alt #0 in the currently active
		// config.
		_cfg, err := dev.Config(1)
		if err != nil {
			log.Printf("ERR: %s.Config(1): %v", dev, err)
			return
		}
		defer _cfg.Close()

		intf, done, err := dev.DefaultInterface()
		if err != nil {
			log.Printf("ERR: %s.DefaultInterface(): %v", dev, err)
			return
		}
		defer done()

		// Read from endpoint
		ep, err := intf.InEndpoint(0x01)
		if err != nil {
			log.Printf("ERR: %s.InEndpoint(0x01): %v", intf, err)
			return
		}

		stream, err := ep.NewStream(ep.Endpoints[0x01].MaxPacketSize*10, 1)
		if err != nil {
			log.Printf("ERR: %s.InEndpoint(0x01).NewStream: %v", intf, err)
			return
		}

		dropCR := // dropCR drops a terminal \r from the data.
			func(data []byte) []byte {
				if len(data) > 0 && data[len(data)-1] == '\r' {
					return data[0 : len(data)-1]
				}
				return data
			}

		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			if i := bytes.IndexByte(data, '\r'); i >= 0 {
				// We have a full newline-terminated line.
				return i + 1, dropCR(data[0:i]), nil
			}
			// If we're at EOF, we have a final, non-terminated line. Return it.
			if atEOF {
				return len(data), dropCR(data), nil
			}
			// Request more data.
			return 0, nil, nil
		}

		scanner := bufio.NewScanner(stream)
		scanner.Split(split)
		d.scanner = scanner

		for {
			select {
			case _end := <-end:
				if _end {
					return
				}
			default:
				if d.scanner.Scan() {
					out <- d.scanner.Text()
				}
			}
		}

	}(done, out)
	return out, nil
}
