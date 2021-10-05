package database

import (
	"backend/observer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObservers_GetByObservable(t *testing.T) {
	is := require.New(t)
	db := &Observers{}
	data := map[string][]observer.IObserver{
		"SOCCER": {},
		"NBA":    {},
	}
	db.SetAll(data)
	soccerObservers := []observer.IObserver{
		&observer.IObserverMock{},
	}

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
			}
		})
	}
}
