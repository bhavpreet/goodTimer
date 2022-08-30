package main

// USB Serial program to output data from a Micro-controller

import (
	"log"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2/usb"
	parser "github.com/bhavpreet/goodTimer/praser"
)

// Channel on which impulses are send to the parser
var impulseChan = make(chan string, 128)

func scanForImpulse() error {
	var err error
	// timy2 := serial.NewTimy2Reader()
	// timy2 := serial.NewTimy2SimReader()
	// timy2 := usb.NewTimy2SimDeviceReader()
	// if err = timy2.Initialize(nil); err != nil {
	// 	return err
	// }

	cfg := usb.GetESP32USBConfig()
	// cfg := usb.GetTimy2USBConfig()

	timy2 := usb.NewTimy2Reader()
	if err := timy2.Initialize(cfg) ; err != nil {
		return err
	}


	var done chan bool = make(chan bool)
	defer func() {
		done <- true
	}()

	impulseChan, err = timy2.SubscribeToImpulses(done)
	if err != nil {
		return err
	}
	parser.ParseImpulse(impulseChan)
	return nil
}


func main() {
	for {
		if err := scanForImpulse(); err != nil {
			log.Printf("ERR: %v", err)
		}
		log.Println("Sleeping 5 secs..")
		time.Sleep(5 * time.Second)
	}
}
