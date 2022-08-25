package sim

import (
	"math/rand"
	"testing"
	"time"
)

func Test_timy2Sim_GenerateImpulses(t *testing.T) {
	type fields struct {
		impulseCounter impulseCounter
		hasStarted     bool
	}
	type args struct {
		done chan bool
	}

	done := make(chan bool)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   chan string
	}{
		{args: args{done: done}},
	}

	rand.Seed(time.Now().UnixMilli())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsim := &timy2Sim{
				impulseCounter: tt.fields.impulseCounter,
				hasStarted:     tt.fields.hasStarted,
			}
			impulse := tsim.GenerateImpulses(tt.args.done, time.Second)
			count := 0
			for i := 0; i < 5; i++ {
				println(<-impulse)
				count++
			}
			done <- true
			if count != 5 {
				t.Errorf("Something is wrong")
			}
		})
	}
}

func Test_timy2Sim_RunningTimer(t *testing.T) {
	type fields struct {
		impulseCounter impulseCounter
		hasStarted     bool
	}
	type args struct {
		done chan bool
	}

	done := make(chan bool)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   chan string
	}{
		{args: args{done: done}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsim := &timy2Sim{
				impulseCounter: tt.fields.impulseCounter,
				hasStarted:     tt.fields.hasStarted,
			}
			runningTimer := tsim.RunningTimer(tt.args.done);
			count := 0
			for i := 0; i < 5; i++ {
				println(<-runningTimer)
				count++
			}
			done <- true
			if count != 5 {
				t.Errorf("Something is wrong")
			}

		})
	}
}
