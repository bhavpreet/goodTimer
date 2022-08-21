package main

import (
	"machine"
	"math/rand"
	"strconv"
	"strings"
	"time"

	timy2 "github.com/bhavpreet/timy2/constants"
)


type ImpulseCounter int

var impulseCounter ImpulseCounter = 30 // start from some arbitary number

func (ic ImpulseCounter) String() string {
	s := strconv.Itoa(int(ic))
	if remaining := 4 - len(s); remaining > 0 {
		s = strings.Repeat("0", remaining) + s
	}
	return s
}

func runningTimer() {
	// Toggle LED
	ledToggle := true
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if ledToggle {
			led.Low()
		} else {
			led.High()
		}
		ledToggle = !ledToggle

		time.Sleep(time.Second / 10)
		print(time.Now().Format(timy2.RunningTimeFormat) + "\r")
	}
}

var hasStarted bool

func padAfter(s string, size int, padWith string) string {
	ret := s
	l := len(s)
	if l < size {
		ret = ret + strings.Repeat(padWith, size - l)
	}
	return ret
}

// Gets channels
func getChannel() string {
	// C0/C0M   => start channel
	// C1/C1M   => finish channel
	// C2 – C8  => timing channels
	startChannels := []string{"C0", "C0M"}
	endChannels := []string{"C1", "C1M"}
	otherChannels := []string{"C2", "C3", "C4", "C5", "C6", "C7", "C8"}
	_ = otherChannels

	if !hasStarted {
		if rand.Int()%2 == 0 {
			// False Start
			return endChannels[rand.Int()%2]
		} else {
			return startChannels[rand.Int()%2]
		}
	} else {
		return startChannels[rand.Int()%2]
	}
}

func waitForImpuse() {
	// Impulse format : B####bCxxbHH:MM:SS:zhtq(CR)
	// Example        :  0033 C0  07:50:40.2828 00
	//
	// B ...................... Blank
	// #### ................. subsequent number of start number
	// Cxx................... channel (see below, if only 2 figures than additional blank) HH.................... hours
	// :........................ separation
	// MM ................... minutes
	// SS .................... seconds
	// z ....................... 1/10 seconds
	// h....................... 1/100 seconds
	// t........................ 1/1.000 seconds
	// q....................... 1/10.000 seconds
	// (CR) ................. Carriage Return

	B := " "
	CR := "\r"
	LF := "\n"
	_ = LF
	var impulse string
	for {
		time.Sleep(time.Second * 5)
		now := time.Now()
		timyChannel := getChannel()
		timeFormatForChannel := timy2.TimeFormatsForChannels[timyChannel]
		impulse = B +
			impulseCounter.String() + B +
			padAfter(timyChannel, 3, " ") + B +
			now.Format(timeFormatForChannel) +
			CR

		impulseCounter++
		_ = impulse
		print(impulse)
	}
}

func init() {
	machine.Serial.Configure(machine.UARTConfig{BaudRate: timy2.BaudRate})
	// Clear
	print("\r\n")
}

func main() {
	rand.Seed(time.Now().UnixMilli()) // Always going to be same on esp?
	go waitForImpuse()
	runningTimer()
}
