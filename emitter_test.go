package goemitter_test

import (
	"testing"
	"time"

	"github.com/shadyzoz/goemitter"
)

type MyEmitter struct {
	goemitter.EventEmitter
}

func TestEventEmitter(t *testing.T) {
	type args struct {
		eventName    string
		eventHandler func(...interface{})
		eventArgs    []interface{}
	}
	resultMap := make(map[string]int)
	tests := []struct {
		name   string
		args   []args
		result map[string]int
	}{
		{
			name: "singleListenerWithoutArguments",
			args: []args{
				args{
					eventName: "event_a",
					eventHandler: func(args ...interface{}) {
						resultMap["event_a"]++
					},
				},
			},
			result: map[string]int{
				"event_a": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emitter := MyEmitter{}
			for _, args := range tt.args {
				emitter.On(args.eventName, args.eventHandler)
				emitter.Emit(args.eventName, args.eventArgs...)
			}
			time.Sleep(time.Millisecond)
			for k, v := range tt.result {
				if resultMap[k] != v {
					t.Errorf("event %s's result does not match, %d, %d", k, resultMap[k], v)
				}
			}
		})
	}
}
