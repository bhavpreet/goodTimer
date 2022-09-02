package timer

import (
	"testing"

	"github.com/bhavpreet/goodTimer/devices/driver"
)

func TestTimer_Run(t *testing.T) {
	type fields struct {
		Reader driver.Reader
		cfg    *TimerConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := NewTimer(&TimerConfig{timerType: ESP32})
			if err != nil {
				t.Errorf("Timer.Run() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := tr.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Timer.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
