package main

// USB Serial program to output data from a Micro-controller

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2"
	"github.com/golang-collections/collections/stack"
	"github.com/tarm/serial"
)

// adjust port and speed for your setup
const SERIAL_PORT_NAME = "/dev/cu.usbserial-0001"
const SERIAL_PORT_BAUD = 38400

var startStack = stack.New()

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func split(
	data []byte,
	atEOF bool) (advance int, token []byte, err error) {

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
const durationFormat = "15:05:05.000"

func (t Timespan) Format(format string) string {
	z := time.Unix(0, 0).UTC()
	return z.Add(time.Duration(t)).Format(format)
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
		println("Parsing Impulse:",
			ii.Channel, ii.Timestamp.String(), "["+ii.String()+"]")

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

func scanForImpulse() {
	go parseImpulse()
	conf := &serial.Config{Name: SERIAL_PORT_NAME, Baud: SERIAL_PORT_BAUD}
	ser, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	scanner := bufio.NewScanner(ser)
	scanner.Split(split)
	for scanner.Scan() {
		impulseChan <- scanner.Text()
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

func main() {
	scanForImpulse()
}
