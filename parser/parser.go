package parser

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2"
)

type Impulse struct {
	s string // Stores the impulse string
	// Impulse format : B####bCxxbHH:MM:SS:zhtq(CR)
	Channel   string
	Timestamp time.Time
}

func NewImpulse(s string) *Impulse {
	ii := new(Impulse)
	ii.s = timy2.B + strings.TrimSpace(s)
	return ii
}

func (ii *Impulse) IsValidImpulse() bool {
	if len(ii.s) == timy2.ImpulseLength && ii.s[:1] == timy2.B {
		return true
	}
	return false
}

func (ii *Impulse) String() string {
	return ii.s
}

func (ii *Impulse) parse() error {
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
	// SPACE
	// 00
	// (CR) ................. Carriage Return

	var err error
	if !ii.IsValidImpulse() {
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
		log.Printf("Error parsing time, err: %v", err)
		return err
	}

	return nil
}

type Timespan time.Duration

const DurationFormat = "15:04:05.000"

func (t Timespan) Format(format string) string {

	z := time.Unix(0, 0).UTC()
	return z.Add(time.Duration(t)).Format(format)
}

func _parseImpulse(ii *Impulse) {
	// Standard Time Format
	if ii.IsValidImpulse() {
		err := ii.parse()
		if err != nil {
			println("Error Occured", ii.String(),
				"format:", timy2.TimeFormatsForChannels[ii.Channel],
				"err:", err.Error())
		}
		// println("Got Impulse:", "["+ii.String()+"]")
	}
}

// return nil in `chan *Impulse in case of error`
func ParseImpulse(impulseChan chan string) (chan *Impulse, func(), error) {
	var done chan bool = make(chan bool)
	var exited bool
	close := func() {
		if !exited {
			done <- true
		}
	}

	var ic chan *Impulse = make(chan *Impulse, 128)

	go func() {
		defer func() {
			exited = true
		}()

		for {
			select {
			case <-done:
				return
			default:
				i := <-impulseChan
				if i == "EOF" {
					ic <- nil
					return
				}
				ii := NewImpulse(i)
				_parseImpulse(ii)
				ic <- ii
			}
		}
	}()

	return ic, close, nil
}
