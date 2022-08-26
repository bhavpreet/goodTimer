package main

// USB Serial program to output data from a Micro-controller

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2"
	"github.com/bhavpreet/goodTimer/timer/serial"
	"github.com/golang-collections/collections/stack"
)

// adjust port and speed for your setup
const SERIAL_PORT_NAME = "/dev/cu.usbserial-0001"
const SERIAL_PORT_BAUD = 38400

var startStack = stack.New()

// Channel on which impulses are send to the parser
var impulseChan = make(chan string, 128)

func parseImpulse() {
	var impulse string
	for {
		impulse = <-impulseChan
		_parseImpulse(newImpulseInput(impulse))
	}
}

type impulseInput struct {
	s string // Stores the impulse string
	// Impulse format : B####bCxxbHH:MM:SS:zhtq(CR)
	Channel   string
	Timestamp time.Time
}

func newImpulseInput(s string) *impulseInput {
	ii := new(impulseInput)
	ii.s = s
	return ii
}

func (ii *impulseInput) isValidImpulse() bool {
	if len(ii.s) == timy2.ImpulseLength && ii.s[:1] == timy2.B {
		return true
	}
	return false
}

func (ii *impulseInput) String() string {
	return ii.s
}

func (ii *impulseInput) parse() error {
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

	var err error
	if !ii.isValidImpulse() {
		return errors.New("Not a valid impulse: " + ii.String())
	}

	s := ii.String()

	// Inpulse Number - ignore
	s = s[4+1+1:]

	// Channel
	ii.Channel = strings.TrimSpace(s[:4])
	s = s[4:]

	// Timestamp
	ii.Timestamp, err = time.Parse(
		timy2.TimeFormatsForChannels[ii.Channel],
		strings.TrimSpace(s[:13]))
	if err != nil {
		return err
	}

	return nil
}

type Timespan time.Duration

const durationFormat = "15:04:05.000"

func (t Timespan) Format(format string) string {
	_t := time.Date(0, 0, 0, 0, 0, 0, int(time.Duration(t).Nanoseconds()), time.UTC)
	return _t.Format(format)
}

func _parseImpulse(ii *impulseInput) {
	// Standard Time Format
	if ii.isValidImpulse() {
		err := ii.parse()
		if err != nil {
			println("Error Occured", ii.String(),
				"format:", timy2.TimeFormatsForChannels[ii.Channel],
				"err:", err.Error())
		}
		println("Got Impulse:", "["+ii.String()+"]")

		// check channel type / start or end
		if channelType, ok := timy2.ChannelType[ii.Channel]; ok {
			switch channelType {
			case timy2.START_IMPULSE:
				startStack.Push(ii)
			case timy2.END_IMPULSE:
				if start := startStack.Peek(); start == nil {
					println("False Start", ii.Channel)
				} else {
					_start, _ := startStack.Pop().(*impulseInput)
					var t Timespan
					t = Timespan(ii.Timestamp.Sub(_start.Timestamp))
					println("FINISH:", t.Format(durationFormat))
				}
			}

		} else {
			println("Unknown channel type " + ii.Channel)
		}
	}
}

func scanForImpulse() error {
	var err error
	// timy2 := serial.NewTimy2Reader()
	timy2 := serial.NewTimy2SimReader()
	if err = timy2.Initialize(SERIAL_PORT_NAME, SERIAL_PORT_BAUD); err != nil {
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
	parseImpulse()
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
