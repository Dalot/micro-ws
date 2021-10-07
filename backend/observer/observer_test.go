package observer

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_packet_register(t *testing.T) {
	is := require.New(t)
	obs := &IObserverMock{
		GetIDFunc:  func() string { return "ID_1" },
		UpdateFunc: func([]byte) {},
	}

	tests := []struct {
		name           string
		observableName string
		p              *Packet
		o              *IObserverMock
		err            string
	}{
		{
			name:           "success",
			observableName: "NBA",
			p: &Packet{
				Msg: []byte("a message"),
				DB: &databaseMock{
					SetFunc: func(name string, obs IObserver) error {
						is.EqualValues("NBA", name)
						is.EqualValues("ID_1", obs.GetID())
						return nil
					},
				},
			},
			o: obs,
		},
		{
			name:           "not found",
			observableName: "NBA",
			p: &Packet{
				Msg: []byte("a message"),
				DB: &databaseMock{
					SetFunc: func(name string, obs IObserver) error {
						is.EqualValues("NBA", name)
						is.EqualValues("ID_1", obs.GetID())
						return errors.New("observable not found")
					},
				},
			},
			o:   obs,
			err: "register observer: observable not found",
		},
		{
			name:           "empty name",
			observableName: "",
			p: &Packet{
				Msg: []byte("a message"),
				DB:  &databaseMock{},
			},
			o:   &IObserverMock{},
			err: "name must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.Register(tt.o, tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}
			is.NoError(err)
		})
	}
}

func Test_packet_deregister(t *testing.T) {
	is := require.New(t)

	obs := &IObserverMock{
		GetIDFunc: func() string { return "OBSERVABLE_1" },
	}

	tests := []struct {
		name           string
		observableName string
		p              *Packet
		observer       *IObserverMock
		err            string
	}{
		{
			name:           "success",
			observableName: "CRICKET",
			p: &Packet{
				Msg: []byte("a message"),
				DB: &databaseMock{
					GetByObservableFunc: func(name string) ([]IObserver, error) {
						is.EqualValues("CRICKET", name)
						return []IObserver{
							&IObserverMock{
								GetIDFunc: func() string { return "OBSERVABLE_1" },
							},
						}, nil
					},
					SetObservableFunc: func(name string, observers []IObserver) error {
						is.EqualValues("CRICKET", name)
						return nil
					},
				},
			},
			observer: obs,
		},
		{
			name:           "observable not found",
			observableName: "ESPORTS",
			p: &Packet{
				Msg: []byte("a message"),
				DB: &databaseMock{
					GetByObservableFunc: func(name string) ([]IObserver, error) {
						is.EqualValues("ESPORTS", name)
						return nil, errors.New("observable ESPORTS not found")
					},
				},
			},
			observer: obs,
			err:      "deregister observer: observable ESPORTS not found",
		},
		{
			name:           "observer not found",
			observableName: "ESPORTS",
			p: &Packet{
				Msg: []byte("a message"),
				DB: &databaseMock{
					GetByObservableFunc: func(name string) ([]IObserver, error) {
						is.EqualValues("ESPORTS", name)
						return []IObserver{
							&IObserverMock{
								GetIDFunc: func() string { return "OBSERVABLE_1 " },
							},
						}, nil
					},
				},
			},
			observer: &IObserverMock{
				GetIDFunc: func() string { return "OBSERVABLE_2" },
			},
			err: "could not deregister, did not find any observer with id OBSERVABLE_2 on the observable ESPORTS",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.Deregister(tt.observer, tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}
		})
	}
}

func Test_packet_notifyAll(t *testing.T) {
	is := require.New(t)
	obs1 := &IObserverMock{
		GetIDFunc:  func() string { return "OBSERVABLE_1 " },
		UpdateFunc: func(bytes []byte) {
			is.EqualValues([]byte("a message"), bytes)
		},
	}

	obs2 := &IObserverMock{
		GetIDFunc:  func() string { return "OBSERVABLE_2 " },
		UpdateFunc: func(bytes []byte) {
			is.EqualValues([]byte("a message"), bytes)
		},
	}

	tests := []struct {
		name           string
		p              *Packet
		observableName string
		err            string
	}{
		{
			name:           "success",
			observableName: "FOOTBALL",
			p: &Packet{
				Msg: []byte("a message"), Name: "FOOTBALL",
				DB: &databaseMock{
					GetByObservableFunc: func(name string) ([]IObserver, error) {
						is.EqualValues("FOOTBALL", name)
						return []IObserver{obs1, obs2}, nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.NotifyAll()
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}

			observers, err := tt.p.DB.GetByObservable(tt.observableName)
			is.NoError(err)
			for _, obs := range observers {
				obs, ok := obs.(*IObserverMock)
				is.True(ok)
				is.Len(obs.UpdateCalls(), 1)
			}
		})
	}
}
