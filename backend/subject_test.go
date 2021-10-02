package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_packet_register(t *testing.T) {
	is := require.New(t)

	type args struct {
		o *observerMock
	}
	tests := []struct {
		name           string
		observableName string
		p              *packet
		args           args
		err            string
	}{
		{
			name:           "success",
			observableName: "NBA",
			p:              &packet{msg: []byte("a message")},
			args: args{
				o: &observerMock{
					getIDFunc: func() string { return "ok" },
				},
			},
		},
		{
			name:           "not found",
			observableName: "GYM",
			p:              &packet{msg: []byte("a message")},
			args: args{
				o: &observerMock{},
			},
			err: "could not register, did not find any observable named GYM",
		},
		{
			name:           "empty name",
			observableName: "",
			p:              &packet{msg: []byte("a message")},
			args: args{
				o: &observerMock{},
			},
			err: "name must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.register(tt.args.o, tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}
			is.Len(observers[tt.observableName], 1)
			is.EqualValues("ok", (*observers[tt.observableName][0]).getID())
		})
	}
}

func Test_packet_deregister(t *testing.T) {
	is := require.New(t)

	obs := &observerMock{
		getIDFunc: func() string { return "OBSERVABLE_1" },
	}

	tests := []struct {
		name           string
		observableName string
		p              *packet
		deregister     *observerMock
		err            string
	}{
		{
			name:           "success",
			observableName: "CRICKET",
			p:              &packet{msg: []byte("a message")},
			deregister:     obs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.register(tt.deregister, tt.observableName)
			is.NoError(err)

			is.Len(observers[tt.observableName], 1)
			is.EqualValues(tt.deregister.getID(), (*observers[tt.observableName][0]).getID())

			err = tt.p.deregister(tt.deregister, tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}
			is.Len(observers[tt.observableName], 0)
		})
	}
}

func Test_packet_notifyAll(t *testing.T) {
	is := require.New(t)

	tests := []struct {
		name                string
		p                   *packet
		registeredObservers []*observerMock
		observableName      string
	}{
		{
			name: "success",
			p:    &packet{msg: []byte("a message"), name: "FOOTBALL"},
			registeredObservers: []*observerMock{
				{
					updateFunc: func(bytes []byte) {
						is.EqualValues(bytes, []byte("a message"))
					},
				},
				{
					updateFunc: func(bytes []byte) {
						is.EqualValues(bytes, []byte("a message"))
					},
				},
			},
			observableName: "FOOTBALL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, obs := range tt.registeredObservers {
				err := tt.p.register(obs, tt.observableName)
				is.NoError(err)
			}
			tt.p.notifyAll()
			for _, obs := range tt.registeredObservers {
				is.Len(obs.updateCalls(), 1)
			}
		})
	}
}
