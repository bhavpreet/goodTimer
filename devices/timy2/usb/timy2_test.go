package usb

import (
	"bufio"
	"fmt"
	"testing"
)

func Test_timy2USBReader_SubscribeToImpulses(t *testing.T) {
	type fields struct {
		cfg     *Timy2USBConfig
		scanner *bufio.Scanner
	}
	type args struct {
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    chan string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields:  fields{cfg: GetESP32USBConfig()},
			// fields:  fields{cfg: GetTimy2USBConfig()},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &timy2USBReader{
				cfg:     tt.fields.cfg,
				scanner: tt.fields.scanner,
			}
			d.Initialize(d.cfg)
			fmt.Println("Initialized")
			ch, close, err := d.SubscribeToImpulses()
			if (err != nil) != tt.wantErr {
				t.Errorf("timy2USBReader.SubscribeToImpulses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// for i := 0; i < 10; i++ {
			for {
				res := <-ch
				println(res)
				if res == "EOF" {
					break
				}
			}
			close()
		})
	}
}
