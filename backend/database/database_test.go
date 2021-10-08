package database

import (
	"backend/observer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestObservers_GetByObservable(t *testing.T) {
	is := require.New(t)
	db := &Observers{}
	obs := &observer.IObserverMock{
		GetIDFunc: func() string { return "ID_1" },
	}
	data := map[string][]observer.IObserver{
		"SOCCER": {obs},
		"NBA":    {},
	}
	db.SetAll(data)
	soccerObservers := []observer.IObserver{obs}

	tests := []struct {
		name           string
		observableName string
		expected       []observer.IObserver
		err            string
	}{
		{
			name:           "success",
			observableName: "SOCCER",
			expected:       soccerObservers,
		},
		{
			name:           "not found",
			observableName: "FOOTBALL",
			err:            "observable FOOTBALL not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.GetByObservable(tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}

			for i, obs := range got {
				obs, ok := obs.(*observer.IObserverMock)
				is.True(ok)
				is.EqualValues(tt.expected[i], obs)
				is.EqualValues(tt.expected[i].GetID(), obs.GetID())
			}
		})
	}
}

func TestObservers_Set(t *testing.T) {
	is := require.New(t)
	db := &Observers{}
	obs := &observer.IObserverMock{
		GetIDFunc: func() string { return "ID_1" },
	}
	data := map[string][]observer.IObserver{
		"SOCCER": {},
		"NBA":    {},
	}
	db.SetAll(data)

	tests := []struct {
		name           string
		observableName string
		input          observer.IObserver
		expected       []observer.IObserver
		len            int
		err            string
	}{
		{
			name:           "success",
			observableName: "SOCCER",
			input:          obs,
			len:            1,
			expected:       []observer.IObserver{obs},
		},
		{
			name:           "not found",
			observableName: "PINBALL",
			input:          obs,
			len:            1,

			err: "observable PINBALL not found",
		},
		{
			name:           "already exists",
			observableName: "SOCCER",
			input:          obs,
			len:            1,
			err:            "observer ID_1 is already listening to SOCCER",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.Set(tt.observableName, tt.input)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}

			got, err := db.GetByObservable(tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}
			for i, obs := range got {
				obs, ok := obs.(*observer.IObserverMock)
				is.True(ok)
				is.EqualValues(tt.expected[i], obs)
				is.EqualValues(tt.expected[i].GetID(), obs.GetID())
			}

		})
	}
}

func TestObservers_SetObservable(t *testing.T) {
	is := require.New(t)
	db := &Observers{}
	obs := &observer.IObserverMock{
		GetIDFunc: func() string { return "ID_1" },
	}
	data := map[string][]observer.IObserver{
		"SOCCER": {},
		"NBA":    {},
	}
	db.SetAll(data)

	tests := []struct {
		name           string
		observableName string
		input          []observer.IObserver
		expected       []observer.IObserver
		len            int
		err            string
	}{
		{
			name:           "success",
			observableName: "SOCCER",
			input:          []observer.IObserver{obs},
			len:            1,
			expected:       []observer.IObserver{obs},
		},
		{
			name:           "not found",
			observableName: "GYM",
			input:          []observer.IObserver{obs},
			err:            "observable GYM not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.SetObservable(tt.observableName, tt.input)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}

			got, err := db.GetByObservable(tt.observableName)
			if err != nil {
				is.EqualError(err, tt.err)
				return
			}
			is.Len(got, 1)
			for i, obs := range got {
				obs, ok := obs.(*observer.IObserverMock)
				is.True(ok)
				is.EqualValues(tt.expected[i], obs)
				is.EqualValues(tt.expected[i].GetID(), obs.GetID())
			}

		})
	}
}
