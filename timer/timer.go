package main

// USB Serial program to output data from a Micro-controller

import (
	"fmt"
	"log"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2/usb"
	"github.com/bhavpreet/goodTimer/parser"
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
	if err := timy2.Initialize(cfg); err != nil {
		return err
	}

	impulseChan, stiClose, err := timy2.SubscribeToImpulses()
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
			// return fmt.Errorf("Error occured got nil")
		}
		if ii.IsValidImpulse() {
			fmt.Println("XX", ii)
		}
	}
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
