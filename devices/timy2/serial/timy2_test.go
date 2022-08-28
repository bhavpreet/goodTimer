package serial

import (
	"bufio"
	"testing"
)

func Test_timyReader_SubscribeToImpulses(t *testing.T) {
	type fields struct {
		scanner *bufio.Scanner
	}
	type args struct {
		done chan bool
	}

	var done chan bool = make(chan bool)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    chan string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:    args{done: done},
			wantErr: false,
		},
	}

	const SERIAL_PORT_NAME = "/dev/cu.usbserial-0001"
	const SERIAL_PORT_BAUD = 38400

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &timyReader{sf: &SerialConfig{
				SerialPort: SERIAL_PORT_NAME,
				BaudRate:   SERIAL_PORT_BAUD,
			}}
			if err := tr.InitializeSerial(tr.sf); err != nil {
				t.Errorf("timyReader.Initialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			stream, err := tr.SubscribeToImpulses(tt.args.done)
			if (err != nil) != tt.wantErr {
				t.Errorf("timyReader.SubscribeToImpulses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			count := 0
			for {
				select {
				case s := <-stream:
					println(s)
					count++
				}
				if count > 10 {
					done <- true
					break
				}

			}
			// if we are here, we are good
		})
	}
}
