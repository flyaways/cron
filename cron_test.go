package cron

import (
	"reflect"
	"testing"
	"time"
)

func TestNewWithLocation(t *testing.T) {
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Daemon
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithLocation(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Daemon
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
