package cron

import (
	"testing"
	"time"
)

func TestDaemon_Add(t *testing.T) {
	type fields struct {
		schedules Schedules
		snapshot  chan Schedules
		running   bool
		stop      chan struct{}
		add       chan *Schedule
		del       chan *Schedule
		update    chan *Schedule
		now       time.Time
		location  *time.Location
	}
	type args struct {
		j Job
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Daemon{
				schedules: tt.fields.schedules,
				snapshot:  tt.fields.snapshot,
				running:   tt.fields.running,
				stop:      tt.fields.stop,
				add:       tt.fields.add,
				del:       tt.fields.del,
				update:    tt.fields.update,
				now:       tt.fields.now,
				location:  tt.fields.location,
			}
			if err := c.Add(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("Daemon.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDaemon_Del(t *testing.T) {
	type fields struct {
		schedules Schedules
		snapshot  chan Schedules
		running   bool
		stop      chan struct{}
		add       chan *Schedule
		del       chan *Schedule
		update    chan *Schedule
		now       time.Time
		location  *time.Location
	}
	type args struct {
		j Job
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Daemon{
				schedules: tt.fields.schedules,
				snapshot:  tt.fields.snapshot,
				running:   tt.fields.running,
				stop:      tt.fields.stop,
				add:       tt.fields.add,
				del:       tt.fields.del,
				update:    tt.fields.update,
				now:       tt.fields.now,
				location:  tt.fields.location,
			}
			if err := c.Del(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("Daemon.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDaemon_Update(t *testing.T) {
	type fields struct {
		schedules Schedules
		snapshot  chan Schedules
		running   bool
		stop      chan struct{}
		add       chan *Schedule
		del       chan *Schedule
		update    chan *Schedule
		now       time.Time
		location  *time.Location
	}
	type args struct {
		j Job
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Daemon{
				schedules: tt.fields.schedules,
				snapshot:  tt.fields.snapshot,
				running:   tt.fields.running,
				stop:      tt.fields.stop,
				add:       tt.fields.add,
				del:       tt.fields.del,
				update:    tt.fields.update,
				now:       tt.fields.now,
				location:  tt.fields.location,
			}
			if err := c.Update(tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("Daemon.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDaemon_runWithRecovery(t *testing.T) {
	type fields struct {
		schedules Schedules
		snapshot  chan Schedules
		running   bool
		stop      chan struct{}
		add       chan *Schedule
		del       chan *Schedule
		update    chan *Schedule
		now       time.Time
		location  *time.Location
	}
	type args struct {
		j Job
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Daemon{
				schedules: tt.fields.schedules,
				snapshot:  tt.fields.snapshot,
				running:   tt.fields.running,
				stop:      tt.fields.stop,
				add:       tt.fields.add,
				del:       tt.fields.del,
				update:    tt.fields.update,
				now:       tt.fields.now,
				location:  tt.fields.location,
			}
			c.runWithRecovery(tt.args.j)
		})
	}
}

func TestDaemon_run(t *testing.T) {
	type fields struct {
		schedules Schedules
		snapshot  chan Schedules
		running   bool
		stop      chan struct{}
		add       chan *Schedule
		del       chan *Schedule
		update    chan *Schedule
		now       time.Time
		location  *time.Location
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Daemon{
				schedules: tt.fields.schedules,
				snapshot:  tt.fields.snapshot,
				running:   tt.fields.running,
				stop:      tt.fields.stop,
				add:       tt.fields.add,
				del:       tt.fields.del,
				update:    tt.fields.update,
				now:       tt.fields.now,
				location:  tt.fields.location,
			}
			c.run()
		})
	}
}
