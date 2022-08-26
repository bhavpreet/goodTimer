package sim

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bhavpreet/goodTimer/devices/timy2"
)

type Timy2Sim interface {
	GenerateImpulses(done chan bool, interval time.Duration) chan string
	RunningTimer(done chan bool) chan string
}


// NewTimy2Sim returns fresh instance of timy2 simulator
func NewTimy2Sim() Timy2Sim {
	rand.Seed(time.Now().UnixMilli())
	return new(timy2Sim)
}

type impulseCounter int

func (ic impulseCounter) String() string {
	s := strconv.Itoa(int(ic))
	if remaining := 4 - len(s); remaining > 0 {
		s = strings.Repeat("0", remaining) + s
	}
	return s
}

type timy2Sim struct {
	impulseCounter impulseCounter
	hasStarted     bool
}

// GenerateImpulses generates impulses
func (tsim *timy2Sim) GenerateImpulses(done chan bool, interval time.Duration) chan string {
	impulses := make(chan string, 1024)
	if interval == 0 {
		interval = 5 * time.Second
	}
	go func(done chan bool) {
		t := time.NewTicker(interval)
		for {
			select {
			case _done := <-done:
				if _done {
					return
				}
			case <-t.C:
				impulses <- tsim.getImpulse()
			}
		}
	}(done)
	return impulses
}


func (tsim *timy2Sim) RunningTimer(done chan bool) chan string {
	runningTimer := make(chan string)
	go func(done chan bool) {
		t := time.NewTicker(time.Second / 10)
		for {
			select {
			case _done := <-done:
				if _done {
					return
				}
			case <-t.C:
				runningTimer <- time.Now().Format(timy2.RunningTimeFormat) + "\r"
			}
		}
	}(done)
	return runningTimer
}

// Gets channels
func (tsim *timy2Sim) getChannel() string {
	// C0/C0M   => start channel
	// C1/C1M   => finish channel
	// C2 – C8  => timing channels
	startChannels := []string{"C0", "C0M"}
	endChannels := []string{"C1", "C1M"}
	otherChannels := []string{"C2", "C3", "C4", "C5", "C6", "C7", "C8"}
	_ = otherChannels

	if !tsim.hasStarted {
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

func (tsim *timy2Sim) getImpulse() string {
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
	now := time.Now()
	timyChannel := tsim.getChannel()
	timeFormatForChannel := timy2.TimeFormatsForChannels[timyChannel]

	impulse = B +
		tsim.impulseCounter.String() + B +
		padAfter(timyChannel, 3, " ") + B +
		padAfter(now.Format(timeFormatForChannel),
			timy2.ImpulseTimeStampLength, " ") +
		CR

	tsim.impulseCounter++
	_ = impulse
	return impulse
}

func padAfter(s string, size int, padWith string) string {
	ret := s
	l := len(s)
	if l < size {
		ret = ret + strings.Repeat(padWith, size-l)
	}
	return ret
}
