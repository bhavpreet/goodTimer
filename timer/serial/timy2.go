package serial

import (
	"bufio"
	"bytes"

	"log"

	tsr "github.com/tarm/serial"
)

func NewTimy2Reader() SerialReader {
	return new(timyReader)
}

type timyReader struct {
	scanner    *bufio.Scanner
	serialPort string
	baudRate   int
}

func (tr *timyReader) Initialize(serialPort string, baudRate int) error {
	conf := &tsr.Config{Name: serialPort, Baud: baudRate}
	ser, err := tsr.OpenPort(conf)
	if err != nil {
		log.Printf(
			"Unable to serial.OpenPort on %v, err: %v",
			serialPort, err)
		return err
	}
	tr.scanner = bufio.NewScanner(ser)
	tr.scanner.Split(split)
	tr.serialPort = serialPort
	tr.baudRate = baudRate
	return nil
}

func (tr *timyReader) SubscribeToImpulses(done chan bool) (chan string, error) {
	var impulseChan = make(chan string, 1024)
	go func(done chan bool) {
		for {
			select {
			case done := <-done:
				if done {
					return
				}
			default:
				if tr.scanner.Scan() {
					impulseChan <- tr.scanner.Text()
				}

				if tr.scanner.Err() != nil {
					log.Printf(
						"Error occured while scanning, err: %v",
						tr.scanner.Err())
					// return err
					// We better reinitialize continue!
					tr.Initialize(tr.serialPort, tr.baudRate)
					continue
				}
			}
		}
	}(done)

	return impulseChan, nil
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

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
